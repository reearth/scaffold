package usecase

import (
	"github.com/reearth/scaffold/server/internal/usecase/assetuc"
	"github.com/reearth/scaffold/server/internal/usecase/projectuc"
	"github.com/reearth/scaffold/server/internal/usecase/useruc"
	"github.com/reearth/scaffold/server/internal/usecase/workspaceuc"
)

type Usecases struct {
	Asset     *assetuc.Usecase
	Project   *projectuc.Usecase
	Workspace *workspaceuc.Usecase
	User      *useruc.Usecase
}
