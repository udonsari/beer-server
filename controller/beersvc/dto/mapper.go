package dto

import (
	"math"

	"github.com/UdonSari/beer-server/domain/beer"
)

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) MapBeerToDTOBeer(beer beer.Beer, comments []beer.Comment, rates []beer.Rate, rateOwner *beer.Rate) Beer {
	rateAvg := float64(0)
	if len(rates) > 0 {
		for _, rate := range rates {
			rateAvg += rate.Ratio
		}
		rateAvg /= float64(len(rates))
		rateAvg = math.Floor(rateAvg*100) / 100
	}

	var dtoComments []Comment
	for _, comment := range comments {
		dtoComments = append(dtoComments, m.mapCommentToDTOComment(comment))
	}

	beer.ABV = math.Floor(beer.ABV*100) / 100

	return Beer{
		ID:        beer.ID,
		Name:      beer.Name,
		Brewery:   beer.Brewery,
		ABV:       beer.ABV,
		Country:   beer.Country,
		BeerStyle: beer.BeerStyle,
		Aroma:     beer.Aroma,
		ImageURL:  beer.ImageURL,
		Comments:  dtoComments,
		RateAvg:   rateAvg,
		RateOwner: rateOwner,
	}
}

func (m *Mapper) MapBeerToDTReducedBeer(beer beer.Beer) ReducedBeer {
	beer.ABV = math.Floor(beer.ABV*100) / 100

	// TODO 여기서 AverageRatio를 구하기 어려운 면이 있음

	return ReducedBeer{
		ID:        beer.ID,
		Name:      beer.Name,
		Brewery:   beer.Brewery,
		ABV:       beer.ABV,
		Country:   beer.Country,
		BeerStyle: beer.BeerStyle,
		Aroma:     beer.Aroma,
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
		dtoRelatedBeers.AromaRelatedBeer = append(dtoRelatedBeers.AromaRelatedBeer, m.MapBeerToDTReducedBeer(b))
	}
	for _, b := range relatedBeer.RandomlyRelatedBeer {
		dtoRelatedBeers.AromaRelatedBeer = append(dtoRelatedBeers.AromaRelatedBeer, m.MapBeerToDTReducedBeer(b))
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
