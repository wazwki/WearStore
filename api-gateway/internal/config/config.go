package config

import (
	"os"
)

type Config struct {
	Host     string
	Port     string
	RESTPort string
	Level    string
}

func LoadFromEnv() (*Config, error) {

	cfg := &Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		RESTPort: os.Getenv("RESTPORT"),
		Level:    os.Getenv("LOG_LEVEL"),
	}

	return cfg, nil
}
