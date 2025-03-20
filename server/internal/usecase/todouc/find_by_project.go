package todouc

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type FindByProject struct {
	TodoRepo      todo.Repo
	ProjectRepo   project.Repo
	WorkspaceRepo workspace.Repo
	TodoPolicy    todo.Policy
}

func (uc *FindByProject) Execute(ctx context.Context, pid project.ID, user *user.User) (todo.List, error) {
	_, project, _, err := build(ctx, user).
		FindProjectByID(pid, uc.ProjectRepo, uc.WorkspaceRepo).
		CanListTodo(uc.TodoPolicy).
		Result()
	if err != nil {
		return nil, err
	}

	assets, err := uc.TodoRepo.FindByProject(ctx, project.ID())
	if err != nil {
		return nil, err
	}

	return assets, nil
}
