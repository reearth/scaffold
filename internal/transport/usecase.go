package transport

import (
	"github.com/reearth/server-scaffold/internal/usecase"
	"github.com/reearth/server-scaffold/internal/usecase/assetuc"
	"github.com/reearth/server-scaffold/internal/usecase/projectuc"
	"github.com/reearth/server-scaffold/internal/usecase/useruc"
	"github.com/reearth/server-scaffold/internal/usecase/workspaceuc"
)

type Usecases struct {
	Asset     *assetuc.Usecase
	Project   *projectuc.Usecase
	Workspace *workspaceuc.Usecase
	User      *useruc.Usecase
}

func NewUsecases(
	uc usecase.Deps,
) Usecases {
	return Usecases{
		Asset:     assetuc.New(uc),
		Project:   projectuc.New(uc),
		Workspace: workspaceuc.New(uc),
		User:      useruc.New(uc),
	}
}
