package useruc

import "github.com/google/wire"

var Wire = wire.NewSet(
	wire.Struct(new(FindBySub), "*"),
	wire.Struct(new(Usecase), "*"),
)

type Usecase struct {
	FindBySub *FindBySub
}
