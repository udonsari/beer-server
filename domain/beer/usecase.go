package beer

type UseCase interface {
	// TODO Add arguments
	GetBeers() ([]Beer, error)
}

type useCase struct {
	beerRpeo BeerRepo
}

func NewUseCase(beerRepo BeerRepo) UseCase {
	return &useCase{
		beerRpeo: beerRepo,
	}
}

func (u *useCase) GetBeers() ([]Beer, error) {
	return u.beerRpeo.GetBeers()
}
