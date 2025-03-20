package todouc

import (
	"context"
	"errors"

	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
	"github.com/samber/lo"
)

type UpdateParam struct {
	ID   todo.ID
	Name *string
}

func (p UpdateParam) Validate() error {
	if lo.IsEmpty(p.ID) {
		return errors.New("id is required")
	}
	return nil
}

type Update struct {
	TodoRepo      todo.Repo
	ProjectRepo   project.Repo
	WorkspaceRepo workspace.Repo
	TodoPolicy    todo.Policy
}

func (uc *Update) Execute(ctx context.Context, param UpdateParam, user *user.User) (*todo.Todo, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	todo, _, _, err := build(ctx, user).
		FindTodoByID(param.ID, uc.TodoRepo).
		FindProjectByTodo(uc.ProjectRepo, uc.WorkspaceRepo).
		CanUpdateTodo(uc.TodoPolicy).
		Result()
	if err != nil {
		return nil, err
	}

	if param.Name != nil {
		todo.SetName(*param.Name)
	}

	if err := uc.TodoRepo.Save(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}
