package assetuc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/project"
	"github.com/reearth/server-scaffold/pkg/user"
)

func (uc *Usecase) FindByProject(ctx context.Context, pid project.ID, user *user.User) (asset.List, error) {
	_, project, _, err := uc.Builder(ctx, user).
		FindProjectByID(pid).
		CanListAssets().
		Result()
	if err != nil {
		return nil, err
	}

	assets, err := uc.Repos.Asset.FindByProject(ctx, project.ID())
	if err != nil {
		return nil, err
	}

	return assets, nil
}
