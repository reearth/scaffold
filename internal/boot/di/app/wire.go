//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	"github.com/reearth/server-scaffold/internal/boot"
	"github.com/reearth/server-scaffold/internal/infra"
	"github.com/reearth/server-scaffold/internal/infra/gcp"
	"github.com/reearth/server-scaffold/internal/infra/mongo"
	"github.com/reearth/server-scaffold/internal/transport/echo"
	"github.com/reearth/server-scaffold/internal/usecase"
	"github.com/reearth/server-scaffold/internal/usecase/assetuc"
	"github.com/reearth/server-scaffold/internal/usecase/projectuc"
	"github.com/reearth/server-scaffold/internal/usecase/useruc"
	"github.com/reearth/server-scaffold/internal/usecase/workspaceuc"
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

func InitializeEcho(ctx context.Context, dev bool) (*echo.Server, error) {
	wire.Build(
		// boot
		boot.LoadConfig,
		boot.InitMongo,

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
		wire.Bind(new(infra.Storage), new(*gcp.Storage)),
		gcp.NewStorage,

		// policy
		asset.NewPolicy,

		// usecases
		assetuc.NewFindByIDsUsecase,
		assetuc.NewFindByProjectUsecase,
		assetuc.NewCreateUsecase,
		assetuc.NewUpdateUsecase,
		assetuc.New,

		projectuc.New,
		workspaceuc.New,

		useruc.NewFindBySubUsecase,
		useruc.New,

		usecase.NewUsecases,

		// echo
		echo.NewEchoConfig,
		echo.New,
	)

	return nil, nil
}
