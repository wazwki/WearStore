package handlers

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/auth-service/api/proto/authpb"
	"github.com/wazwki/WearStore/auth-service/internal/service"
	"github.com/wazwki/WearStore/auth-service/pkg/metrics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
	service service.ServiceInterface
}

func NewControllers(s service.ServiceInterface) *Server {
	return &Server{service: s}
}

func (s *Server) CreateToken(ctx context.Context, req *authpb.CreateTokenRequest) (*authpb.CreateTokenResponse, error) {
	start := time.Now()
	id := req.GetUserID()
	access, refresh, err := s.service.CreateToken(ctx, id)
	if err != nil {
		metrics.ControllersDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
		return nil, status.Errorf(codes.Internal, "error creating token: %v", err)
	}
	metrics.ControllersDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
	return &authpb.CreateTokenResponse{AccessToken: access, RefreshToken: refresh}, nil
}

func (s *Server) CheckToken(ctx context.Context, req *authpb.CheckTokenRequest) (*authpb.CheckTokenResponse, error) {
	start := time.Now()
	token := req.GetToken()
	isValid, err := s.service.CheckToken(ctx, token)
	if err != nil {
		metrics.ControllersDuration.WithLabelValues("auth-service.CheckToken").Observe(time.Since(start).Seconds())
		return nil, status.Errorf(codes.Internal, "error checking token: %v", err)
	}
	metrics.ControllersDuration.WithLabelValues("auth-service.CheckToken").Observe(time.Since(start).Seconds())
	return &authpb.CheckTokenResponse{Valid: isValid}, nil
}

func (s *Server) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	start := time.Now()
	token := req.GetRefreshToken()
	access, err := s.service.RefreshToken(ctx, token)
	if err != nil {
		metrics.ControllersDuration.WithLabelValues("auth-service.RefreshToken").Observe(time.Since(start).Seconds())
		return nil, status.Errorf(codes.Internal, "error refreshing token: %v", err)
	}
	metrics.ControllersDuration.WithLabelValues("auth-service.RefreshToken").Observe(time.Since(start).Seconds())
	return &authpb.RefreshTokenResponse{AccessToken: access}, nil
}
