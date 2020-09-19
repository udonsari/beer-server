package dto

import (
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/util"
)

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) MapBeerToDTOBeer(beer beer.Beer, reviews []beer.Review, reviewOwner *beer.Review) Beer {
	var dtoReviews []Review
	for _, review := range reviews {
		dtoReview := m.MapReviewToDTOReview(&review)
		if dtoReview != nil {
			dtoReviews = append(dtoReviews, *dtoReview)
		}
	}

	return Beer{
		ID:             beer.ID,
		Name:           beer.Name,
		Brewery:        beer.Brewery,
		ABV:            util.Floor(beer.ABV, 2),
		Country:        beer.Country,
		BeerStyle:      beer.BeerStyle,
		Aroma:          beer.Aroma,
		ImageURL:       beer.ImageURL,
		ThumbnailImage: beer.ThumbnailImage,
		Reviews:        dtoReviews,
		RateAvg:        util.Floor(beer.RateAvg, 2),
		ReviewOwner:    m.MapReviewToDTOReview(reviewOwner),
	}
}

func (m *Mapper) MapBeerToDTReducedBeer(beer beer.Beer) ReducedBeer {
	return ReducedBeer{
		ID:             beer.ID,
		Name:           beer.Name,
		Brewery:        beer.Brewery,
		ABV:            util.Floor(beer.ABV, 2),
		Country:        beer.Country,
		BeerStyle:      beer.BeerStyle,
		Aroma:          beer.Aroma,
		ThumbnailImage: beer.ThumbnailImage,
		RateAvg:        util.Floor(beer.RateAvg, 2),
	}
}

func (m *Mapper) MapRelatedBeersToDTORelatedBeers(relatedBeer *beer.RelatedBeers) *RelatedBeers {
	if relatedBeer == nil {
		return nil
	}
	var dtoRelatedBeers RelatedBeers
	for _, b := range relatedBeer.AromaRelatedBeer {
		dtoRelatedBeers.AromaRelatedBeer = append(dtoRelatedBeers.AromaRelatedBeer, m.MapBeerToDTReducedBeer(b))
	}
	for _, b := range relatedBeer.StyleRelatedBeer {
		dtoRelatedBeers.StyleRelatedBeer = append(dtoRelatedBeers.StyleRelatedBeer, m.MapBeerToDTReducedBeer(b))
	}
	for _, b := range relatedBeer.RandomlyRelatedBeer {
		dtoRelatedBeers.RandomlyRelatedBeer = append(dtoRelatedBeers.RandomlyRelatedBeer, m.MapBeerToDTReducedBeer(b))
	}
	return &dtoRelatedBeers
}

func (m *Mapper) MapReviewToDTOReview(review *beer.Review) *Review {
	if review == nil {
		return nil
	}
	return &Review{
		BeerID:  review.BeerID,
		Content: review.Content,
		Ratio:   util.Floor(review.Ratio, 2),
		UserID:  review.UserID,
	}
}

func (m *Mapper) MapGetBeersRequestToBeerQueryArgs(req GetBeersRequest) (*beer.BeerQueryArgs, error) {
	if (req.MinABV != nil && req.MaxABV == nil) || (req.MinABV == nil && req.MaxABV != nil) {
		return nil, NewMapperError("MinABV and MaxABV should come together")
	}

	var args beer.BeerQueryArgs
	if req.MinABV != nil {
		args.ABVInterval = &beer.ABVInterval{
			MinABV: *req.MinABV,
			MaxABV: *req.MaxABV,
		}
	}

	args.Name = req.Name
	args.Country = req.Country
	args.BeerStyle = req.BeerStyle
	args.Aroma = req.Aroma

	args.Cursor = req.Cursor
	args.MaxCount = req.MaxCount
	return &args, nil
}
