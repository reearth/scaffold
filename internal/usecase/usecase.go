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
