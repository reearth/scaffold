package useruc

type Usecase struct {
	FindBySub *FindBySub
}

func New(findBySub *FindBySub) *Usecase {
	return &Usecase{
		FindBySub: findBySub,
	}
}
