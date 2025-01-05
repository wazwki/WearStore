package config

import (
	"fmt"
	"os"
)

type Config struct {
	Level    string
	Host     string
	Port     string
	DBDSN    string
	HTTPPort string
	AuthAddr string
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Level: os.Getenv("LOG_LEVEL"),
		Host:  os.Getenv("HOST"),
		Port:  os.Getenv("PRODUCT_PORT"),
		DBDSN: fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
			os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB_NAME")),
		HTTPPort: os.Getenv("PRODUCT_HTTP_PORT"),
		AuthAddr: fmt.Sprintf("%s:%s", os.Getenv("AUTH_HOST"), os.Getenv("AUTH_PORT")),
	}

	return cfg, nil
}
