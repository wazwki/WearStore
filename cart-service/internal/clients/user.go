package clients

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/cart-service/internal/domain"
	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type UserClient interface {
	GetUser(ctx context.Context, user_id, token string) (*domain.User, error)
}

type UserClientImpl struct {
	client userpb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) UserClient {
	return &UserClientImpl{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (c *UserClientImpl) GetUser(ctx context.Context, user_id, token string) (*domain.User, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "access_token", token)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	res, err := c.client.GetUser(ctx, &userpb.GetUserRequest{Id: user_id})
	if err != nil {
		return &domain.User{}, err
	}

	return &domain.User{
		ID:        res.User.Id,
		Name:      res.User.Name,
		Email:     res.User.Email,
		Role:      res.User.Role,
		CreatedAt: res.User.CreatedAt,
		UpdatedAt: res.User.UpdatedAt,
	}, nil
}
