package config

import (
	"os"
)

type Config struct {
	Level string
}

func LoadFromEnv() (*Config, error) {

	cfg := &Config{
		Level: os.Getenv("LOG_LEVEL"),
	}

	return cfg, nil
}
