package app

import (
	"fmt"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"github.com/wazwki/WearStore/user-service/db"
	"github.com/wazwki/WearStore/user-service/internal/config"
	server "github.com/wazwki/WearStore/user-service/internal/controllers/grpc"
	"github.com/wazwki/WearStore/user-service/internal/repository"
	"github.com/wazwki/WearStore/user-service/internal/services"
	"github.com/wazwki/WearStore/user-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	pool       *pgxpool.Pool
	server     *grpc.Server
	migrateDSN string
	serverHost string
	serverPort string
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "user-service"))

	pool, err := db.ConnectPool(cfg.DBdsn)
	if err != nil {
		logger.Error("Fail connect pool", zap.Error(err), zap.String("module", "user-service"))
		return nil, err
	}

	repository := repository.NewRepository(pool)
	service := services.NewService(repository)
	serv := server.NewServer(service)

	grpcServer := grpc.NewServer()
	logger.Info("Success creating server", zap.String("module", "user-service"))

	userpb.RegisterUserServiceServer(grpcServer, serv)
	logger.Info("Success register service server", zap.String("module", "user-service"))

	return &App{pool: pool, server: grpcServer, migrateDSN: cfg.DBdsn, serverHost: cfg.Host, serverPort: cfg.Port}, nil
}

func (a *App) Run() error {
	logger.Info("Running app...", zap.String("module", "user-service"))

	if err := db.RunMigrations(a.migrateDSN); err != nil {
		logger.Error("Fail migrate", zap.Error(err), zap.String("module", "user-service"))
		return err
	}
	logger.Info("Migrate success", zap.String("module", "user-service"))

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", a.serverHost, a.serverPort))
	if err != nil {
		logger.Error("Fail listen", zap.Error(err), zap.String("module", "user-service"))
		return err
	}

	go func() {
		logger.Info(fmt.Sprintf("Server is running at %v", lis.Addr()), zap.String("module", "user-service"))
		if err := a.server.Serve(lis); err != nil {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "user-service"))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.pool.Close()
	a.server.GracefulStop()

	return nil
}
