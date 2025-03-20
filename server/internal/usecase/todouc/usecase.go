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
	NewFindByIDs,
	NewFindByProject,
	NewCreate,
	NewUpdate,
	wire.Struct(new(Usecase), "*"),
)

type Usecase struct {
	FindByIDs     *FindByIDs
	FindByProject *FindByProject
	Create        *Create
	Update        *Update
}

type Builder struct {
	ctx       context.Context
	err       error
	user      *user.User
	todo      *todo.Todo
	project   *project.Project
	workspace *workspace.Workspace
}

func UsecaseBuilder(ctx context.Context, user *user.User) *Builder {
	return &Builder{ctx: ctx, user: user}
}

func (b *Builder) Result() (*todo.Todo, *project.Project, *workspace.Workspace, error) {
	if b.err != nil {
		return nil, nil, nil, b.err
	}
	return b.todo, b.project, b.workspace, b.err
}

func (b *Builder) FindTodoByID(id todo.ID, assetRepo todo.Repo) *Builder {
	if b.err != nil {
		return b
	}
	b.todo, b.err = assetRepo.FindByID(b.ctx, id)
	return b
}

func (b *Builder) FindProjectByID(id project.ID, projectRepo project.Repo, workspaceRepo workspace.Repo) *Builder {
	if b.err != nil {
		return b
	}
	b.project, b.err = projectRepo.FindByID(b.ctx, id)
	if b.err == nil {
		b.workspace, b.err = workspaceRepo.FindByID(b.ctx, b.project.Workspace())
	}
	return b
}

func (b *Builder) FindProjectByTodo(projectRepo project.Repo, workspaceRepo workspace.Repo) *Builder {
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

func (b *Builder) CanReadTodo(assetPolicy todo.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanRead(b.ctx, b.user, b.workspace, b.project, b.todo)
	return b
}

func (b *Builder) CanListTodo(assetPolicy todo.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanList(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *Builder) CanCreateTodo(assetPolicy todo.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanCreate(b.ctx, b.user, b.workspace, b.project)
	return b
}

func (b *Builder) CanUpdateTodo(assetPolicy todo.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanUpdate(b.ctx, b.user, b.workspace, b.project, b.todo)
	return b
}

func (b *Builder) CanDeleteTodo(assetPolicy todo.Policy) *Builder {
	if b.err != nil {
		return b
	}
	b.err = assetPolicy.CanDelete(b.ctx, b.user, b.workspace, b.project, b.todo)
	return b
}
