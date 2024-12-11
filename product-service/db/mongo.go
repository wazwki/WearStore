package db

import (
	"context"
	"fmt"

	"github.com/wazwki/WearStore/product-service/pkg/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

func Connect(ctx context.Context, mongoDSN string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoDSN))
	if err != nil {
		return nil, fmt.Errorf("Connect mongo: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Ping mongo: %w", err)
	}

	go func() {
		<-ctx.Done()

		logger.Debug("Disconnecting mongo")

		if err := client.Disconnect(ctx); err != nil {
			logger.Warn("Failed to disconnect mongo", zap.Error(err), zap.String("module", "product-service"))
		}

		logger.Debug("Disconnected from mongo successfully", zap.String("module", "product-service"))
	}()

	return client, nil
}
