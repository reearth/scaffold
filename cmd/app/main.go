package main

import (
	"context"

	"github.com/reearth/server-scaffold/internal/boot"
	"github.com/reearth/server-scaffold/internal/boot/di/app"
)

func main() {
	cfg := boot.LoadConfig()
	cfg.Print()

	server, err := app.InitializeEcho(context.Background(), true)
	if err != nil {
		panic(err)
	}

	server.Start()

}
