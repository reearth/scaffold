package todouc

import (
	"context"
	"errors"

	"github.com/google/wire"
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
	"github.com/reearth/scaffold/server/pkg/user"
	"github.com/reearth/scaffold/server/pkg/workspace"
)

var Wire = wire.NewSet(
	wire.Struct(new(FindByIDs), "*"),
	wire.Struct(new(FindByProject), "*"),
	wire.Struct(new(Create), "*"),
	wire.Struct(new(Update), "*"),
	wire.Struct(new(Usecase), "*"),
)

type Usecase struct {
	FindByIDs     *FindByIDs
	FindByProject *FindByProject
	Create        *Create
	Update        *Update
}

type builder struct {
	ctx       context.Context
	err       error
	user      *user.User
	todo      *todo.Todo
	project   *project.Project
	workspace *workspace.Workspace
}

func build(ctx context.Context, user *user.User) *builder {
	return &builder{ctx: ctx, user: user}
}

func (b *builder) Result() (*todo.Todo, *project.Project, *workspace.Workspace, error) {
	if b.err != nil {
		return nil, nil, nil, b.err
	}
	return b.todo, b.project, b.workspace, b.err
}

func (b *builder) FindTodoByID(id todo.ID, assetRepo todo.Repo) *builder {
	if b.err != nil {
		return b
	}
	b.todo, b.err = assetRepo.FindByID(b.ctx, id)
	return b
}

func (b *builder) FindProjectByID(id project.ID, projectRepo project.Repo, workspaceRepo workspace.Repo) *builder {
	if b.err != nil {
		return b
	}
	b.project, b.err = projectRepo.FindByID(b.ctx, id)
	if b.err == nil {
		b.workspace, b.err = workspaceRepo.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *builder) FindProjectByTodo(projectRepo project.Repo, workspaceRepo workspace.Repo) *builder {
	if b.err != nil {
		return b
	}
	if b.todo == nil {
		b.err = errors.New("asset not found")
		return b
	}
	b.project, b.err = projectRepo.FindByID(b.ctx, b.todo.Project())
	if b.err == nil {
		b.workspace, b.err = workspaceRepo.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *builder) CanReadTodo(assetPolicy todo.Policy) *builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanRead(b.ctx, b.user, b.workspace, b.project, b.todo)
	return b
}

func (b *builder) CanListTodo(assetPolicy todo.Policy) *builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanList(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *builder) CanCreateTodo(assetPolicy todo.Policy) *builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanCreate(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *builder) CanUpdateTodo(assetPolicy todo.Policy) *builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanUpdate(b.ctx, b.user, b.workspace, b.project, b.todo)
	return b
}

func (b *builder) CanDeleteTodo(assetPolicy todo.Policy) *builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanDelete(b.ctx, b.user, b.workspace, b.project, b.todo)
	return b
}
