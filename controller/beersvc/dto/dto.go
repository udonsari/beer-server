package dto

import "github.com/UdonSari/beer-server/domain/beer"

type GetBeersRequest struct {
	// ABV : Alcohol by Volume
	MinABV *float64 `query:"min_abv"`
	MaxABV *float64 `query:"max_abv"`

	Name *string `query:"name"`

	// Array Type Fields
	Country   []string `query:"country"`
	BeerStyle []string `query:"beer_style"`
	Aroma     []string `query:"aroma"`
}

type GetBeersResponse struct {
	Beers []Beer
}

type GetBeerRequest struct {
	BeerID int64 `query:"beer_id"`
}

type GetBeerResponse struct {
	Beer Beer `json:"beer"`
}

type AddRateRequest struct {
	BeerID int64   `form:"beer_id"`
	Ratio  float64 `form:"ratio"`
}

type AddCommentRequest struct {
	BeerID  int64  `form:"beer_id"`
	Content string `form:"content"`
}

type Beer struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Brewery   string         `json:"brewery"`
	ABV       float64        `json:"abv"`
	Country   string         `json:"country"`
	BeerStyle string         `json:"beer_style"`
	Aroma     []string       `json:"aroma"`
	Comments  []beer.Comment `json:"comments"`
	RateAvg   float64        `json:"rate_avg"`

	// TODO 지금 접속한 사람이 이 맥주에 대해 매긴 Rate도 내려주기
}
