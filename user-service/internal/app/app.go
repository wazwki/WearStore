package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wazwki/WearStore/user-service/api/proto/userpb"
	"github.com/wazwki/WearStore/user-service/db"
	"github.com/wazwki/WearStore/user-service/internal/clients"
	"github.com/wazwki/WearStore/user-service/internal/config"
	server "github.com/wazwki/WearStore/user-service/internal/controllers/grpc"
	"github.com/wazwki/WearStore/user-service/internal/repository"
	"github.com/wazwki/WearStore/user-service/internal/services"
	"github.com/wazwki/WearStore/user-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	pool       *pgxpool.Pool
	server     *grpc.Server
	migrateDSN string
	serverHost string
	serverPort string
	httpServer *http.Server
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "user-service"))

	pool, err := db.ConnectPool(cfg.DBdsn)
	if err != nil {
		logger.Error("Fail connect pool", zap.Error(err), zap.String("module", "user-service"))
		return nil, err
	}

	authConn, err := grpc.NewClient(cfg.AuthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Fail connect to grpc", zap.Error(err), zap.String("module", "user-service"))
		return nil, err
	}
	auth := clients.NewAuthClient(authConn)
	logger.Info("Success creating user client", zap.String("module", "user-service"))

	repository := repository.NewRepository(pool)
	service := services.NewService(repository, auth)
	serv := server.NewServer(service)

	grpcServer := grpc.NewServer()
	logger.Info("Success creating server", zap.String("module", "user-service"))

	userpb.RegisterUserServiceServer(grpcServer, serv)
	logger.Info("Success register service server", zap.String("module", "user-service"))

	//http

	runtimeMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = userpb.RegisterUserServiceHandlerFromEndpoint(context.Background(), runtimeMux, fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), opts)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	mux.Handle("/", runtimeMux)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.Host, cfg.HTTPPort),
		Handler: mux,
	}

	return &App{pool: pool, server: grpcServer, migrateDSN: cfg.DBdsn, serverHost: cfg.Host, serverPort: cfg.Port, httpServer: srv}, nil
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

	go func() {
		logger.Info(fmt.Sprintf("HTTP server is running at %v", a.httpServer.Addr), zap.String("module", "user-service"))
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "user-service"))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.pool.Close()
	a.server.GracefulStop()
	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}
