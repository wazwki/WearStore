package app

import (
	"context"
	"fmt"
	"net"

	"github.com/wazwki/WearStore/product-service/api/proto/productpb"
	"github.com/wazwki/WearStore/product-service/db"
	"github.com/wazwki/WearStore/product-service/internal/config"
	server "github.com/wazwki/WearStore/product-service/internal/controllers/grpc"
	"github.com/wazwki/WearStore/product-service/internal/repository"
	"github.com/wazwki/WearStore/product-service/internal/services"
	"github.com/wazwki/WearStore/product-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	server     *grpc.Server
	serverHost string
	serverPort string
	mongoctx   context.Context
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "product-service"))

	mongoctx := context.Background()
	client, err := db.Connect(mongoctx, cfg.DBDSN)
	if err != nil {
		logger.Error("Fail connect to mongodb", zap.Error(err), zap.String("module", "product-service"))
		return nil, err
	}
	collection := client.Database("wearstore").Collection("products")
	logger.Info("Success connect to mongodb", zap.String("module", "product-service"))

	repository := repository.NewRepository(collection)
	service := services.NewService(repository)
	srv := server.NewServer(service)

	grpcServer := grpc.NewServer()
	logger.Info("Success creating server", zap.String("module", "product-service"))

	productpb.RegisterProductServiceServer(grpcServer, srv)
	logger.Info("Success register service server", zap.String("module", "product-service"))

	return &App{server: grpcServer, serverHost: cfg.Host, serverPort: cfg.Port, mongoctx: mongoctx}, nil
}

func (a *App) Run() error {
	logger.Info("Running app...", zap.String("module", "product-service"))
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", a.serverHost, a.serverPort))
	if err != nil {
		logger.Error("Fail listen", zap.Error(err), zap.String("module", "product-service"))
		return err
	}

	go func() {
		logger.Info(fmt.Sprintf("Server is running at %v", lis.Addr()), zap.String("module", "product-service"))
		if err := a.server.Serve(lis); err != nil {
			logger.Error("Failed to start server", zap.Error(err), zap.String("module", "product-service"))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.server.GracefulStop()
	a.mongoctx.Done()

	return nil
}
