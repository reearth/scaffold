//go:build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/reearth/server-scaffold/internal/infra/gcp"
	"github.com/reearth/server-scaffold/internal/infra/mongo"
	"github.com/reearth/server-scaffold/internal/transport/cli"
	"github.com/reearth/server-scaffold/internal/transport/echo"
	"github.com/reearth/server-scaffold/internal/usecase"
	"github.com/reearth/server-scaffold/internal/usecase/assetuc"
	"github.com/reearth/server-scaffold/internal/usecase/gateway"
	"github.com/reearth/server-scaffold/internal/usecase/projectuc"
	"github.com/reearth/server-scaffold/internal/usecase/useruc"
	"github.com/reearth/server-scaffold/internal/usecase/workspaceuc"
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

func InitEcho(ctx context.Context) (*echo.Server, error) {
	wire.Build(
		// boot
		LoadConfig,
		InitMongo,

		// infra: mongo
		mongo.NewAsset,
		wire.Bind(new(asset.Repo), new(*mongo.Asset)),
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
		asset.NewPolicy,

		// usecases
		assetuc.NewFindByIDs,
		assetuc.NewFindByProject,
		assetuc.NewCreate,
		assetuc.NewUpdate,
		assetuc.New,

		projectuc.New,
		workspaceuc.New,

		useruc.NewFindBySub,
		useruc.New,

		wire.Struct(new(usecase.Usecases), "*"),

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
