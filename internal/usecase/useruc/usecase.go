package useruc

import (
	"github.com/reearth/server-scaffold/internal/usecase"
)

type Usecase struct {
	usecase.Deps
}

func New(uc usecase.Deps) *Usecase {
	return &Usecase{
		Deps: uc,
	}
}
