package useruc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/user"
)

type FindBySub struct {
	userRepo user.Repo
}

func NewFindBySub(userRepo user.Repo) *FindBySub {
	return &FindBySub{
		userRepo: userRepo,
	}
}

func (uc *FindBySub) Execute(ctx context.Context, sub string) (*user.User, error) {
	return uc.userRepo.FindBySub(ctx, sub)
}
