package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wazwki/WearStore/api-gateway/api/proto/apigatewaypb"
	"github.com/wazwki/WearStore/api-gateway/internal/config"
	"github.com/wazwki/WearStore/api-gateway/internal/middleware"
	"github.com/wazwki/WearStore/api-gateway/pkg/logger"
	"go.uber.org/zap"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/wazwki/WearStore/api-gateway/pkg/jwtutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	mux            *runtime.ServeMux
	authMiddleware func(http.Handler) http.Handler
	Host           string
	Port           string
	serverContext  context.Context
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.Level)
	logger.Info("Success logger init", zap.String("module", "api-gateway"))

	jwt := jwtutil.NewJWTUtil(cfg.JWTcfg)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := apigatewaypb.RegisterApiGatewayServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		logger.Fatal("Failed to start HTTP Gateway", zap.Error(err))
	}

	protectedRoutes := []middleware.ProtectedRoute{
		{Path: "/v1/users/{id}", Method: http.MethodDelete},
		{Path: "/v1/users/{id}", Method: http.MethodPut},
		{Path: "/v1/users/{id}", Method: http.MethodGet},
		{Path: "/v1/products", Method: http.MethodPost},
		{Path: "/v1/products/{id}", Method: http.MethodPut},
		{Path: "/v1/products/{id}", Method: http.MethodDelete},
		{Path: "/v1/carts/add", Method: http.MethodPost},
		{Path: "/v1/cart/remove", Method: http.MethodPost},
		{Path: "/v1/cart/{user_id}", Method: http.MethodGet},
		{Path: "/v1/cart/clear", Method: http.MethodPost},
		{Path: "/v1/cart/checkout", Method: http.MethodPost},
	}

	authMiddleware := middleware.AuthMiddleware(jwt, protectedRoutes)

	return &App{mux: mux, authMiddleware: authMiddleware, Host: cfg.Host, Port: cfg.RESTPort, serverContext: ctx}, nil
}

func (a *App) Run() error {
	go func() {
		logger.Info(fmt.Sprintf("Starting HTTP server on %v", a.Host), zap.String("module", "api-gateway"))
		if err := http.ListenAndServe(fmt.Sprintf("%v:%v", a.Host, a.Port), a.authMiddleware(a.mux)); err != nil {
			logger.Fatal("Failed to start HTTP server", zap.Error(err))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	a.serverContext.Done()
	if err := a.serverContext.Err(); err != nil {
		return err
	}

	return nil
}
