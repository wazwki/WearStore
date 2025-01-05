package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wazwki/WearStore/auth-service/pkg/jwtutil"
)

type Config struct {
	Level    string
	Host     string
	Port     string
	HTTPPort string
	JWTcfg   jwtutil.Config
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Level:    os.Getenv("LOG_LEVEL"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("AUTH_PORT"),
		HTTPPort: os.Getenv("AUTH_HTTP_PORT"),
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
