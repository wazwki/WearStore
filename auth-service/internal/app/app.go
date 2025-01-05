package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	authpb "github.com/wazwki/WearStore/auth-service/api/proto/authpb"
	"github.com/wazwki/WearStore/auth-service/internal/config"
	controllers "github.com/wazwki/WearStore/auth-service/internal/controllers/grpc"
	"github.com/wazwki/WearStore/auth-service/internal/service"
	"github.com/wazwki/WearStore/auth-service/pkg/jwtutil"
	"github.com/wazwki/WearStore/auth-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	serverHost string
	serverPort string
	server     *grpc.Server
	httpServer *http.Server
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "auth-service"))

	jwt := jwtutil.NewJWTUtil(cfg.JWTcfg)

	service := service.NewService(jwt)
	controller := controllers.NewControllers(service)

	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServiceServer(grpcServer, controller)
	logger.Info("Success grpc server init", zap.String("module", "auth-service"))

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.Host, cfg.HTTPPort),
		Handler: mux,
	}

	return &App{server: grpcServer, serverHost: cfg.Host, serverPort: cfg.Port, httpServer: httpServer}, nil
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", a.serverHost, a.serverPort))
	if err != nil {
		logger.Error("Fail listen", zap.Error(err), zap.String("module", "auth-service"))
		return err
	}

	go func() {
		logger.Info(fmt.Sprintf("Server is running at %v", lis.Addr()), zap.String("module", "auth-service"))
		if err := a.server.Serve(lis); err != nil {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "auth-service"))
		}
	}()

	go func() {
		logger.Info(fmt.Sprintf("HTTP server is running at %v", a.httpServer.Addr), zap.String("module", "auth-service"))
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "auth-service"))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.server.GracefulStop()
	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}
