package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wazwki/WearStore/cart-service/api/proto/cartpb"
	"github.com/wazwki/WearStore/cart-service/db"
	"github.com/wazwki/WearStore/cart-service/internal/clients"
	"github.com/wazwki/WearStore/cart-service/internal/config"
	server "github.com/wazwki/WearStore/cart-service/internal/controllers/grpc"
	"github.com/wazwki/WearStore/cart-service/internal/repository"
	"github.com/wazwki/WearStore/cart-service/internal/services"
	"github.com/wazwki/WearStore/cart-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	server     *grpc.Server
	serverHost string
	serverPort string

	httpServer *http.Server
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "cart-service"))

	client, err := db.Config(cfg.DBHost, cfg.DBPort, cfg.DBPassword, cfg.DBNumber)
	if err != nil {
		logger.Error("Fail connect to redis", zap.Error(err), zap.String("module", "cart-service"))
		return nil, err
	}
	logger.Info("Success connect to redis", zap.String("module", "cart-service"))

	productConn, err := grpc.NewClient(cfg.ProductAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Fail connect to grpc", zap.Error(err), zap.String("module", "cart-service"))
		return nil, err
	}
	product := clients.NewProductClient(productConn)
	logger.Info("Success creating product client", zap.String("module", "cart-service"))

	userConn, err := grpc.NewClient(cfg.UserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Fail connect to grpc", zap.Error(err), zap.String("module", "cart-service"))
		return nil, err
	}
	user := clients.NewUserClient(userConn)
	logger.Info("Success creating user client", zap.String("module", "cart-service"))

	repository := repository.NewRepository(client)
	service := services.NewService(repository, product, user)
	serv := server.NewServer(service)

	grpcServer := grpc.NewServer()
	logger.Info("Success creating server", zap.String("module", "cart-service"))

	cartpb.RegisterCartServiceServer(grpcServer, serv)
	logger.Info("Success register service server", zap.String("module", "cart-service"))

	//http

	runtimeMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = cartpb.RegisterCartServiceHandlerFromEndpoint(context.Background(), runtimeMux, fmt.Sprintf("%v:%v", cfg.Host, cfg.Port), opts)
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

	return &App{server: grpcServer, serverHost: cfg.Host, serverPort: cfg.Port, httpServer: srv}, nil
}

func (a *App) Run() error {
	logger.Info("Running app...", zap.String("module", "cart-service"))
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", a.serverHost, a.serverPort))
	if err != nil {
		logger.Error("Fail listen", zap.Error(err), zap.String("module", "cart-service"))
		return err
	}

	go func() {
		logger.Info(fmt.Sprintf("Server is running at %v", lis.Addr()), zap.String("module", "cart-service"))
		if err := a.server.Serve(lis); err != nil {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "cart-service"))
		}
	}()

	go func() {
		logger.Info(fmt.Sprintf("HTTP server is running at %v", a.httpServer.Addr), zap.String("module", "cart-service"))
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "cart-service"))
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
