package services

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/product-service/internal/domain"
	"github.com/wazwki/WearStore/product-service/internal/repository"
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
	product, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, domain.ErrProductNotFound
	}
	return product, nil
}

func (s *Service) ListProducts(ctx context.Context, limit, offset int64) ([]*domain.Product, error) {
	products, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) AddProduct(ctx context.Context, newProduct *domain.Product) (string, error) {
	newProduct.CreatedAt = time.Now()
	newProduct.UpdatedAt = time.Now()
	id, err := s.repo.Create(ctx, newProduct)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *Service) UpdateProduct(ctx context.Context, updatingProduct *domain.Product) (*domain.Product, error) {
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
	return updatedProduct, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id string) (bool, error) {
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
	return ok, nil
}
