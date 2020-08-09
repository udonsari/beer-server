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
	for _, rate := range rates {
		rateAvg += rate.Ratio
	}
	rateAvg /= float64(len(rates))
	rateAvg = math.Floor(rateAvg*100) / 100

	return Beer{
		ID:        beer.ID,
		Name:      beer.Name,
		Brewery:   beer.Brewery,
		ABV:       beer.ABV,
		Country:   beer.Country,
		BeerStyle: beer.BeerStyle,
		Aroma:     beer.Aroma,
		Comments:  comments,
		RateAvg:   rateAvg,
		RateOwner: rateOwner,
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
