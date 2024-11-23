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
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")),
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	return cfg, nil
}
