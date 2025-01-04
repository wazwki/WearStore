package repository

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/product-service/internal/domain"
	"github.com/wazwki/WearStore/product-service/pkg/metrics"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type RepositoryInterface interface {
	Get(ctx context.Context, id string) (*domain.Product, error)
	List(ctx context.Context, limit, offset int64) ([]*domain.Product, error)
	Create(ctx context.Context, newProduct *domain.Product) (string, error)
	Update(ctx context.Context, updatingProduct *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type Repository struct {
	DataBase *mongo.Collection
}

func NewRepository(db *mongo.Collection) RepositoryInterface {
	return &Repository{DataBase: db}
}

func (s *Repository) Get(ctx context.Context, id string) (*domain.Product, error) {
	start := time.Now()
	var product *domain.Product
	filter := bson.M{"product_id": id}
	err := s.DataBase.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}

	metrics.RepositoryDuration.WithLabelValues("product-service.Get").Observe(time.Since(start).Seconds())
	return product, nil
}

func (s *Repository) List(ctx context.Context, limit, offset int64) ([]*domain.Product, error) {
	start := time.Now()
	var products []*domain.Product
	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := s.DataBase.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	metrics.RepositoryDuration.WithLabelValues("product-service.List").Observe(time.Since(start).Seconds())
	return products, nil
}

func (s *Repository) Create(ctx context.Context, newProduct *domain.Product) (string, error) {
	start := time.Now()
	_, err := s.DataBase.InsertOne(ctx, newProduct)
	if err != nil {
		return "", err
	}

	metrics.RepositoryDuration.WithLabelValues("product-service.Create").Observe(time.Since(start).Seconds())
	return newProduct.ID, nil
}

func (s *Repository) Update(ctx context.Context, updatingProduct *domain.Product) (*domain.Product, error) {
	start := time.Now()
	filter := bson.M{"product_id": updatingProduct.ID}
	update := bson.M{
		"$set": bson.M{
			"name":        updatingProduct.Name,
			"description": updatingProduct.Description,
			"price":       updatingProduct.Price,
			"stock":       updatingProduct.Stock,
			"updated_at":  updatingProduct.UpdatedAt,
		},
	}
	result, err := s.DataBase.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, domain.ErrProductNotFound
	}

	metrics.RepositoryDuration.WithLabelValues("product-service.Update").Observe(time.Since(start).Seconds())
	return updatingProduct, nil
}

func (s *Repository) Delete(ctx context.Context, id string) (bool, error) {
	start := time.Now()
	filter := bson.M{"product_id": id}
	result, err := s.DataBase.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, domain.ErrProductNotFound
	}

	metrics.RepositoryDuration.WithLabelValues("product-service.Delete").Observe(time.Since(start).Seconds())
	return true, nil
}
