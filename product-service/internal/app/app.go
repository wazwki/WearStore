package app

import (
	"github.com/wazwki/WearStore/product-service/internal/config"
	"github.com/wazwki/WearStore/product-service/pkg/logger"
	"go.uber.org/zap"
)

type App struct{}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "user-service"))

	return &App{}, nil
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}
