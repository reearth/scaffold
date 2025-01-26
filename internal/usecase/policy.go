package usecase

import (
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

type Policies struct {
	Asset     asset.Policy
	Project   project.Policy
	Workspace workspace.Policy
}

func DefaultPolicies() Policies {
	return Policies{
		Asset:     asset.DefaultPolicy{},
		Project:   project.DefaultPolicy{},
		Workspace: workspace.DefaultPolicy{},
	}
}
