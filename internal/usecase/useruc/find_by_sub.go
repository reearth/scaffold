package useruc

import (
	"context"

	"github.com/reearth/server-scaffold/pkg/user"
)

type FindBySubUsecase struct {
	userRepo user.Repo
}

func NewFindBySubUsecase(userRepo user.Repo) *FindBySubUsecase {
	return &FindBySubUsecase{
		userRepo: userRepo,
	}
}

func (uc *FindBySubUsecase) Execute(ctx context.Context, sub string) (*user.User, error) {
	return uc.userRepo.FindBySub(ctx, sub)
}
