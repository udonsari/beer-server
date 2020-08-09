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
	// TODO Beer Detail 요청시 Related (향, 스타일, 랜덤 ...) Beer 리스트 추가 필요. 깔끔하게 5개 정도만 하자. (ReducedBeer)
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
	// TODO 여러 Beer 리스트에 대한 아래 정보를 모두 내린다면 무겁지 않은가 ? Comments는 Pagination ?
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Brewery   string   `json:"brewery"`
	ABV       float64  `json:"abv"`
	Country   string   `json:"country"`
	BeerStyle string   `json:"beer_style"`
	Aroma     []string `json:"aroma"`
	ImageURL  []string `json:"image_url"`

	Comments  []beer.Comment `json:"comments"`
	RateAvg   float64        `json:"rate_avg"`
	RateOwner *beer.Rate     `json:"rate_owner,omitempty"`
}

// type ReducedBeer struct {
// 	ID        int64    `json:"id"`
// 	Name      string   `json:"name"`
// 	Brewery   string   `json:"brewery"`
// 	ABV       float64  `json:"abv"`
// 	Country   string   `json:"country"`
// 	BeerStyle string   `json:"beer_style"`
// 	Aroma     []string `json:"aroma"`
// 	RateAvg   float64  `json:"rate_avg"`
// }
