//go:build wireinject
// +build wireinject

package cliapp

import (
	"context"

	"github.com/google/wire"
	"github.com/reearth/server-scaffold/internal/transport/cli"
)

func InitializeCLI(ctx context.Context, conf cli.Config) (*cli.CLI, error) {
	wire.Build(
		cli.NewCLI,
	)

	return nil, nil
}
