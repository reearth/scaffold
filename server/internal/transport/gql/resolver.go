package gql

import (
	"context"

	"github.com/reearth/scaffold/server/internal/usecase"
	"github.com/reearth/scaffold/server/pkg/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	uc usecase.Usecases
}

func NewResolver(uc usecase.Usecases) *Resolver {
	return &Resolver{uc: uc}
}

type UserKey struct{}

func (r *Resolver) user(ctx context.Context) *user.User {
	u, _ := ctx.Value(UserKey{}).(*user.User)
	return u
}
