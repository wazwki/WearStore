package clients

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/auth-service/api/proto/authpb"
	"google.golang.org/grpc"
)

type AuthClient interface {
	CreateToken(ctx context.Context, id string) (string, string, error)
	CheckToken(ctx context.Context, token string) (bool, error)
}

type AuthClientImpl struct {
	client authpb.AuthServiceClient
}

func NewAuthClient(conn *grpc.ClientConn) AuthClient {
	return &AuthClientImpl{
		client: authpb.NewAuthServiceClient(conn),
	}
}

func (c *AuthClientImpl) CreateToken(ctx context.Context, id string) (string, string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	data, err := c.client.CreateToken(ctx, &authpb.CreateTokenRequest{UserID: id})
	if err != nil {
		return "", "", err
	}

	return data.AccessToken, data.RefreshToken, nil
}

func (c *AuthClientImpl) CheckToken(ctx context.Context, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	data, err := c.client.CheckToken(ctx, token)
	if err != nil {
		return false, err
	}

	return data.Valid, nil
}
