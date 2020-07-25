package beer

type BeerRepo interface {
	GetBeers(args BeerQueryArgs) ([]Beer, error)
}
