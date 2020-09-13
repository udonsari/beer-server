package beer

type BeerRepo interface {
	Addbeer(beer Beer) error
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	UpdateBeerRateAvg(beerID int64, rateAvg float64) error
	AddRate(rate Rate) error
	GetRates(beerID int64) ([]Rate, error)
	GetRatesCount(beerID int64) (int64, error)
	GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*Rate, error)
	AddComment(comment Comment) error
	GetComments(beerID int64) ([]Comment, error)
}
