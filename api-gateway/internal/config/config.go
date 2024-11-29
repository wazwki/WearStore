package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtutil "github.com/wazwki/WearStore/api-gateway/pkg/jwt"
)

type Config struct {
	Level  string
	JWTcfg jwtutil.Config
}

func LoadFromEnv() (*Config, error) {

	cfg := &Config{
		Level: os.Getenv("LOG_LEVEL"),
		JWTcfg: jwtutil.Config{
			AccessTokenSecret:  []byte("your-very-secret-key"),
			RefreshTokenSecret: []byte("your-very-secret-refresh-key"),
			AccessTokenTTL:     15 * time.Minute,
			RefreshTokenTTL:    7 * 24 * time.Hour,
			SigningMethod:      jwt.SigningMethodHS256,
		},
	}

	return cfg, nil
}
