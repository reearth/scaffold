package todouc

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type FindByIDs struct {
	TodoRepo      todo.Repo
	ProjectRepo   project.Repo
	WorkspaceRepo workspace.Repo
	TodoPolicy    todo.Policy
}

func (uc *FindByIDs) Execute(ctx context.Context, ids todo.IDList, user *user.User) (todo.List, error) {
	assets, err := uc.TodoRepo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	projects, err := uc.ProjectRepo.FindByIDs(ctx, assets.ProjectIDs())
	if err != nil {
		return nil, err
	}

	workspaces, err := uc.WorkspaceRepo.FindByIDs(ctx, projects.WorkspaceIDs())
	if err != nil {
		return nil, err
	}

	assets, err = uc.TodoPolicy.Filter(ctx, user, workspaces, projects, assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
