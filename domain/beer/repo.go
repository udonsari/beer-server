package beer

type BeerRepo interface {
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	AddRate(beerID int64, ratio float64, UserID int64) error
	GetRates(beerID int64) ([]Rate, error)
	AddComment(beerID int64, Content string, userID int64) error
	GetComments(beerID int64) ([]Comment, error)
}
