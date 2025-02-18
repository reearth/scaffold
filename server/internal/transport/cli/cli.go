package cli

import (
	"errors"

	"github.com/reearth/scaffold/server/internal/infra/mongo/mongomigrate"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/mongo"
)

type CLI struct {
	args  []string
	mongo *mongo.Database
}

func New(args []string, mongo *mongo.Database) *CLI {
	return &CLI{
		args:  args,
		mongo: mongo,
	}
}

func (c *CLI) Must() {
	lo.Must0(c.Do())
}

func (c *CLI) Do() error {
	var command string
	if len(c.args) > 1 {
		command = c.args[1]
	}

	if command == "migrate" {
		return migrate(c.mongo)
	}

	// TODO: add more commands for workers

	return errors.New("invalid command")
}

func migrate(mongo *mongo.Database) error {
	if err := mongomigrate.NewMigrator(mongo).Migrate1(); err != nil {
		return err
	}
	return nil
}
