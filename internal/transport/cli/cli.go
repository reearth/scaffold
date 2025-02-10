package cli

import (
	"errors"

	"github.com/reearth/server-scaffold/internal/infra/mongo/mongomigrate"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Args  []string
	Mongo *mongo.Database
}

func (c *CLI) Do() error {
	var command string
	if len(c.conf.Args) > 1 {
		command = c.conf.Args[1]
	}

	if command == "migrate" {
		return migrate(c.conf.Mongo)
	}

	// TODO: add more commands for workers

	return errors.New("invalid command")
}

type CLI struct {
	conf Config
}

func NewCLI(conf Config) *CLI {
	return &CLI{
		conf: conf,
	}
}

func (c *CLI) Must() {
	lo.Must0(c.Do())
}

func migrate(mongo *mongo.Database) error {
	if err := mongomigrate.NewMigrator(mongo).Migrate1(); err != nil {
		return err
	}
	return nil
}
