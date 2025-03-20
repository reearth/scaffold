package todouc

import (
	"context"
	"errors"
	"io"

	"github.com/reearth/scaffold/server/internal/usecase/gateway"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
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

type Create struct {
	assetRepo     todo.Repo
	projectRepo   project.Repo
	workspaceRepo workspace.Repo
	assetPolicy   todo.Policy
	storage       gateway.Storage
}

func NewCreate(
	assetRepo todo.Repo,
	projectRepo project.Repo,
	workspaceRepo workspace.Repo,
	assetPolicy todo.Policy,
	storage gateway.Storage,
) *Create {
	return &Create{
		assetRepo:     assetRepo,
		projectRepo:   projectRepo,
		workspaceRepo: workspaceRepo,
		assetPolicy:   assetPolicy,
		storage:       storage,
	}
}

func (uc *Create) Execute(ctx context.Context, param CreateParam, user *user.User) (*todo.Todo, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}

	_, project, _, err := UsecaseBuilder(ctx, user).
		FindProjectByID(param.ProjectID, uc.projectRepo, uc.workspaceRepo).
		CanCreateTodo(uc.assetPolicy).
		Result()

	if err != nil {
		return nil, err
	}

	if err := uc.storage.Save(ctx, param.Name, param.Data); err != nil {
		return nil, err
	}

	asset, err := todo.New().
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
