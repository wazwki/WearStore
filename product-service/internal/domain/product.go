package domain

import (
	"errors"
	"time"

	"github.com/wazwki/WearStore/product-service/api/proto/productpb"
)

var ErrProductNotFound error = errors.New("Product not found")

type Product struct {
	ID          string    `bson:"_id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Price       float64   `bson:"price"`
	Stock       int64     `bson:"stock"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func ProductEntityToDTO(product *Product) *productpb.Product {
	return &productpb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.GoString(),
		UpdatedAt:   product.UpdatedAt.GoString(),
	}
}
