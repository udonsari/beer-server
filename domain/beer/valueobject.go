package beer

const (
	relatedBeersMaxLen = 3
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
