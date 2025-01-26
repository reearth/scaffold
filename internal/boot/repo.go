package boot

import (
	"context"
	"fmt"

	imongo "github.com/reearth/server-scaffold/internal/infra/mongo"
	"github.com/reearth/server-scaffold/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitRepos(cfg *Config) (usecase.Repos, *mongo.Database) {
	ctx := context.Background()
	client, err := initMongo(ctx, cfg)
	if err != nil {
		panic(err)
	}

	db := client.Database(cfg.DB_APP)
	return usecase.Repos{
		Asset: imongo.NewAsset(db),
		// ...
	}, db
}

func initMongo(ctx context.Context, cfg *Config) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %v", err)
	}
	return client, nil
}
