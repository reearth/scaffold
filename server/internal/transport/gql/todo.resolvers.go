package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"context"
	"fmt"

	"github.com/reearth/scaffold/server/internal/transport/gql/gqlmodel"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input gqlmodel.CreateTodoInput) (*gqlmodel.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input gqlmodel.UpdateTodoInput) (*gqlmodel.Todo, error) {
	panic(fmt.Errorf("not implemented: UpdateTodo - updateTodo"))
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID gqlmodel.ID) (gqlmodel.ID, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, filter gqlmodel.TodoFilter) (*gqlmodel.TodoConnection, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Project is the resolver for the project field.
func (r *todoResolver) Project(ctx context.Context, obj *gqlmodel.Todo) (*gqlmodel.Project, error) {
	panic(fmt.Errorf("not implemented: Project - project"))
}

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
