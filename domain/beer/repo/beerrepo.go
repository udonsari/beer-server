package repo

import (
	"log"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
)

// TODO Attach Real DB and use ORM. Maybe ElasticSearch ?
// TODO Add Cache ?
type beerRepo struct {
	o orm.Ormer
}

func New(o orm.Ormer) *beerRepo {
	return &beerRepo{o: o}
}

func (r *beerRepo) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	log.Printf("BeerRepo - GetBeers() - args %+v", spew.Sdump(args))

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
