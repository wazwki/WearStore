package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type RepositoryInterface interface{}

type Repository struct {
	DataBase *mongo.Collection
}

func NewRepository(db *mongo.Collection) RepositoryInterface {
	return &Repository{DataBase: db}
}
