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
