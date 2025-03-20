package gqlmodel

import (
	"github.com/reearth/scaffold/server/pkg/project"
	"github.com/reearth/scaffold/server/pkg/todo"
)

func NewTodo(a *todo.Todo) *Todo {
	if a == nil {
		return nil
	}
	return &Todo{
		ID:        ID(a.ID()),
		Name:      a.Name(),
		ProjectID: ID(a.Project()),
	}
}

func NewAssets(assets todo.List) []*Todo {
	if assets == nil {
		return nil
	}
	res := make([]*Todo, 0, len(assets))
	for _, a := range assets {
		res = append(res, NewTodo(a))
	}
	return res
}

func (a *Todo) Into() (*todo.Todo, error) {
	return todo.New().
		ID(todo.ID(a.ID)).
		Name(a.Name).
		Project(project.ID(a.ProjectID)).
		Build()
}
