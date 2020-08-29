package repo

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/davecgh/go-spew/spew"
)

type MockBeerRepo struct {
}

func NewMockBeerRepo() *MockBeerRepo {
	return &MockBeerRepo{}
}

func (r *MockBeerRepo) GetBeer(beerID int64) (*beer.Beer, error) {
	return &beer.Beer{
		ID:        beerID,
		Name:      "Wonder Pale Ale",
		Brewery:   "CraftBros",
		ABV:       5.7,
		Country:   "korea",
		BeerStyle: "ipa",
		Aroma:     []string{"grape", "apple"},
		ImageURL:  []string{"www.test_image_url.com"},
	}, nil
}

func (r *MockBeerRepo) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	log.Printf("BeerRepo - GetBeers() - args %+v", spew.Sdump(args))

	return []beer.Beer{
		beer.Beer{
			ID:        int64(1),
			Name:      "Wonder Pale Ale",
			Brewery:   "CraftBros",
			ABV:       5.7,
			Country:   "korea",
			BeerStyle: "ipa",
			Aroma:     []string{"grape", "apple"},
			ImageURL:  []string{"www.test_image_url.com"},
		},
		beer.Beer{
			ID:        int64(2),
			Name:      "Super Pale Ale",
			Brewery:   "CraftBros",
			ABV:       6.3,
			Country:   "korea",
			BeerStyle: "ipa",
			Aroma:     []string{"orange", "apple"},
			ImageURL:  []string{"www.test_image_url.com"},
		},
	}, nil
}

func (r *MockBeerRepo) AddRate(comment beer.Rate) error {
	return nil
}

func (r *MockBeerRepo) GetRates(beerID int64) ([]beer.Rate, error) {
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

func (r *MockBeerRepo) GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*beer.Rate, error) {
	return &beer.Rate{
		BeerID: beerID,
		Ratio:  rand.Float64() * 5,
		UserID: userID,
	}, nil
}

func (r *MockBeerRepo) AddComment(comment beer.Comment) error {
	return nil
}

func (r *MockBeerRepo) GetComments(beerID int64) ([]beer.Comment, error) {
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
