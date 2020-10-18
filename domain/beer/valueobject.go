package beer

const (
	relatedBeersMaxLen = 3
)

const (
	SortByRateAvgAsc      = "rate_avg_asc"
	SortByRateAvgDesc     = "rate_avg_desc"
	SortByReviewCountAsc  = "review_count_asc"
	SortByReviewCountDesc = "review_count_desc"
)

type BeerQueryArgs struct {
	ABVInterval *ABVInterval
	Name        *string

	// Array Type Fields
	Country   []string
	BeerStyle []string
	Aroma     []string

	// Cursor Pagination
	Cursor   *int64
	MaxCount *int64

	SortBy *string
}

type ABVInterval struct {
	MinABV float64
	MaxABV float64
}

type RelatedBeers struct {
	AromaRelatedBeer    []Beer
	StyleRelatedBeer    []Beer
	RandomlyRelatedBeer []Beer
}
