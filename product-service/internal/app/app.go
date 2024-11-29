package app

import (
	"log/slog"

	"github.com/wazwki/WearStore/product-service/internal/config"
	"github.com/wazwki/WearStore/product-service/pkg/logger"
)

type App struct{}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	slog.Info("Success logger init", slog.String("module", "product-service"))

	return &App{}, nil
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}
