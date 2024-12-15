package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/wazwki/WearStore/cart-service/internal/domain"
)

type RepositoryInterface interface {
	Add(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	Delete(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	Get(ctx context.Context, user_id string) (*domain.Cart, error)
	Clear(ctx context.Context, user_id string) (bool, error)
	Checkout(ctx context.Context, user_id string) (bool, float64, error)
}

type Repository struct {
	DataBase *redis.Client
}

func NewRepository(db *redis.Client) RepositoryInterface {
	return &Repository{DataBase: db}
}

func (repo *Repository) Add(ctx context.Context, user_id, product_id string, quantity int) (bool, error) {
	_, err := repo.DataBase.HIncrBy(ctx, user_id, product_id, int64(quantity)).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *Repository) Delete(ctx context.Context, user_id, product_id string, quantity int) (bool, error) {
	_, err := repo.DataBase.HIncrBy(ctx, user_id, product_id, int64(-quantity)).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *Repository) Get(ctx context.Context, user_id string) (*domain.Cart, error) {
	cart := &domain.Cart{
		UserID: user_id,
	}

	members, err := repo.DataBase.HKeys(ctx, user_id).Result()
	if err != nil {
		return nil, err
	}

	for _, member := range members {
		value, err := repo.DataBase.HGet(ctx, user_id, member).Int()
		if err != nil {
			return nil, err
		}
		cart.Items = append(cart.Items, &domain.CartItem{
			ProductID: member,
			Quantity:  value,
		})
	}

	return cart, nil
}

func (repo *Repository) Clear(ctx context.Context, user_id string) (bool, error) {
	_, err := repo.DataBase.Del(ctx, user_id).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *Repository) Checkout(ctx context.Context, user_id string) (bool, float64, error) {
	cart, err := repo.Get(ctx, user_id)
	if err != nil {
		return false, 0, err
	}
	success, err := repo.Clear(ctx, user_id)
	if err != nil {
		return false, 0, err
	}
	return success, cart.TotalCost, nil
}
