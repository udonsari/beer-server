package beer

import (
	"math/rand"
	"time"

	"github.com/UdonSari/beer-server/util"
)

type UseCase interface {
	Addbeer(beer Beer) error
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	AddRate(rate Rate) error
	GetRates(beerID int64) ([]Rate, error)
	GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*Rate, error)
	AddComment(comment Comment) error
	GetComments(beerID int64) ([]Comment, error)
	GetRelatedBeers(beerID int64) (*RelatedBeers, error)
}

type useCase struct {
	beerRepo BeerRepo
}

func NewUseCase(beerRepo BeerRepo) UseCase {
	rand.Seed(time.Now().UnixNano())
	return &useCase{
		beerRepo: beerRepo,
	}
}

func (u *useCase) Addbeer(beer Beer) error {
	return u.beerRepo.Addbeer(beer)
}

func (u *useCase) GetBeers(args BeerQueryArgs) ([]Beer, error) {
	return u.beerRepo.GetBeers(args)
}

func (u *useCase) GetBeer(beerID int64) (*Beer, error) {
	return u.beerRepo.GetBeer(beerID)
}

func (u *useCase) AddRate(rate Rate) error {
	beer, err := u.GetBeer(rate.BeerID)
	if err != nil {
		return err
	}
	ratesLen, err := u.beerRepo.GetRatesCount(rate.BeerID)
	if err != nil {
		return err
	}

	var newRateAvg float64
	if ratesLen == 0 {
		newRateAvg = rate.Ratio
	} else {
		newRateAvg = (beer.RateAvg*float64(ratesLen) + rate.Ratio) / (float64(ratesLen) + 1.0)
	}
	err = u.beerRepo.UpdateBeerRateAvg(rate.BeerID, newRateAvg)
	if err != nil {
		return err
	}

	return u.beerRepo.AddRate(rate)
}

func (u *useCase) GetRates(beerID int64) ([]Rate, error) {
	return u.beerRepo.GetRates(beerID)
}

func (u *useCase) GetRatesByBeerIDAndUserID(beerID int64, userID int64) (*Rate, error) {
	return u.beerRepo.GetRatesByBeerIDAndUserID(beerID, userID)
}

func (u *useCase) AddComment(comment Comment) error {
	return u.beerRepo.AddComment(comment)
}

func (u *useCase) GetComments(beerID int64) ([]Comment, error) {
	return u.beerRepo.GetComments(beerID)
}

func (u *useCase) GetRelatedBeers(beerID int64) (*RelatedBeers, error) {
	// TODO Improve
	baseBeer, err := u.beerRepo.GetBeer(beerID)
	if err != nil {
		return nil, err
	}

	var relatedBeers RelatedBeers

	var aromaQueryArgs BeerQueryArgs
	aromaQueryArgs.Aroma = baseBeer.Aroma
	aromaRelatedBeers, err := u.getRelatedBeersWithQueryArgs(aromaQueryArgs)
	if err != nil {
		return nil, err
	}
	relatedBeers.AromaRelatedBeer = aromaRelatedBeers

	var styleQueryArgs BeerQueryArgs
	styleQueryArgs.BeerStyle = append(styleQueryArgs.BeerStyle, baseBeer.BeerStyle)
	styleRelatedBeers, err := u.getRelatedBeersWithQueryArgs(styleQueryArgs)
	if err != nil {
		return nil, err
	}
	relatedBeers.StyleRelatedBeer = styleRelatedBeers

	var randomlyQueryArgs BeerQueryArgs
	randomlyRelatedBeers, err := u.getRelatedBeersWithQueryArgs(randomlyQueryArgs)
	if err != nil {
		return nil, err
	}
	relatedBeers.RandomlyRelatedBeer = randomlyRelatedBeers

	return &relatedBeers, nil
}

func (u *useCase) getRelatedBeersWithQueryArgs(args BeerQueryArgs) ([]Beer, error) {
	relatedBeers, err := u.beerRepo.GetBeers(args)
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(relatedBeers), func(i, j int) {
		relatedBeers[i], relatedBeers[j] = relatedBeers[j], relatedBeers[i]
	})
	return relatedBeers[0:util.Min(len(relatedBeers), relatedBeersMaxLen)], nil
}
