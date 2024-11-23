package app

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"github.com/wazwki/WearStore/user-service/db"
	"github.com/wazwki/WearStore/user-service/internal/config"
	"github.com/wazwki/WearStore/user-service/pkg/logger"
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
	slog.Info("Success logger init", slog.String("module", "user-service"))

	pool, err := db.ConnectPool(cfg.DBdsn)
	if err != nil {
		slog.Error("Fail connect pool", slog.Any("error", err), slog.String("module", "user-service"))
		return nil, err
	}

	repository := repository.NewRepository(pool)
	service := service.NewService(repository)
	serv := controllers_grpc.NewServer(service)

	grpcServer := grpc.NewServer()
	slog.Info("Success creating server", slog.String("module", "user-service"))

	userpb.RegisterUserServiceServer(grpcServer, serv)
	slog.Info("Success register service server", slog.String("module", "user-service"))

	return &App{pool: pool, server: grpcServer, migrateDSN: cfg.DBdsn, serverHost: cfg.Host, serverPort: cfg.Port}, nil
}

func (a *App) Run() error {
	slog.Info("Running app...", slog.String("module", "user-service"))

	if err := db.RunMigrations(a.migrateDSN); err != nil {
		slog.Error("Fail migrate", slog.Any("error", err), slog.String("module", "user-service"))
		return err
	}
	slog.Info("Migrate success", slog.String("module", "user-service"))

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", a.serverHost, a.serverPort))
	if err != nil {
		slog.Error("Fail listen", slog.Any("error", err), slog.String("module", "user-service"))
		return err
	}

	go func() {
		slog.Info(fmt.Sprintf("Server is running at %v", lis.Addr()), slog.String("module", "user-service"))
		if err := a.server.Serve(lis); err != nil {
			slog.Error("Fail start server", slog.Any("error", err), slog.String("module", "user-service"))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.pool.Close()
	a.server.GracefulStop()

	return nil
}
