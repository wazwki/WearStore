package clients

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/cart-service/internal/domain"
	"github.com/wazwki/WearStore/product-service/api/proto/productpb"
	"google.golang.org/grpc"
)

type ProductClient interface {
	GetProduct(ctx context.Context, productID string) (*domain.CartItem, error)
	UpdateProduct(ctx context.Context, product *domain.CartItem) (*domain.CartItem, error)
}

type ProductClientImpl struct {
	client productpb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) ProductClient {
	return &ProductClientImpl{
		client: productpb.NewProductServiceClient(conn),
	}
}

func (c *ProductClientImpl) GetProduct(ctx context.Context, productID string) (*domain.CartItem, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	res, err := c.client.GetProduct(ctx, &productpb.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}

	return &domain.CartItem{
		ProductID: res.Product.Id,
		Name:      res.Product.Name,
		Price:     res.Product.Price,
		Quantity:  int(res.Product.Stock),
	}, nil
}

func (c *ProductClientImpl) UpdateProduct(ctx context.Context, product *domain.CartItem) (*domain.CartItem, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	res, err := c.client.UpdateProduct(ctx, &productpb.UpdateProductRequest{
		Id:    product.ProductID,
		Name:  product.Name,
		Price: product.Price,
		Stock: int64(product.Quantity),
	})
	if err != nil {
		return nil, err
	}

	return &domain.CartItem{
		ProductID: res.Product.Id,
		Name:      res.Product.Name,
		Price:     res.Product.Price,
		Quantity:  int(res.Product.Stock),
	}, nil
}
