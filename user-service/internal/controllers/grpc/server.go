package server

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"github.com/wazwki/WearStore/user-service/internal/domain"
	"github.com/wazwki/WearStore/user-service/internal/services"
	"github.com/wazwki/WearStore/user-service/pkg/logger"
	"github.com/wazwki/WearStore/user-service/pkg/metrics"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
	service services.ServiceInterface
}

func NewServer(s services.ServiceInterface) userpb.UserServiceServer {
	return &Server{service: s}
}

func (server *Server) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	start := time.Now()
	id, err := server.service.Register(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		logger.Error("RegisterUser failed",
			zap.Error(err),
			zap.String("module", "user-service"),
			zap.String("email", req.GetEmail()))
		return nil, status.Errorf(codes.Unknown, "User not create")
	}

	logger.Info("User registered successfully",
		zap.String("module", "user-service"),
		zap.String("userID", id))
	metrics.ControllersDuration.WithLabelValues("user-service.RegisterUser", "/v1/users/register").Observe(time.Since(start).Seconds())
	return &userpb.RegisterUserResponse{Id: id}, nil
}

func (server *Server) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	start := time.Now()
	user, access, refresh, err := server.service.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		logger.Error("LoginUser failed",
			zap.Error(err),
			zap.String("module", "user-service"),
			zap.String("email", req.GetEmail()))
		return nil, status.Errorf(codes.Unknown, "User not login")
	}

	logger.Info("User logged in successfully",
		zap.String("module", "user-service"),
		zap.String("email", req.GetEmail()))

	header := metadata.Pairs("access_token", access, "refresh_token", refresh)
	err = grpc.SendHeader(ctx, header)
	if err != nil {
		logger.Error("LoginUser failed",
			zap.Error(err),
			zap.String("module", "user-service"),
			zap.String("email", req.GetEmail()))
		return nil, status.Errorf(codes.Unknown, "User not login")
	}

	metrics.ControllersDuration.WithLabelValues("user-service.LoginUser", "/v1/users/login").Observe(time.Since(start).Seconds())
	return &userpb.LoginUserResponse{User: domain.UserEntityToDTO(user)}, nil
}

func (server *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("access_token")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("access_token")[0]

	user, err := server.service.Get(ctx, req.GetId(), accessToken)
	if err != nil {
		logger.Error("GetUser failed",
			zap.Error(err),
			zap.String("module", "user-service"),
			zap.String("userID", req.GetId()))
		return nil, status.Errorf(codes.Unknown, "User not exist")
	}

	logger.Info("User retrieved successfully",
		zap.String("module", "user-service"),
		zap.String("userID", req.GetId()))
	metrics.ControllersDuration.WithLabelValues("user-service.GetUser", "/v1/users/{id}").Observe(time.Since(start).Seconds())
	return &userpb.GetUserResponse{User: domain.UserEntityToDTO(user)}, nil
}

func (server *Server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("access_token")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("access_token")[0]

	user, err := server.service.Update(ctx, req.GetId(), req.GetName(), req.GetEmail(), req.GetPassword(), accessToken)
	if err != nil {
		logger.Error("UpdateUser failed",
			zap.Error(err),
			zap.String("module", "user-service"),
			zap.String("userID", req.GetId()))
		return nil, status.Errorf(codes.Unknown, "User not exist")
	}

	logger.Info("User updated successfully",
		zap.String("module", "user-service"),
		zap.String("userID", req.GetId()))
	metrics.ControllersDuration.WithLabelValues("user-service.UpdateUser", "/v1/users/{id}").Observe(time.Since(start).Seconds())
	return &userpb.UpdateUserResponse{User: domain.UserEntityToDTO(user)}, nil
}

func (server *Server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("access_token")) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing or invalid access_token")
	}
	accessToken := md.Get("access_token")[0]

	ok, err := server.service.Delete(ctx, req.GetId(), accessToken)
	if err != nil {
		logger.Error("DeleteUser failed",
			zap.Error(err),
			zap.String("module", "user-service"),
			zap.String("userID", req.GetId()))
		return nil, status.Errorf(codes.Unknown, "User not delete")
	}

	logger.Info("User deleted successfully",
		zap.String("module", "user-service"),
		zap.String("userID", req.GetId()))
	metrics.ControllersDuration.WithLabelValues("user-service.DeleteUser", "/v1/users/{id}").Observe(time.Since(start).Seconds())
	return &userpb.DeleteUserResponse{Success: ok}, nil
}
