package assetuc

import (
	"context"
	"errors"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/user"
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

func (uc *Usecase) Update(ctx context.Context, param UpdateParam, user *user.User) (*asset.Asset, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	asset, _, _, err := uc.Builder(ctx, user).
		FindAssetByID(param.ID).
		FindProjectByAsset().
		CanUpdateAsset().
		Result()
	if err != nil {
		return nil, err
	}

	if param.Name != nil {
		asset.SetName(*param.Name)
	}

	if err := uc.Repos.Asset.Save(ctx, asset); err != nil {
		return nil, err
	}

	return asset, nil
}
