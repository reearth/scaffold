package workspaceuc

import "github.com/google/wire"

var Wire = wire.NewSet(
	wire.Struct(new(Usecase), "*"),
)

type Usecase struct{}
