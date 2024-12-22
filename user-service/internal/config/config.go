package config

import (
	"fmt"
	"os"
)

type Config struct {
	Level string
	DBdsn string
	Host  string
	Port  string
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Level: os.Getenv("LOG_LEVEL"),
		DBdsn: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("USER_DB_USER"), os.Getenv("USER_DB_PASSWORD"),
			os.Getenv("USER_DB_HOST"), os.Getenv("USER_DB_PORT"),
			os.Getenv("USER_DB_NAME")),
		Host: os.Getenv("HOST"),
		Port: os.Getenv("USER_PORT"),
	}

	return cfg, nil
}
