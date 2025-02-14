package assetuc

import (
	"context"
	"errors"
	"io"

	"github.com/reearth/server-scaffold/internal/usecase/gateway"
	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
	"github.com/reearth/server-scaffold/pkg/workspace"
	"github.com/samber/lo"
)

type CreateParam struct {
	ProjectID project.ID
	Name      string
	Data      io.Reader
}

func (p CreateParam) Validate() error {
	if lo.IsEmpty(p.ProjectID) {
		return errors.New("project id is required")
	}
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Data == nil {
		return errors.New("data is required")
	}
	return nil
}

type CreateUsecase struct {
	assetRepo     asset.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo
	assetPolicy   asset.Policy
	storage       gateway.Storage
}

func NewCreateUsecase(
	assetRepo asset.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy asset.Policy,
	storage gateway.Storage,
) *CreateUsecase {
	return &CreateUsecase{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
		storage:       storage,
	}
}

func (uc *CreateUsecase) Execute(ctx context.Context, param CreateParam, user *user.User) (*asset.Asset, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	_, project, _, err := UsecaseBuilder(ctx, user).
		FindProjectByID(param.ProjectID, uc.projectRepo, uc.workspaceRepo).
		CanCreateAsset(uc.assetPolicy).
		Result()

	if err != nil {
		return nil, err
	}

	if err := uc.storage.Save(ctx, param.Name, param.Data); err != nil {
		return nil, err
	}

	asset, err := asset.New().
		NewID().
		Project(project.ID()).
		Name(param.Name).
		Build()
	if err != nil {
		return nil, err
	}

	if err := uc.assetRepo.Save(ctx, asset); err != nil {
		return nil, err
	}

	return asset, nil
}
