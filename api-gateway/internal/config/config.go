package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wazwki/WearStore/api-gateway/pkg/jwtutil"
)

type Config struct {
	Host   string
	Port   string
	Level  string
	JWTcfg jwtutil.Config
}

func LoadFromEnv() (*Config, error) {

	cfg := &Config{
		Host:  os.Getenv("HOST"),
		Port:  os.Getenv("GATEWAY_PORT"),
		Level: os.Getenv("LOG_LEVEL"),
		JWTcfg: jwtutil.Config{
			AccessTokenSecret:  []byte(os.Getenv("JWT_ACCESS_SECRET")),
			RefreshTokenSecret: []byte(os.Getenv("JWT_REFRESH_SECRET")),
			AccessTokenTTL:     15 * time.Minute,
			RefreshTokenTTL:    7 * 24 * time.Hour,
			SigningMethod:      jwt.SigningMethodHS256,
		},
	}

	return cfg, nil
}
