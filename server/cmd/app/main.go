package main

import (
	"context"
	"os"

	"github.com/reearth/scaffold/server/internal/di"
	"github.com/samber/lo"
)

func main() {
	ctx := context.Background()

	if len(os.Args) > 1 {
		cliApp := lo.Must(di.InitCLI(ctx, os.Args))
		cliApp.Must()
		return
	}

	server := lo.Must(di.InitEcho(ctx))
	lo.Must0(server.Start())
}
