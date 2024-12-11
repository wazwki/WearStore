package services

import (
	"github.com/wazwki/WearStore/product-service/internal/repository"
)

type ServiceInterface interface{}

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &Service{repo: repo}
}
