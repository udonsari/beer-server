package dto

import "github.com/UdonSari/beer-server/domain/beer"

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) MapBeerToDTOBeer(beer beer.Beer) Beer {
	return Beer{
		Name:      beer.Name,
		Brewery:   beer.Brewery,
		ABV:       beer.ABV,
		Country:   beer.Country,
		BeerStyle: beer.BeerStyle,
		Aroma:     beer.Aroma,
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

	args.Country = req.Country
	args.BeerStyle = req.BeerStyle
	args.Aroma = req.Aroma
	return &args, nil
}
