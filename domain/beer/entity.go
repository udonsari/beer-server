package beer

type Beer struct {
	ID        int64
	Name      string
	Brewery   string
	ABV       float64
	Country   string
	BeerStyle string // TODO *** 스타일 세분화 필요. ex) 대분류 에일, 중분류 IPA, 소분류 NEIPA
	Aroma     []string
	ImageURL  []string
}

type Comment struct {
	ID      int64
	BeerID  int64
	Content string
	UserID  int64
}

type Rate struct {
	ID     int64
	BeerID int64
	Ratio  float64
	UserID int64
}
