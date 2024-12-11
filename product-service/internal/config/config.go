package config

import (
	"fmt"
	"os"
)

type Config struct {
	Level string
	Host  string
	Port  string
	DBDSN string
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Level: os.Getenv("LOG_LEVEL"),
		Host:  os.Getenv("HOST"),
		Port:  os.Getenv("PORT"),
		DBDSN: fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")),
	}

	return cfg, nil
}
