package assetuc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/asset"
	"github.com/reearth/server-scaffold/pkg/user"
)

func (uc *Usecase) FindByIDs(ctx context.Context, ids asset.IDList, user *user.User) (asset.List, error) {
	assets, err := uc.Repos.Asset.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	projects, err := uc.Repos.Project.FindByIDs(ctx, assets.ProjectIDs())
	if err != nil {
		return nil, err
	}

	workspaces, err := uc.Repos.Workspace.FindByIDs(ctx, projects.WorkspaceIDs())
	if err != nil {
		return nil, err
	}

	assets, err = uc.Policies.Asset.Filter(ctx, user, workspaces, projects, assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
