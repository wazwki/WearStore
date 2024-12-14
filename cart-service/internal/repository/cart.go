package repository

import "github.com/redis/go-redis/v9"

type RepositoryInterface interface{}

type Repository struct {
	DataBase *redis.Client
}

func NewRepository(db *redis.Client) RepositoryInterface {
	return &Repository{DataBase: db}
}
