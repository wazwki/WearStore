package config

import (
	"os"
	"strconv"
)

type Config struct {
	Level      string
	Host       string
	Port       string
	DBHost     string
	DBPort     string
	DBPassword string
	DBNumber   int
}

func LoadFromEnv() (*Config, error) {
	dbnum, err := strconv.Atoi(os.Getenv("DB_NUMBER"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Level:      os.Getenv("LOG_LEVEL"),
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBNumber:   dbnum,
	}

	return cfg, nil
}
