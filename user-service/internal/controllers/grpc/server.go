package server

import (
	"context"

	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"github.com/wazwki/WearStore/user-service/internal/domain"
	"github.com/wazwki/WearStore/user-service/internal/services"
	"github.com/wazwki/WearStore/user-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
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
	return &userpb.RegisterUserResponse{Id: id}, nil
}

func (server *Server) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	user, err := server.service.Login(ctx, req.GetEmail(), req.GetPassword())
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
	return &userpb.LoginUserResponse{User: domain.UserEntityToDTO(user)}, nil
}

func (server *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := server.service.Get(ctx, req.GetId())
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
	return &userpb.GetUserResponse{User: domain.UserEntityToDTO(user)}, nil
}

func (server *Server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user, err := server.service.Update(ctx, req.GetId(), req.GetName(), req.GetEmail(), req.GetPassword())
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
	return &userpb.UpdateUserResponse{User: domain.UserEntityToDTO(user)}, nil
}

func (server *Server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	ok, err := server.service.Delete(ctx, req.GetId())
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
	return &userpb.DeleteUserResponse{Success: ok}, nil
}
