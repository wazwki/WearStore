package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wazwki/WearStore/product-service/internal/domain"
	"github.com/wazwki/WearStore/product-service/internal/repository"
	"github.com/wazwki/WearStore/product-service/pkg/metrics"
)

type ServiceInterface interface {
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	ListProducts(ctx context.Context, limit, offset int64) ([]*domain.Product, error)
	AddProduct(ctx context.Context, newProduct *domain.Product) (string, error)
	UpdateProduct(ctx context.Context, updatingProduct *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
}

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{repo: repo}
}

func (s *Service) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	start := time.Now()
	product, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, domain.ErrProductNotFound
	}

	metrics.ServiceDuration.WithLabelValues("product-service.GetProduct").Observe(time.Since(start).Seconds())
	return product, nil
}

func (s *Service) ListProducts(ctx context.Context, limit, offset int64) ([]*domain.Product, error) {
	start := time.Now()
	products, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	metrics.ServiceDuration.WithLabelValues("product-service.ListProducts").Observe(time.Since(start).Seconds())
	return products, nil
}

func (s *Service) AddProduct(ctx context.Context, newProduct *domain.Product) (string, error) {
	start := time.Now()
	newProduct.CreatedAt = time.Now()
	newProduct.UpdatedAt = time.Now()
	newProduct.ID = uuid.New().String()
	id, err := s.repo.Create(ctx, newProduct)
	if err != nil {
		return "", err
	}

	metrics.ServiceDuration.WithLabelValues("product-service.AddProduct").Observe(time.Since(start).Seconds())
	return id, nil
}

func (s *Service) UpdateProduct(ctx context.Context, updatingProduct *domain.Product) (*domain.Product, error) {
	start := time.Now()
	product, err := s.repo.Get(ctx, updatingProduct.ID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, domain.ErrProductNotFound
	}

	product.Name = updatingProduct.Name
	product.Description = updatingProduct.Description
	product.Price = updatingProduct.Price
	product.Stock = updatingProduct.Stock
	product.UpdatedAt = time.Now()

	updatedProduct, err := s.repo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	metrics.ServiceDuration.WithLabelValues("product-service.UpdateProduct").Observe(time.Since(start).Seconds())
	return updatedProduct, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id string) (bool, error) {
	start := time.Now()
	product, err := s.repo.Get(ctx, id)
	if err != nil {
		return false, err
	}
	if product == nil {
		return false, domain.ErrProductNotFound
	}

	ok, err := s.repo.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	metrics.ServiceDuration.WithLabelValues("product-service.DeleteProduct").Observe(time.Since(start).Seconds())
	return ok, nil
}
