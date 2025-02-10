package assetuc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

type FindByProjectUsecase struct {
	assetRepo     asset.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo
	assetPolicy   asset.Policy
}

func NewFindByProjectUsecase(
	assetRepo asset.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy asset.Policy,
) *FindByProjectUsecase {
	return &FindByProjectUsecase{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *FindByProjectUsecase) Execute(ctx context.Context, pid project.ID, user *user.User) (asset.List, error) {
	_, project, _, err := UsecaseBuilder(ctx, user).
		FindProjectByID(pid, uc.projectRepo, uc.workspaceRepo).
		CanListAssets(uc.assetPolicy).
		Result()
	if err != nil {
		return nil, err
	}

	assets, err := uc.assetRepo.FindByProject(ctx, project.ID())
	if err != nil {
		return nil, err
	}

	return assets, nil
}
