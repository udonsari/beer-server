package repo

const (
	maxAromaListLen    = 3
	maxImageURLListLen = 5

	listSplitChar = "___"
)

// TODO List 이런식으로 넣지 말고 정규화 해서 넣을까 고민
type DBBeer struct {
	ID        int64
	Name      string
	Brewery   string
	ABV       float64
	Country   string
	BeerStyle string

	AromaList    string
	ImageURLList string `gorm:"image_url_list"`

	RateAvg float64
}

func (DBBeer) TableName() string {
	return "beer_info"
}

type DBComment struct {
	ID      int64
	BeerID  int64
	Content string
	UserID  int64
}

func (DBComment) TableName() string {
	return "beer_comment"
}

type DBRate struct {
	ID     int64
	BeerID int64
	Ratio  float64
	UserID int64
}

func (DBRate) TableName() string {
	return "beer_rate"
}