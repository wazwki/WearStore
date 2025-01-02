package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/wazwki/WearStore/cart-service/internal/domain"
)

type RepositoryInterface interface {
	Add(ctx context.Context, user_id string, product *domain.CartItem) (bool, error)
	Delete(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	Get(ctx context.Context, user_id string) (*domain.Cart, error)
	Clear(ctx context.Context, user_id string) (bool, error)
}

type Repository struct {
	DataBase *redis.Client
}

func NewRepository(db *redis.Client) RepositoryInterface {
	return &Repository{DataBase: db}
}

func (r *Repository) Add(ctx context.Context, userID string, product *domain.CartItem) (bool, error) {
	key := fmt.Sprintf("cart:%s", userID)

	existeng, err := r.DataBase.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}

	var cartItems []domain.CartItem
	if existeng != "" {
		if err := json.Unmarshal([]byte(existeng), &cartItems); err != nil {
			return false, err
		}
	}

	changed := false
	for i := range cartItems {
		if cartItems[i].ProductID == product.ProductID {
			cartItems[i].Quantity += product.Quantity
			changed = true
			break
		}
	}

	if !changed {
		cartItems = append(cartItems, *product)
	}

	data, err := json.Marshal(cartItems)
	if err != nil {
		return false, err
	}
	if err := r.DataBase.Set(ctx, key, data, 0).Err(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) Delete(ctx context.Context, userID, productID string, quantity int) (bool, error) {
	key := fmt.Sprintf("cart:%s", userID)

	existeng, err := r.DataBase.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	var cartItems []domain.CartItem
	if existeng != "" {
		if err := json.Unmarshal([]byte(existeng), &cartItems); err != nil {
			return false, err
		}
	}

	for i, item := range cartItems {
		if item.ProductID == productID {
			cartItems[i].Quantity -= quantity
			if cartItems[i].Quantity <= 0 {
				cartItems = append(cartItems[:i], cartItems[i+1:]...)
			}
			break
		}
	}

	data, err := json.Marshal(cartItems)
	if err != nil {
		return false, err
	}
	if err := r.DataBase.Set(ctx, key, data, 0).Err(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) Get(ctx context.Context, userID string) (*domain.Cart, error) {
	key := fmt.Sprintf("cart:%s", userID)

	items, err := r.DataBase.Get(ctx, key).Result()
	if err == redis.Nil {
		return &domain.Cart{
			UserID:    userID,
			Items:     make([]*domain.CartItem, 0),
			TotalCost: 0.0,
		}, nil
	} else if err != nil {
		return nil, err
	}

	cart := &domain.Cart{
		UserID: userID,
		Items:  make([]*domain.CartItem, 0),
	}

	if items != "" {
		var cartItems []*domain.CartItem
		if err := json.Unmarshal([]byte(items), &cartItems); err != nil {
			return nil, err
		}
		cart.Items = cartItems
	}

	totalCost, err := r.recalculateTotalCost(ctx, key)
	if err != nil {
		return nil, err
	}

	cart.TotalCost = totalCost

	return cart, nil
}

func (r *Repository) Clear(ctx context.Context, userID string) (bool, error) {
	key := fmt.Sprintf("cart:%s", userID)
	if err := r.DataBase.Del(ctx, key).Err(); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Repository) recalculateTotalCost(ctx context.Context, key string) (float64, error) {
	items, err := r.DataBase.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	var cartItems []domain.CartItem
	if items != "" {
		if err := json.Unmarshal([]byte(items), &cartItems); err != nil {
			return 0, err
		}
	}

	totalCost := 0.0
	for _, item := range cartItems {
		totalCost += item.Price * float64(item.Quantity)
	}

	return totalCost, nil
}
