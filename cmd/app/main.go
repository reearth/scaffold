package main

import (
	"os"

	"github.com/reearth/server-scaffold/internal/boot"
	"github.com/reearth/server-scaffold/internal/transport/cli"
	"github.com/reearth/server-scaffold/internal/transport/echo"
)

func main() {
	cfg := boot.LoadConfig()
	cfg.Print()

	usecases, mongo := boot.InitUsecases(cfg)

	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if command != "" {
		cli.Must(cli.Config{
			Args:  os.Args,
			UC:    usecases,
			Mongo: mongo,
		})
		return
	}

	if err := echo.New(echo.Config{
		Port:     cfg.Port,
		Usecases: usecases,
	}).Start(); err != nil {
		panic(err)
	}
}
