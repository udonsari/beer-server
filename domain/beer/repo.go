package beer

type BeerRepo interface {
	Addbeer(beer Beer) error
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	UpdateBeerRateAvg(beerID int64, rateAvg float64) error
	AddReview(review Review) error
	GetReviews(beerID int64) ([]Review, error)
	GetReviewCount(beerID int64) (int64, error)
	GetReviewByBeerIDAndUserID(beerID int64, userID int64) (*Review, error)
}
