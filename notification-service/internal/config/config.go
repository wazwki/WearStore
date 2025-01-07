package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Level        string
	Host         string
	HTTPPort     string
	KafkaAddress string
	Topic        string
	Partition    int32
	SMTPConfig   map[string]string
}

func LoadFromEnv() (*Config, error) {
	partition, err := strconv.Atoi(os.Getenv("NOTIFICATION_PARTITION"))
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		Level:        os.Getenv("LOG_LEVEL"),
		Host:         os.Getenv("HOST"),
		HTTPPort:     os.Getenv("NOTIFICATION_HTTP_PORT"),
		KafkaAddress: fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		Topic:        os.Getenv("NOTIFICATION_TOPIC"),
		Partition:    int32(partition),
		SMTPConfig: map[string]string{
			"host":     os.Getenv("SMTP_HOST"),
			"port":     os.Getenv("SMTP_PORT"),
			"username": os.Getenv("SMTP_USERNAME"),
			"password": os.Getenv("SMTP_PASSWORD"),
			"from":     os.Getenv("SMTP_FROM"),
		},
	}

	return cfg, nil
}
