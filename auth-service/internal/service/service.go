package service

import (
	"context"
	"time"

	"github.com/wazwki/WearStore/auth-service/pkg/jwtutil"
	"github.com/wazwki/WearStore/auth-service/pkg/metrics"
)

type ServiceInterface interface {
	CreateToken(ctx context.Context, id string) (string, string, error)
	CheckToken(ctx context.Context, token string) (bool, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type Service struct {
	jwt *jwtutil.JWTUtil
}

func NewService(jwt *jwtutil.JWTUtil) ServiceInterface {
	return &Service{jwt: jwt}
}

func (s *Service) CreateToken(ctx context.Context, id string) (string, string, error) {
	start := time.Now()
	accessToken, err := s.jwt.GenerateAccessToken(ctx, id)
	if err != nil {
		metrics.ServiceDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
		return "", "", err
	}
	refreshToken, err := s.jwt.GenerateRefreshToken(ctx, id)
	if err != nil {
		metrics.ServiceDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
		return "", "", err
	}
	metrics.ServiceDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
	return accessToken, refreshToken, nil
}

func (s *Service) CheckToken(ctx context.Context, token string) (bool, error) {
	start := time.Now()
	_, err := s.jwt.ValidateToken(ctx, token)
	if err != nil {
		metrics.ServiceDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
		return false, err
	}
	metrics.ServiceDuration.WithLabelValues("auth-service.CheckToken").Observe(time.Since(start).Seconds())
	return true, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	start := time.Now()
	access, err := s.jwt.RefreshAccessToken(ctx, refreshToken)
	if err != nil {
		metrics.ServiceDuration.WithLabelValues("auth-service.CreateToken").Observe(time.Since(start).Seconds())
		return "", err
	}
	metrics.ServiceDuration.WithLabelValues("auth-service.RefreshToken").Observe(time.Since(start).Seconds())
	return access, nil
}
