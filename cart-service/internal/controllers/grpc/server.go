package server

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/cart-service/api/proto/cartpb"
	"github.com/wazwki/WearStore/cart-service/internal/domain"
	"github.com/wazwki/WearStore/cart-service/internal/services"
	"github.com/wazwki/WearStore/cart-service/pkg/metrics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("authorization")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("authorization")[0]

	add, err := s.service.AddToCart(ctx, req.GetUserId(), req.GetProductId(), int(req.GetQuantity()), accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error adding to cart: %v", err)
	}

	metrics.ControllersDuration.WithLabelValues("cart-service.AddToCart", "/v1/cart/add").Observe(time.Since(start).Seconds())
	return &cartpb.AddToCartResponse{Success: add}, nil
}

func (s *Server) RemoveFromCart(ctx context.Context, req *cartpb.RemoveFromCartRequest) (*cartpb.RemoveFromCartResponse, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("authorization")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("authorization")[0]

	remove, err := s.service.RemoveFromCart(ctx, req.GetUserId(), req.GetProductId(), int(req.GetQuantity()), accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error removing from cart: %v", err)
	}

	metrics.ControllersDuration.WithLabelValues("cart-service.RemoveFromCart", "/v1/cart/remove").Observe(time.Since(start).Seconds())
	return &cartpb.RemoveFromCartResponse{Success: remove}, nil
}

func (s *Server) GetCart(ctx context.Context, req *cartpb.GetCartRequest) (*cartpb.GetCartResponse, error) {
	start := time.Now()
	cart, err := s.service.GetCart(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting cart: %v", err)
	}

	cartdto := domain.CartEntityToDTO(cart)

	metrics.ControllersDuration.WithLabelValues("cart-service.GetCart", "/v1/cart/{user_id}").Observe(time.Since(start).Seconds())
	return &cartpb.GetCartResponse{Cart: cartdto}, nil
}

func (s *Server) ClearCart(ctx context.Context, req *cartpb.ClearCartRequest) (*cartpb.ClearCartResponse, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("authorization")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("authorization")[0]

	success, err := s.service.ClearCart(ctx, req.GetUserId(), accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error clearing cart: %v", err)
	}

	metrics.ControllersDuration.WithLabelValues("cart-service.ClearCart", "/v1/cart/clear").Observe(time.Since(start).Seconds())
	return &cartpb.ClearCartResponse{Success: success}, nil
}

func (s *Server) Checkout(ctx context.Context, req *cartpb.CheckoutRequest) (*cartpb.CheckoutResponse, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("authorization")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("authorization")[0]

	totalPrice, err := s.service.Checkout(ctx, req.GetUserId(), accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error during checkout: %v", err)
	}

	metrics.ControllersDuration.WithLabelValues("cart-service.Checkout", "/v1/cart/checkout").Observe(time.Since(start).Seconds())
	return &cartpb.CheckoutResponse{TotalPrice: totalPrice}, nil
}
