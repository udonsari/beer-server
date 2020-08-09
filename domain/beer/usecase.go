package beer

type UseCase interface {
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	AddRate(beerID int64, ratio float64, UserID int64) error
	GetRates(beerID int64) ([]Rate, error)
	GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*Rate, error)
	AddComment(beerID int64, Content string, userID int64) error
	GetComments(beerID int64) ([]Comment, error)
}

type useCase struct {
	beerRepo BeerRepo
}

func NewUseCase(beerRepo BeerRepo) UseCase {
	return &useCase{
		beerRepo: beerRepo,
	}
}

func (u *useCase) GetBeers(args BeerQueryArgs) ([]Beer, error) {
	return u.beerRepo.GetBeers(args)
}

func (u *useCase) GetBeer(beerID int64) (*Beer, error) {
	return u.beerRepo.GetBeer(beerID)
}

func (u *useCase) AddRate(beerID int64, ratio float64, userID int64) error {
	return u.beerRepo.AddRate(beerID, ratio, userID)
}

func (u *useCase) GetRates(beerID int64) ([]Rate, error) {
	return u.beerRepo.GetRates(beerID)
}

func (u *useCase) GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*Rate, error) {
	return u.beerRepo.GetRatesByBeerIDAndUserID(beerID, userID)
}

func (u *useCase) AddComment(beerID int64, Content string, userID int64) error {
	return u.beerRepo.AddComment(beerID, Content, userID)
}

func (u *useCase) GetComments(beerID int64) ([]Comment, error) {
	return u.beerRepo.GetComments(beerID)
}

// TODO Implemnt. 이거 뭔가 이쁘게 만들어 보고 싶다
func (u *useCase) GetRelatedBeers(beerID int64) (*RelatedBeers, error) {
	return nil, nil
}
