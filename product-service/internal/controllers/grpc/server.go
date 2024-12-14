package server

import (
	"context"

	"github.com/wazwki/WearStore/product-service/api/proto/productpb"
	"github.com/wazwki/WearStore/product-service/internal/domain"
	"github.com/wazwki/WearStore/product-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	productpb.UnimplementedProductServiceServer
	service services.ServiceInterface
}

func NewServer(s services.ServiceInterface) productpb.ProductServiceServer {
	return &Server{service: s}
}

func (s *Server) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	product, err := s.service.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get product: %v", err)
	}
	if product == nil {
		return nil, status.Errorf(codes.NotFound, "product %s does not exist", req.GetId())
	}

	return &productpb.GetProductResponse{Product: domain.ProductEntityToDTO(product)}, nil
}
func (s *Server) ListProducts(ctx context.Context, req *productpb.ListProductsRequest) (*productpb.ListProductsResponse, error) {
	products, err := s.service.ListProducts(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not list products: %v", err)
	}

	var productList []*productpb.Product
	for _, product := range products {
		productList = append(productList, domain.ProductEntityToDTO(product))
	}

	return &productpb.ListProductsResponse{Products: productList}, nil
}
func (s *Server) AddProduct(ctx context.Context, req *productpb.AddProductRequest) (*productpb.AddProductResponse, error) {
	product := &domain.Product{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Stock:       req.GetStock(),
	}
	id, err := s.service.AddProduct(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not add product: %v", err)
	}

	return &productpb.AddProductResponse{Id: id}, nil

}
func (s *Server) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	product := &domain.Product{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Stock:       req.GetStock(),
	}
	updatedProduct, err := s.service.UpdateProduct(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update product: %v", err)
	}

	return &productpb.UpdateProductResponse{Product: domain.ProductEntityToDTO(updatedProduct)}, nil
}
func (s *Server) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	ok, err := s.service.DeleteProduct(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not delete product: %v", err)
	}

	return &productpb.DeleteProductResponse{Success: ok}, nil
}
