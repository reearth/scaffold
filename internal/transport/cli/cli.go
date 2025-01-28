package cli

import (
	"errors"

	"github.com/reearth/server-scaffold/internal/infra/mongo/mongomigrate"
	"github.com/reearth/server-scaffold/internal/transport"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Args     []string
	Usecases transport.Usecases
	Mongo    *mongo.Database
}

func Do(conf Config) error {
	var command string
	if len(conf.Args) > 1 {
		command = conf.Args[1]
	}

	if command == "migrate" {
		return migrate(conf.Mongo)
	}

	// TODO: add more commands for workers

	return errors.New("invalid command")
}

func Must(conf Config) {
	lo.Must0(Do(conf))
}

func migrate(mongo *mongo.Database) error {
	if err := mongomigrate.NewMigrator(mongo).Migrate1(); err != nil {
		return err
	}
	return nil
}
