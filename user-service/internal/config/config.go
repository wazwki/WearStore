package config

import (
	"fmt"
	"os"
)

type Config struct {
	Level    string
	DBdsn    string
	Host     string
	Port     string
	HTTPPort string
	AuthAddr string
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Level: os.Getenv("LOG_LEVEL"),
		DBdsn: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("USER_DB_USER"), os.Getenv("USER_DB_PASSWORD"),
			os.Getenv("USER_DB_HOST"), os.Getenv("USER_DB_PORT"),
			os.Getenv("USER_DB_NAME")),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("USER_PORT"),
		HTTPPort: os.Getenv("USER_HTTP_PORT"),
		AuthAddr: fmt.Sprintf("%s:%s", os.Getenv("AUTH_HOST"), os.Getenv("AUTH_PORT")),
	}

	return cfg, nil
}
