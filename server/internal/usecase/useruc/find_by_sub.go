package useruc

import (
	"context"

	"github.com/reearth/scaffold/server/pkg/user"
)

type FindBySub struct {
	UserRepo user.Repo
}

func (uc *FindBySub) Execute(ctx context.Context, sub string) (*user.User, error) {
	return uc.UserRepo.FindBySub(ctx, sub)
}
