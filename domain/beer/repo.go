package beer

type BeerRepo interface {
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	AddRate(rate Rate) error
	GetRates(beerID int64) ([]Rate, error)
	GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*Rate, error)
	AddComment(comment Comment) error
	GetComments(beerID int64) ([]Comment, error)
}
