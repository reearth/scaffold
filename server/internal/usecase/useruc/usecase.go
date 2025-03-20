package useruc

import "github.com/google/wire"

var Wire = wire.NewSet(
	NewFindBySub,
	wire.Struct(new(Usecase), "*"),
)

type Usecase struct {
	FindBySub *FindBySub
}
