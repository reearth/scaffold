package assetuc

import (
	"context"
	"errors"
	"io"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
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

func (uc *Usecase) Create(ctx context.Context, param CreateParam, user *user.User) (*asset.Asset, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	_, project, _, err := uc.Builder(ctx, user).
		FindProjectByID(param.ProjectID).
		CanCreateAsset().
		Result()

	if err != nil {
		return nil, err
	}

	if err := uc.Gateways.Storage.Save(ctx, param.Name, param.Data); err != nil {
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

	if err := uc.Repos.Asset.Save(ctx, asset); err != nil {
		return nil, err
	}

	return asset, nil
}
