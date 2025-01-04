package services

import (
	"context"
	"errors"
	"time"

	"github.com/wazwki/WearStore/user-service/internal/domain"
	"github.com/wazwki/WearStore/user-service/internal/repository"
	"github.com/wazwki/WearStore/user-service/pkg/hasher"
	"github.com/wazwki/WearStore/user-service/pkg/metrics"
)

type ServiceInterface interface {
	Register(ctx context.Context, name, email, password string) (string, error)
	Login(ctx context.Context, email, password string) (*domain.User, error)
	Get(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, id, name, email, password string) (*domain.User, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{repo: repo}
}

func (service *Service) Register(ctx context.Context, name, email, password string) (string, error) {
	start := time.Now()
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &domain.User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := service.repo.Create(ctx, user)
	if err != nil {
		return "", err
	}

	metrics.ServiceDuration.WithLabelValues("user-service.Register").Observe(time.Since(start).Seconds())
	return id, nil
}

func (service *Service) Login(ctx context.Context, email, password string) (*domain.User, error) {
	start := time.Now()
	user, err := service.repo.FindByMail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !hasher.ComparePassword(user.Password, password) {
		return nil, errors.New("Incorrect password")
	}

	metrics.ServiceDuration.WithLabelValues("user-service.Login").Observe(time.Since(start).Seconds())
	return user, nil
}

func (service *Service) Get(ctx context.Context, id string) (*domain.User, error) {
	start := time.Now()
	user, err := service.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	metrics.ServiceDuration.WithLabelValues("user-service.Get").Observe(time.Since(start).Seconds())
	return user, nil
}

func (service *Service) Update(ctx context.Context, id, name, email, password string) (*domain.User, error) {
	start := time.Now()
	hashedPassword, err := hasher.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		UpdatedAt: time.Now(),
	}

	updatedUser, err := service.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	metrics.ServiceDuration.WithLabelValues("user-service.Update").Observe(time.Since(start).Seconds())
	return updatedUser, nil
}

func (service *Service) Delete(ctx context.Context, id string) (bool, error) {
	start := time.Now()
	err := service.repo.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	metrics.ServiceDuration.WithLabelValues("user-service.Delete").Observe(time.Since(start).Seconds())
	return true, nil
}
