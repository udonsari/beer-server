package repo

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/davecgh/go-spew/spew"
)

// TODO *** Attach Real DB and use ORM. Maybe ElasticSearch ?
// TODO *** Add Cache ?
type beerRepo struct {
}

func New() *beerRepo {
	return &beerRepo{}
}

func (r *beerRepo) GetBeer(beerID int64) (*beer.Beer, error) {
	return &beer.Beer{
		ID:        beerID,
		Name:      "Wonder Pale Ale",
		Brewery:   "CraftBros",
		ABV:       5.7,
		Country:   "korea",
		BeerStyle: "ipa",
		Aroma:     []string{"grape", "apple"},
	}, nil
}

func (r *beerRepo) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	log.Printf("BeerRepo - GetBeers() - args %+v", spew.Sdump(args))

	// TODO *** Implement query based on args
	return []beer.Beer{
		beer.Beer{
			ID:        int64(1),
			Name:      "Wonder Pale Ale",
			Brewery:   "CraftBros",
			ABV:       5.7,
			Country:   "korea",
			BeerStyle: "ipa",
			Aroma:     []string{"grape", "apple"},
		},
		beer.Beer{
			ID:        int64(2),
			Name:      "Super Pale Ale",
			Brewery:   "CraftBros",
			ABV:       6.3,
			Country:   "korea",
			BeerStyle: "ipa",
			Aroma:     []string{"orange", "apple"},
		},
	}, nil
}

func (r *beerRepo) AddRate(beerID int64, ratio float64, UserID int64) error {
	return nil
}

func (r *beerRepo) GetRates(beerID int64) ([]beer.Rate, error) {
	var rates []beer.Rate
	for i := 1; i <= 3; i++ {
		rates = append(rates, beer.Rate{
			BeerID: beerID,
			Ratio:  rand.Float64() * 5,
			UserID: int64(i),
		})
	}
	return rates, nil
}

func (r *beerRepo) AddComment(beerID int64, Content string, userID int64) error {
	return nil
}

func (r *beerRepo) GetComments(beerID int64) ([]beer.Comment, error) {
	var comments []beer.Comment
	for i := 1; i <= 3; i++ {
		comments = append(comments, beer.Comment{
			BeerID:  beerID,
			Content: "TEST_Comment_" + strconv.Itoa(i),
			UserID:  int64(i),
		})
	}
	return comments, nil
}
