package beer

type Beer struct {
	ID        int64
	Name      string
	Brewery   string
	ABV       float64
	Country   string
	BeerStyle string // TODO 스타일 세분화 필요. ex) 대분류 에일, 중분류 IPA, 소분류 NEIPA
	Aroma     []string

	// TODO
	// image string
}

// TODO Make dtoComment
type Comment struct {
	BeerID  int64  `json:"beer_id"`
	Content string `json:"content"`
	UserID  int64  `json:"user_id"`
}

type Rate struct {
	BeerID int64
	Ratio  float64
	UserID int64
}
