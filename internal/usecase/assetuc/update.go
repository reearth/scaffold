package assetuc

import (
	"context"
	"errors"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
	"github.com/samber/lo"
)

type UpdateParam struct {
	ID   asset.ID
	Name *string
}

func (p UpdateParam) Validate() error {
	if lo.IsEmpty(p.ID) {
		return errors.New("id is required")
	}
	return nil
}

type UpdateUsecase struct {
	assetRepo     asset.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo

	assetPolicy asset.Policy
}

func NewUpdateUsecase(
	assetRepo asset.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy asset.Policy,
) *UpdateUsecase {
	return &UpdateUsecase{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *UpdateUsecase) Execute(ctx context.Context, param UpdateParam, user *user.User) (*asset.Asset, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	asset, _, _, err := UsecaseBuilder(ctx, user).
		FindAssetByID(param.ID, uc.assetRepo).
		FindProjectByAsset(uc.projectRepo, uc.workspaceRepo).
		CanUpdateAsset(uc.assetPolicy).
		Result()
	if err != nil {
		return nil, err
	}

	if param.Name != nil {
		asset.SetName(*param.Name)
	}

	if err := uc.assetRepo.Save(ctx, asset); err != nil {
		return nil, err
	}

	return asset, nil
}
