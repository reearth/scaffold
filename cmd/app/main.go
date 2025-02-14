package main

import (
	"context"
	"os"

	"github.com/reearth/server-scaffold/internal/boot"
	"github.com/reearth/server-scaffold/internal/boot/di"
)

func main() {
	cfg := boot.LoadConfig()
	cfg.Print()

	if len(os.Args) > 1 {
		cliApp, err := di.InitializeCLI(context.Background(), os.Args)
		if err != nil {
			panic(err)
		}
		cliApp.Must()
		return
	}

	server, err := di.InitializeEcho(context.Background(), true)
	if err != nil {
		panic(err)
	}

	server.Start()

}
