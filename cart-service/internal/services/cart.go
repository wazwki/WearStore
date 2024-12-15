package services

import (
	"context"
	"fmt"

	"github.com/wazwki/WearStore/cart-service/internal/domain"
	"github.com/wazwki/WearStore/cart-service/internal/repository"
)

type ServiceInterface interface {
	AddToCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	RemoveFromCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	GetCart(ctx context.Context, user_id string) (*domain.Cart, error)
	ClearCart(ctx context.Context, user_id string) (bool, error)
	Checkout(ctx context.Context, user_id string) (bool, float64, error)
}

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{repo: repo}
}

func (s *Service) AddToCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error) {
	_, err := s.repo.Add(ctx, user_id, product_id, quantity)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) RemoveFromCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error) {
	_, err := s.repo.Delete(ctx, user_id, product_id, quantity)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) GetCart(ctx context.Context, user_id string) (*domain.Cart, error) {
	cart, err := s.repo.Get(ctx, user_id)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *Service) ClearCart(ctx context.Context, user_id string) (bool, error) {
	success, err := s.repo.Clear(ctx, user_id)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *Service) Checkout(ctx context.Context, user_id string) (bool, float64, error) {
	cart, err := s.repo.Get(ctx, user_id)
	if err != nil {
		return false, 0, err
	}
	success, err := s.repo.Clear(ctx, user_id)
	if err != nil {
		return false, 0, err
	}
	if !success {
		return false, 0, fmt.Errorf("error clearing cart: %v", err)
	}
	return success, cart.TotalCost, nil
}
