package usecase

import (
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
	assetuc *assetuc.Usecase,
	projectuc *projectuc.Usecase,
	workspaceuc *workspaceuc.Usecase,
	useruc *useruc.Usecase,
) Usecases {
	return Usecases{
		Asset:     assetuc,
		Project:   projectuc,
		Workspace: workspaceuc,
		User:      useruc,
	}
}
