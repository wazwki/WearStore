package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wazwki/WearStore/notification-service/internal/config"
	"github.com/wazwki/WearStore/notification-service/internal/controllers/kafka"
	controllers "github.com/wazwki/WearStore/notification-service/internal/controllers/kafka/consumer"
	"github.com/wazwki/WearStore/notification-service/internal/service"
	"github.com/wazwki/WearStore/notification-service/pkg/logger"
	"go.uber.org/zap"
)

type App struct {
	server     *http.Server
	consumer   sarama.Consumer
	controller *controllers.Consumer
	topic      string
	partition  int32
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "notification-service"))

	consumer, err := kafka.NewConsumer(cfg.KafkaAddress)
	if err != nil {
		return nil, err
	}
	serv := service.NewService(cfg.SMTPConfig)
	con := controllers.NewConsumer(consumer, serv)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.Host, cfg.HTTPPort),
		Handler: mux,
	}

	return &App{server: srv, consumer: consumer, controller: con, topic: cfg.Topic, partition: cfg.Partition}, nil
}

func (a *App) Run() error {
	logger.Info("Running app...", zap.String("module", "notification-service"))

	go func() {
		logger.Info(fmt.Sprintf("HTTP server is running at %v", a.server.Addr), zap.String("module", "notification-service"))
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "notification-service"))
		}
	}()

	go func() {
		ticker := time.NewTicker(15 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			if err := a.controller.GetMessages(a.topic, a.partition); err != nil {
				logger.Error("Failed to get messages", zap.Error(err), zap.String("module", "notification-service"))
			}
		}
	}()

	return nil
}

func (a *App) Stop() error {
	logger.Info("Stopping app...", zap.String("module", "notification-service"))

	if err := a.server.Shutdown(context.Background()); err != nil {
		logger.Error("Failed to shutdown server", zap.Error(err), zap.String("module", "notification-service"))
	}

	if err := a.consumer.Close(); err != nil {
		logger.Error("Failed to close consumer", zap.Error(err), zap.String("module", "notification-service"))
	}

	return nil
}
