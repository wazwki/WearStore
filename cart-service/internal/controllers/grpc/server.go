package server

import (
	"context"

	"github.com/wazwki/WearStore/cart-service/api/proto/cartpb"
	"github.com/wazwki/WearStore/cart-service/internal/services"
)

type Server struct {
	cartpb.UnimplementedCartServiceServer
	service services.ServiceInterface
}

func NewServer(s services.ServiceInterface) cartpb.CartServiceServer {
	return &Server{service: s}
}

func (s *Server) AddToCart(ctx context.Context, req *cartpb.AddToCartRequest) (*cartpb.AddToCartResponse, error) {

}

func (s *Server) RemoveFromCart(ctx context.Context, req *cartpb.RemoveFromCartRequest) (*cartpb.RemoveFromCartResponse, error) {

}

func (s *Server) GetCart(ctx context.Context, req *cartpb.GetCartRequest) (*cartpb.GetCartResponse, error) {

}

func (s *Server) ClearCart(ctx context.Context, req *cartpb.ClearCartRequest) (*cartpb.ClearCartResponse, error) {

}

func (s *Server) Checkout(ctx context.Context, req *cartpb.CheckoutRequest) (*cartpb.CheckoutResponse, error) {

}
