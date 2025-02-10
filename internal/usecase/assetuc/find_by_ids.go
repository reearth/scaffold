package assetuc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
)

type FindByIDsUsecase struct {
	assetRepo     asset.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo

	assetPolicy asset.Policy
}

func NewFindByIDsUsecase(
	assetRepo asset.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy asset.Policy,
) *FindByIDsUsecase {
	return &FindByIDsUsecase{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *FindByIDsUsecase) Execute(ctx context.Context, ids asset.IDList, user *user.User) (asset.List, error) {
	assets, err := uc.assetRepo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	projects, err := uc.projectRepo.FindByIDs(ctx, assets.ProjectIDs())
	if err != nil {
		return nil, err
	}

	workspaces, err := uc.workspaceRepo.FindByIDs(ctx, projects.WorkspaceIDs())
	if err != nil {
		return nil, err
	}

	assets, err = uc.assetPolicy.Filter(ctx, user, workspaces, projects, assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
