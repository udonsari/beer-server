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

func (m *Mapper) MapBeerToDTOBeer(beer beer.Beer, comments []beer.Comment, rateOwner *beer.Rate) Beer {
	var dtoComments []Comment
	for _, comment := range comments {
		dtoComments = append(dtoComments, m.mapCommentToDTOComment(comment))
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
		Comments:       dtoComments,
		RateAvg:        util.Floor(beer.RateAvg, 2),
		RateOwner:      m.mapRateToDTORate(rateOwner),
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

func (m *Mapper) mapCommentToDTOComment(comment beer.Comment) Comment {
	return Comment{
		BeerID:  comment.BeerID,
		Content: comment.Content,
		UserID:  comment.UserID,
	}
}

func (m *Mapper) mapRateToDTORate(rate *beer.Rate) *Rate {
	if rate == nil {
		return nil
	}
	return &Rate{
		BeerID: rate.BeerID,
		Ratio:  util.Floor(rate.Ratio, 2),
		UserID: rate.UserID,
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
	return &args, nil
}
