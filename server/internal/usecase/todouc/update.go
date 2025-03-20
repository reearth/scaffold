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
	assetRepo     todo.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo
	assetPolicy   todo.Policy
}

func NewUpdate(
	assetRepo todo.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy todo.Policy,
) *Update {
	return &Update{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *Update) Execute(ctx context.Context, param UpdateParam, user *user.User) (*todo.Todo, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	asset, _, _, err := UsecaseBuilder(ctx, user).
		FindTodoByID(param.ID, uc.assetRepo).
		FindProjectByTodo(uc.projectRepo, uc.workspaceRepo).
		CanUpdateTodo(uc.assetPolicy).
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
