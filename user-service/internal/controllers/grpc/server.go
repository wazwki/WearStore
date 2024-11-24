package server

import (
	"context"
	"log/slog"

	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"github.com/wazwki/WearStore/user-service/internal/domain"
	"github.com/wazwki/WearStore/user-service/internal/services"
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
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	id, err := server.service.Register(ctx, name, email, password)
	if err != nil {
		slog.Error("RegisterUser", slog.Any("error", err), slog.String("module", "user-service"))
		return nil, status.Errorf(codes.Unknown, "User not create")
	}

	return &userpb.RegisterUserResponse{Id: id}, nil
}

func (server *Server) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	email := req.GetEmail()
	password := req.GetPassword()

	user, err := server.service.Login(ctx, email, password)
	if err != nil {
		slog.Error("LoginUser", slog.Any("error", err), slog.String("module", "user-service"))
		return nil, status.Errorf(codes.Unknown, "User not login")
	}

	return &userpb.LoginUserResponse{User: domain.EntityToDTO(user)}, nil
}

func (server *Server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id := req.GetId()

	user, err := server.service.Get(ctx, id)
	if err != nil {
		slog.Error("GetUser", slog.Any("error", err), slog.String("module", "user-service"))
		return nil, status.Errorf(codes.Unknown, "User not exist")
	}

	return &userpb.GetUserResponse{User: domain.EntityToDTO(user)}, nil
}

func (server *Server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	id := req.GetId()
	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	user, err := server.service.Update(ctx, id, name, email, password)
	if err != nil {
		slog.Error("UpdateUser", slog.Any("error", err), slog.String("module", "user-service"))
		return nil, status.Errorf(codes.Unknown, "User not exist")
	}

	return &userpb.UpdateUserResponse{User: domain.EntityToDTO(user)}, nil
}

func (server *Server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := req.GetId()

	ok, err := server.service.Delete(ctx, id)
	if err != nil {
		slog.Error("DeleteUser", slog.Any("error", err), slog.String("module", "user-service"))
		return nil, status.Errorf(codes.Unknown, "User not delete")
	}

	return &userpb.DeleteUserResponse{Success: ok}, nil
}
