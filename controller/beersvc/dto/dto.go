package dto

type GetBeersRequest struct {
	// ABV : Alcohol by Volume
	MinABV *float64 `query:"min_abv" validate:"required"`
	MaxABV *float64 `query:"max_abv" validate:"required"`

	Country   *string `query:"country" validate:"required"`
	BeerStyle *string `query:"beer_style" validate:"required"`
	Aroma     *string `query:"aroma" validate:"required"`
}

type GetBeersResponse struct {
	Beers []Beer
}

type Beer struct {
	Name      string  `json:"name"`
	Brewery   string  `json:"brewary"`
	ABV       float64 `json:"abv"`
	Country   string  `json:"country"`
	BeerStyle string  `json:"beer_style"`
	Aroma     string  `json:"aroma"`
}
