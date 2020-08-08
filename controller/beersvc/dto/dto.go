package dto

type GetBeersRequest struct {
	// ABV : Alcohol by Volume
	MinABV *float64 `query:"min_abv" validate:"required"`
	MaxABV *float64 `query:"max_abv" validate:"required"`

	Name    *string `query:"name" validate:"required"`
	Country *string `query:"country" validate:"required"`

	// TODO Support multi query for below args
	BeerStyle *string `query:"beer_style" validate:"required"`
	Aroma     *string `query:"aroma" validate:"required"`
}

type GetBeersResponse struct {
	Beers []Beer
}

type Beer struct {
	// TODO 여기 beer id도 내려가야함. 그래서 details 호출 가능하게
	Name      string  `json:"name"`
	Brewery   string  `json:"brewery"`
	ABV       float64 `json:"abv"`
	Country   string  `json:"country"`
	BeerStyle string  `json:"beer_style"`
	Aroma     string  `json:"aroma"`
}
