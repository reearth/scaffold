package boot

import (
	"github.com/reearth/server-scaffold/internal/transport"
	"github.com/reearth/server-scaffold/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitUsecases(conf *Config) (transport.Usecases, *mongo.Database) {
	repos, mongo := InitRepos(conf)
	policies := InitPolicies(conf)
	gateways := InitGateways(conf)
	usecases := transport.NewUsecases(usecase.Deps{
		Repos:    repos,
		Policies: policies,
		Gateways: gateways,
	})
	return usecases, mongo
}
