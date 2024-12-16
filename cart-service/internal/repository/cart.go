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
	Checkout(ctx context.Context, user_id string) (float64, error)
}

type Repository struct {
	DataBase *redis.Client
}

func NewRepository(db *redis.Client) RepositoryInterface {
	return &Repository{DataBase: db}
}

func (r *Repository) Add(ctx context.Context, userID string, product *domain.CartItem) (bool, error) {
	key := fmt.Sprintf("cart:%s", userID)

	existing, err := r.DataBase.HGet(ctx, key, product.ProductID).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}

	quantity := product.Quantity
	if existing != "" {
		var existingProduct domain.CartItem
		if err := json.Unmarshal([]byte(existing), &existingProduct); err != nil {
			return false, err
		}
		quantity += existingProduct.Quantity
	}

	product.Quantity = quantity

	data, err := json.Marshal(product)
	if err != nil {
		return false, err
	}
	if err := r.DataBase.HSet(ctx, key, product.ProductID, data).Err(); err != nil {
		return false, err
	}

	_, err = r.recalculateTotalCost(ctx, key)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) Delete(ctx context.Context, userID, productID string, quantity int) (bool, error) {
	key := fmt.Sprintf("cart:%s", userID)

	existing, err := r.DataBase.HGet(ctx, key, productID).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	var existingProduct domain.CartItem
	if err := json.Unmarshal([]byte(existing), &existingProduct); err != nil {
		return false, err
	}

	if existingProduct.Quantity <= quantity {
		if err := r.DataBase.HDel(ctx, key, productID).Err(); err != nil {
			return false, err
		}
	} else {
		existingProduct.Quantity -= quantity
		data, err := json.Marshal(existingProduct)
		if err != nil {
			return false, err
		}
		if err := r.DataBase.HSet(ctx, key, productID, data).Err(); err != nil {
			return false, err
		}
	}

	_, err = r.recalculateTotalCost(ctx, key)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *Repository) Get(ctx context.Context, userID string) (*domain.Cart, error) {
	key := fmt.Sprintf("cart:%s", userID)

	items, err := r.DataBase.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	cart := &domain.Cart{
		UserID: userID,
		Items:  make([]*domain.CartItem, 0),
	}

	totalCost := 0.0
	for _, item := range items {
		var cartItem domain.CartItem
		if err := json.Unmarshal([]byte(item), &cartItem); err != nil {
			return nil, err
		}
		cart.Items = append(cart.Items, &cartItem)
		totalCost += cartItem.Price * float64(cartItem.Quantity)
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

func (r *Repository) Checkout(ctx context.Context, userID string) (float64, error) {
	cart, err := r.Get(ctx, userID)
	if err != nil {
		return 0, err
	}
	return cart.TotalCost, nil
}

func (r *Repository) recalculateTotalCost(ctx context.Context, key string) (float64, error) {
	items, err := r.DataBase.HGetAll(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	totalCost := 0.0
	for _, item := range items {
		var cartItem domain.CartItem
		if err := json.Unmarshal([]byte(item), &cartItem); err != nil {
			return 0, err
		}
		totalCost += cartItem.Price * float64(cartItem.Quantity)
	}

	if err := r.DataBase.HSet(ctx, key, "total_cost", totalCost).Err(); err != nil {
		return 0, err
	}

	return totalCost, nil
}
