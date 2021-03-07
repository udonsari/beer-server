package beer

import "time"

type Beer struct {
	ID             int64
	Name           string
	Brewery        string
	ABV            float64
	Country        string
	BeerStyle      string
	Aroma          []string
	ImageURL       []string
	ThumbnailImage string
	RateAvg        float64
	ReviewCount    int64
}

type Review struct {
	ID        int64
	BeerID    int64
	Content   string
	Ratio     float64
	UserID    int64
	CreatedAt time.Time
}

type Favorite struct {
	ID     int64
	BeerID int64
	Flag   bool
	UserID int64
}

type UserBeerConfig struct {
	ID     int64
	UserID int64
	Aroma  []string
	Style  []string
}
