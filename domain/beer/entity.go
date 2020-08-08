package beer

type Beer struct {
	Name      string
	Brewery   string
	ABV       float64
	Country   string
	BeerStyle string // TODO 스타일 세분화 필요. ex) 대분류 에일, 중분류 IPA, 소분류 NEIPA
	Aroma     string // TODO Should be list

	// TODO
	// image string
}
