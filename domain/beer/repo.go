package beer

type BeerRepo interface {
	GetBeers() ([]Beer, error)
}
