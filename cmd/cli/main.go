package main

import (
	"context"
	"os"

	"github.com/reearth/server-scaffold/internal/boot"
	"github.com/reearth/server-scaffold/internal/boot/di/cliapp"
	"github.com/reearth/server-scaffold/internal/transport/cli"
)

func main() {
	cfg := boot.LoadConfig()
	cfg.Print()

	mongo, err := boot.InitMongo(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	cliApp, err := cliapp.InitializeCLI(context.Background(), cli.Config{
		Args:  os.Args,
		Mongo: mongo,
	})
	if err != nil {
		panic(err)
	}

	cliApp.Must()
}
