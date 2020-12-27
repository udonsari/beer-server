package beer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/UdonSari/beer-server/util"
)

type UseCase interface {
	AddBeer(beer Beer) error
	GetBeers(args BeerQueryArgs) ([]Beer, error)
	GetBeer(beerID int64) (*Beer, error)
	AddReview(review Review) error
	GetReviews(beerID int64) ([]Review, error)
	GetReviewsByUserID(userID int64) ([]Review, error)
	GetReviewByBeerIDAndUserID(beerID int64, userID int64) (*Review, error)
	GetRelatedBeers(beerID int64) (*RelatedBeers, error)
	AddFavorite(avorite Favorite) error
	GetFavorites(userID int64) ([]Favorite, error)
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

func (u *useCase) AddBeer(beer Beer) error {
	return u.beerRepo.AddBeer(beer)
}

func (u *useCase) GetBeers(args BeerQueryArgs) ([]Beer, error) {
	return u.beerRepo.GetBeers(args)
}

func (u *useCase) GetBeer(beerID int64) (*Beer, error) {
	return u.beerRepo.GetBeer(beerID)
}

func (u *useCase) AddReview(review Review) error {
	beer, err := u.GetBeer(review.BeerID)
	if err != nil {
		return err
	} else if beer == nil {
		return fmt.Errorf("no matching beer to add review")
	}

	var newRateAvg float64
	if beer.ReviewCount == 0 {
		newRateAvg = review.Ratio
	} else {
		preReview, err := u.beerRepo.GetReviewByBeerIDAndUserID(review.BeerID, review.UserID)
		if err != nil {
			return err
		}
		if preReview == nil {
			newRateAvg = (beer.RateAvg*float64(beer.ReviewCount) + review.Ratio) / (float64(beer.ReviewCount) + 1.0)
		} else {
			newRateAvg = (beer.RateAvg*float64(beer.ReviewCount) + review.Ratio - preReview.Ratio) / (float64(beer.ReviewCount))
		}
	}

	err = u.beerRepo.UpdateBeerRateAvg(review.BeerID, newRateAvg)
	if err != nil {
		return err
	}

	return u.beerRepo.AddReview(review)
}

func (u *useCase) GetReviews(beerID int64) ([]Review, error) {
	return u.beerRepo.GetReviews(beerID)
}

func (u *useCase) GetReviewsByUserID(userID int64) ([]Review, error) {
	return u.beerRepo.GetReviewsByUserID(userID)
}

func (u *useCase) GetReviewByBeerIDAndUserID(beerID int64, userID int64) (*Review, error) {
	return u.beerRepo.GetReviewByBeerIDAndUserID(beerID, userID)
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

func (u *useCase) AddFavorite(favorite Favorite) error {
	return u.beerRepo.AddFavorite(favorite)
}

func (u *useCase) GetFavorites(userID int64) ([]Favorite, error) {
	return u.beerRepo.GetFavorites(userID)
}
