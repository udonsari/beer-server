package dto

import "time"

type GetBeersRequest struct {
	// ABV : Alcohol by Volume
	MinABV *float64 `query:"min_abv"`
	MaxABV *float64 `query:"max_abv"`

	Name *string `query:"name"`

	// Array Type Fields
	Country   []string `query:"country"`
	BeerStyle []string `query:"beer_style"`
	Aroma     []string `query:"aroma"`

	// Cursor Pagination
	Cursor   *int64 `query:"cursor"`
	MaxCount *int64 `query:"max_count"`

	SortBy *string `query:"sort_by"`
}

type GetBeersResponse struct {
	ReducedBeer []ReducedBeer `json:"beers"`

	// Cursor Pagination
	Cursor *int64 `json:"next_cursor"` // If it is last page, this will be nil
}

type GetBeerRequest struct {
	BeerID int64 `query:"beer_id"`
}

type GetBeerResponse struct {
	Beer         Beer          `json:"beer"`
	RelatedBeers *RelatedBeers `json:"related_beers,omitempty"`
}

type AddReviewRequest struct {
	BeerID  int64   `form:"beer_id"`
	Content string  `form:"content"`
	Ratio   float64 `form:"ratio"`
}

type GetReviewResponse struct {
	Reviews []Review `json:"reviews"`
}

type AddFavoriteRequest struct {
	BeerID int64 `form:"beer_id"`
	Flag   bool  `form:"flag"`
}

type Beer struct {
	// TODO Beer 리스트에 대한 아래 정보를 모두 내린다면 무겁지 않은가 ? Reviews는 Pagination ?
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Brewery        string   `json:"brewery"`
	ABV            float64  `json:"abv"`
	Country        string   `json:"country"`
	BeerStyle      string   `json:"beer_style"`
	Aroma          []string `json:"aroma"`
	ImageURL       []string `json:"image_url"`
	ThumbnailImage string   `json:"thumbnail_image"`

	Reviews      []Review `json:"reviews"`
	RateAvg      float64  `json:"rate_avg"`
	ReviewCount  int64    `json:"review_count"`
	ReviewOwner  *Review  `json:"review_owner,omitempty"`
	FavoriteFlag bool     `json:"favorite_flag"`
}

type RelatedBeers struct {
	AromaRelatedBeer    []ReducedBeer `json:"aroma_related"`
	StyleRelatedBeer    []ReducedBeer `json:"style_related"`
	RandomlyRelatedBeer []ReducedBeer `json:"randomly_related"`
}

type ReducedBeer struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Brewery        string   `json:"brewery"`
	ABV            float64  `json:"abv"`
	Country        string   `json:"country"`
	BeerStyle      string   `json:"beer_style"`
	Aroma          []string `json:"aroma"`
	ThumbnailImage string   `json:"thumbnail_image"`
	RateAvg        float64  `json:"rate_avg"`
	ReviewCount    int64    `json:"review_count"`
	FavoriteFlag   bool     `json:"favorite_flag"`
}

type Review struct {
	ReducedBeer ReducedBeer `json:"beer"`
	Content     string      `json:"content"`
	Ratio       float64     `json:"ratio"`
	UserID      int64       `json:"user_id"`
	NickName    string      `json:"nickname"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Favorite struct {
	ReducedBeer ReducedBeer `json:"beer"`
	UserID      int64       `json:"user_id"`
	BeerID      int64       `json:"beer_id"`
}

type AppConfig struct {
	AromaList     []string `json:"aroma_list"`
	CountryList   []string `json:"country_list"`
	BeerStyleList []string `json:"style_list"`
	MinABV        float64  `json:"min_abv"`
	MaxABV        float64  `json:"max_abv"`
}
