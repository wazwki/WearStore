package server

import (
	"context"

	"github.com/wazwki/WearStore/cart-service/api/proto/cartpb"
	"github.com/wazwki/WearStore/cart-service/internal/domain"
	"github.com/wazwki/WearStore/cart-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cartpb.UnimplementedCartServiceServer
	service services.ServiceInterface
}

func NewServer(s services.ServiceInterface) cartpb.CartServiceServer {
	return &Server{service: s}
}

func (s *Server) AddToCart(ctx context.Context, req *cartpb.AddToCartRequest) (*cartpb.AddToCartResponse, error) {
	add, err := s.service.AddToCart(ctx, req.UserId, req.ProductId, int(req.Quantity))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error adding to cart: %v", err)
	}
	return &cartpb.AddToCartResponse{Success: add}, nil
}

func (s *Server) RemoveFromCart(ctx context.Context, req *cartpb.RemoveFromCartRequest) (*cartpb.RemoveFromCartResponse, error) {
	remove, err := s.service.RemoveFromCart(ctx, req.UserId, req.ProductId, int(req.Quantity))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error removing from cart: %v", err)
	}

	return &cartpb.RemoveFromCartResponse{Success: remove}, nil
}

func (s *Server) GetCart(ctx context.Context, req *cartpb.GetCartRequest) (*cartpb.GetCartResponse, error) {
	cart, err := s.service.GetCart(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting cart: %v", err)
	}

	cartdto := domain.CartEntityToDTO(cart)

	return &cartpb.GetCartResponse{Cart: cartdto}, nil
}

func (s *Server) ClearCart(ctx context.Context, req *cartpb.ClearCartRequest) (*cartpb.ClearCartResponse, error) {
	success, err := s.service.ClearCart(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error clearing cart: %v", err)
	}

	return &cartpb.ClearCartResponse{Success: success}, nil
}

func (s *Server) Checkout(ctx context.Context, req *cartpb.CheckoutRequest) (*cartpb.CheckoutResponse, error) {
	totalPrice, err := s.service.Checkout(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during checkout: %v", err)
	}

	return &cartpb.CheckoutResponse{TotalPrice: totalPrice}, nil
}
