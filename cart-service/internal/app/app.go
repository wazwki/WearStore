package app

import (
	"fmt"
	"net"

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
)

type App struct {
	server     *grpc.Server
	serverHost string
	serverPort string
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

	productConn, err := grpc.NewClient(cfg.ProductAddr)
	if err != nil {
		logger.Error("Fail connect to grpc", zap.Error(err), zap.String("module", "cart-service"))
		return nil, err
	}
	product := clients.NewProductClient(productConn)
	logger.Info("Success creating product client", zap.String("module", "cart-service"))

	userConn, err := grpc.NewClient(cfg.UserAddr)
	if err != nil {
		logger.Error("Fail connect to grpc", zap.Error(err), zap.String("module", "cart-service"))
		return nil, err
	}
	user := clients.NewUserClient(userConn)
	logger.Info("Success creating user client", zap.String("module", "cart-service"))

	repository := repository.NewRepository(client)
	service := services.NewService(repository, product, user)
	srv := server.NewServer(service)

	grpcServer := grpc.NewServer()
	logger.Info("Success creating server", zap.String("module", "cart-service"))

	cartpb.RegisterCartServiceServer(grpcServer, srv)
	logger.Info("Success register service server", zap.String("module", "cart-service"))

	return &App{server: grpcServer, serverHost: cfg.Host, serverPort: cfg.Port}, nil
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

	return nil
}

func (a *App) Stop() error {
	a.server.GracefulStop()

	return nil
}
