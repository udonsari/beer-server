package beer

type BeerQueryArgs struct {
	ABVInterval *ABVInterval
	Name        *string
	Country     *string
	BeerStyle   *string
	Aroma       *string
}

type ABVInterval struct {
	MinABV float64
	MaxABV float64
}
