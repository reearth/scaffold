package usecase

import (
	"github.com/google/wire"
	"github.com/reearth/scaffold/server/internal/usecase/projectuc"
	"github.com/reearth/scaffold/server/internal/usecase/todouc"
	"github.com/reearth/scaffold/server/internal/usecase/useruc"
	"github.com/reearth/scaffold/server/internal/usecase/workspaceuc"
)

var Wire = wire.NewSet(
	projectuc.Wire,
	todouc.Wire,
	useruc.Wire,
	workspaceuc.Wire,
	wire.Struct(new(Usecases), "*"),
)

type Usecases struct {
	Todo      *todouc.Usecase
	Project   *projectuc.Usecase
	Workspace *workspaceuc.Usecase
	User      *useruc.Usecase
}
