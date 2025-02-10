package useruc

type Usecase struct {
	FindBySubUsecase *FindBySubUsecase
}

func New(findBySubUsecase *FindBySubUsecase) *Usecase {
	return &Usecase{
		FindBySubUsecase: findBySubUsecase,
	}
}
