package repo

import "github.com/UdonSari/beer-server/domain/beer"

// TODO Attach Real DB and use ORM
// Maybe ElasticSearch ?

type beerRepo struct {
}

func New() *beerRepo {
	return &beerRepo{}
}

func (r *beerRepo) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	// TODO Implement query based on args
	return []beer.Beer{
		beer.Beer{
			Name:      "Wonder Pale Ale",
			Brewery:   "CraftBros",
			ABV:       5.7,
			Country:   "korea",
			BeerStyle: "ipa",
			Aroma:     "grape",
		},
		beer.Beer{
			Name:      "Super Pale Ale",
			Brewery:   "CraftBros",
			ABV:       6.3,
			Country:   "korea",
			BeerStyle: "ipa",
			Aroma:     "orange",
		},
	}, nil
}
