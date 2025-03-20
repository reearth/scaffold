//go:build wireinject

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/reearth/scaffold/server/internal/infra/gcp"
	"github.com/reearth/scaffold/server/internal/infra/mongo"
	"github.com/reearth/scaffold/server/internal/transport/cli"
	"github.com/reearth/scaffold/server/internal/transport/echo"
	"github.com/reearth/scaffold/server/internal/usecase"
	"github.com/reearth/scaffold/server/internal/usecase/gateway"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

func InitEcho(ctx context.Context) (*echo.Server, error) {
	wire.Build(
		// boot
		LoadConfig,
		InitMongo,

		// infra: mongo
		mongo.NewTodo,
		wire.Bind(new(todo.Repo), new(*mongo.Todo)),
		mongo.NewWorkspace,
		wire.Bind(new(workspace.Repo), new(*mongo.Workspace)),
		mongo.NewUser,
		wire.Bind(new(user.Repo), new(*mongo.User)),
		mongo.NewProject,
		wire.Bind(new(project.Repo), new(*mongo.Project)),

		// infra: storage
		wire.Bind(new(gateway.Storage), new(*gcp.Storage)),
		gcp.NewStorage,

		// policy
		todo.NewPolicy,

		// usecases
		usecase.Wire,

		// echo
		newEchoConfig,
		echo.New,
	)

	return nil, nil
}

func InitCLI(ctx context.Context, args []string) (*cli.CLI, error) {
	wire.Build(
		LoadConfig,
		InitMongo,
		cli.New,
	)

	return nil, nil
}
