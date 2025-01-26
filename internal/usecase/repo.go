package usecase

import (
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

type Repos struct {
	User        user.Repo
	Workspace   workspace.Repo
	Project     project.Repo
	Asset       asset.Repo
	Transaction Transaction
}
