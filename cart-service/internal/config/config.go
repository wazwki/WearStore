package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Level       string
	Host        string
	Port        string
	DBHost      string
	DBPort      string
	DBPassword  string
	DBNumber    int
	ProductAddr string
	UserAddr    string
}

func LoadFromEnv() (*Config, error) {
	dbnum, err := strconv.Atoi(os.Getenv("DB_NUMBER"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Level:       os.Getenv("LOG_LEVEL"),
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBNumber:    dbnum,
		ProductAddr: fmt.Sprintf("%s:%s", os.Getenv("PRODUCT_HOST"), os.Getenv("PRODUCT_PORT")),
		UserAddr:    fmt.Sprintf("%s:%s", os.Getenv("USER_HOST"), os.Getenv("USER_PORT")),
	}

	return cfg, nil
}
