package boot

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(ctx context.Context, cfg *Config) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %v", err)
	}

	return client.Database(cfg.DB_APP), nil
}
