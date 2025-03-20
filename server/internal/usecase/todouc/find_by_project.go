package todouc

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

type FindByProject struct {
	assetRepo     todo.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo
	assetPolicy   todo.Policy
}

func NewFindByProject(
	assetRepo todo.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy todo.Policy,
) *FindByProject {
	return &FindByProject{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
	}
}

func (uc *FindByProject) Execute(ctx context.Context, pid project.ID, user *user.User) (todo.List, error) {
	_, project, _, err := UsecaseBuilder(ctx, user).
		FindProjectByID(pid, uc.projectRepo, uc.workspaceRepo).
		CanListTodo(uc.assetPolicy).
		Result()
	if err != nil {
		return nil, err
	}

	assets, err := uc.assetRepo.FindByProject(ctx, project.ID())
	if err != nil {
		return nil, err
	}

	return assets, nil
}
