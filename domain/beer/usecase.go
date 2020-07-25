package beer

type UseCase interface {
	GetBeers(args BeerQueryArgs) ([]Beer, error)
}

type useCase struct {
	beerRpeo BeerRepo
}

func NewUseCase(beerRepo BeerRepo) UseCase {
	return &useCase{
		beerRpeo: beerRepo,
	}
}

func (u *useCase) GetBeers(args BeerQueryArgs) ([]Beer, error) {
	return u.beerRpeo.GetBeers(args)
}
