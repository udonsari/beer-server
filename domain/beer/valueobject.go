package beer

type BeerQueryArgs struct {
	ABVInterval *ABVInterval
	Name        *string

	// Array Type Fields
	// TODO 이 쿼리 기준들은 OR로 계산
	Country   []string
	BeerStyle []string
	Aroma     []string
}

type ABVInterval struct {
	MinABV float64
	MaxABV float64
}
