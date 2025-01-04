package services

import (
	"context"
	"fmt"
	"time"

	"github.com/wazwki/WearStore/cart-service/internal/clients"
	"github.com/wazwki/WearStore/cart-service/internal/domain"
	"github.com/wazwki/WearStore/cart-service/internal/repository"
	"github.com/wazwki/WearStore/cart-service/pkg/metrics"
)

type ServiceInterface interface {
	AddToCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	RemoveFromCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error)
	GetCart(ctx context.Context, user_id string) (*domain.Cart, error)
	ClearCart(ctx context.Context, user_id string) (bool, error)
	Checkout(ctx context.Context, user_id string) (float64, error)
}

type Service struct {
	repo    repository.RepositoryInterface
	product clients.ProductClient
	user    clients.UserClient
}

func NewService(repo repository.RepositoryInterface, product clients.ProductClient, user clients.UserClient) ServiceInterface {
	return &Service{repo: repo, product: product, user: user}
}

func (s *Service) AddToCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error) {
	start := time.Now()
	if quantity <= 0 {
		return false, fmt.Errorf("quantity must be greater than zero")
	}

	product, err := s.product.GetProduct(ctx, product_id)
	if err != nil {
		return false, err
	}
	product.Quantity = quantity

	ok, err := s.repo.Add(ctx, user_id, product)
	if err != nil {
		return false, err
	}

	metrics.ServiceDuration.WithLabelValues("cart-service.AddToCart").Observe(time.Since(start).Seconds())
	return ok, nil
}

func (s *Service) RemoveFromCart(ctx context.Context, user_id, product_id string, quantity int) (bool, error) {
	start := time.Now()
	if quantity <= 0 {
		return false, fmt.Errorf("quantity must be greater than zero")
	}

	_, err := s.repo.Delete(ctx, user_id, product_id, quantity)
	if err != nil {
		return false, err
	}

	metrics.ServiceDuration.WithLabelValues("cart-service.RemoveFromCart").Observe(time.Since(start).Seconds())
	return true, nil
}

func (s *Service) GetCart(ctx context.Context, user_id string) (*domain.Cart, error) {
	start := time.Now()
	cart, err := s.repo.Get(ctx, user_id)
	if err != nil {
		return nil, err
	}

	metrics.ServiceDuration.WithLabelValues("cart-service.GetCart").Observe(time.Since(start).Seconds())
	return cart, nil
}

func (s *Service) ClearCart(ctx context.Context, user_id string) (bool, error) {
	start := time.Now()
	success, err := s.repo.Clear(ctx, user_id)
	if err != nil {
		return false, err
	}

	metrics.ServiceDuration.WithLabelValues("cart-service.ClearCart").Observe(time.Since(start).Seconds())
	return success, nil
}

func (s *Service) Checkout(ctx context.Context, user_id string) (float64, error) {
	start := time.Now()
	_, err := s.user.GetUser(ctx, user_id)
	if err != nil {
		return 0, err
	}

	cart, err := s.repo.Get(ctx, user_id)
	if err != nil {
		return 0, err
	}

	var totalCost float64
	for _, item := range cart.Items {

		product, err := s.product.GetProduct(ctx, item.ProductID)
		if err != nil {
			return 0, err
		}

		if product.Quantity < item.Quantity {
			return 0, fmt.Errorf("not enough stock for product %v", item.ProductID)
		}

		_, err = s.product.UpdateProduct(ctx, &domain.CartItem{
			ProductID: item.ProductID,
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  item.Quantity,
		})
		if err != nil {
			return 0, err
		}

		totalCost += item.Price * float64(item.Quantity)
	}

	success, err := s.repo.Clear(ctx, user_id)
	if err != nil || !success {
		return 0, fmt.Errorf("failed to clear cart after checkout: %w", err)
	}

	metrics.ServiceDuration.WithLabelValues("cart-service.Checkout").Observe(time.Since(start).Seconds())
	return totalCost, nil
}
