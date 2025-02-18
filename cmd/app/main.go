package main

import (
	"context"
	"os"

	"github.com/reearth/server-scaffold/internal/boot"
)

func main() {
	cfg := boot.LoadConfig()
	cfg.Print()

	if len(os.Args) > 1 {
		cliApp, err := boot.InitializeCLI(context.Background(), os.Args)
		if err != nil {
			panic(err)
		}
		cliApp.Must()
		return
	}

	server, err := boot.InitializeEcho(context.Background(), true)
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
