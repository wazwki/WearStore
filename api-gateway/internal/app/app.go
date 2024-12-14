package app

import (
	"github.com/wazwki/WearStore/api-gateway/internal/config"
	"github.com/wazwki/WearStore/api-gateway/pkg/jwtutil"
	"github.com/wazwki/WearStore/api-gateway/pkg/logger"
	"go.uber.org/zap"
)

type App struct {
	JWTservice *jwtutil.JWTUtil
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "api-gateway"))

	jwt := jwtutil.NewJWTUtil(cfg.JWTcfg)

	return &App{JWTservice: jwt}, nil
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop() error {
	return nil
}
