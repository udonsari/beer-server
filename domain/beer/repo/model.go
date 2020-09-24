package repo

import "time"

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

	AromaList      string
	ImageURLList   string `gorm:"image_url_list"`
	ThumbnailImage string

	RateAvg     float64
	ReviewCount int64
}

func (DBBeer) TableName() string {
	return "beer_info"
}

type DBReview struct {
	ID        int64
	BeerID    int64
	Ratio     float64
	Content   string
	UserID    int64
	CreatedAt time.Time
}

func (DBReview) TableName() string {
	return "beer_review"
}
