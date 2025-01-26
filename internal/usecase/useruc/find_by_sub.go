package useruc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/user"
)

func (uc *Usecase) FindBySub(ctx context.Context, sub string) (*user.User, error) {
	return uc.Repos.User.FindBySub(ctx, sub)
}
