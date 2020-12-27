package repo

import (
	"strings"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/util"
)

type mapper struct{}

func (m mapper) mapDBReviewToReview(dbReview DBReview) beer.Review {
	return beer.Review{
		ID:        dbReview.ID,
		BeerID:    dbReview.BeerID,
		Content:   dbReview.Content,
		Ratio:     dbReview.Ratio,
		UserID:    dbReview.UserID,
		CreatedAt: dbReview.CreatedAt,
	}
}

func (m mapper) mapReviewToDBReview(review beer.Review) DBReview {
	return DBReview{
		BeerID:  review.BeerID,
		Content: review.Content,
		Ratio:   review.Ratio,
		UserID:  review.UserID,
	}
}

func (m mapper) mapFavoriteToDBFavorite(favorite beer.Favorite) DBFavorite {
	return DBFavorite{
		BeerID: favorite.BeerID,
		Flag:   favorite.Flag,
		UserID: favorite.UserID,
	}
}

func (m mapper) mapDBFavoriteToFavorite(dbFavorite DBFavorite) beer.Favorite {
	return beer.Favorite{
		BeerID: dbFavorite.BeerID,
		Flag:   dbFavorite.Flag,
		UserID: dbFavorite.UserID,
	}
}
func (m mapper) mapBeerToDBBeer(beer beer.Beer) DBBeer {
	return DBBeer{
		Name:           beer.Name,
		Brewery:        beer.Brewery,
		ABV:            beer.ABV,
		Country:        beer.Country,
		BeerStyle:      beer.BeerStyle,
		AromaList:      m.splitAndGetString(beer.Aroma),
		ImageURLList:   m.splitAndGetString(beer.ImageURL),
		ThumbnailImage: beer.ThumbnailImage,
	}
}

func (m mapper) mapDBBeerToBeer(dbBeer DBBeer) beer.Beer {
	return beer.Beer{
		ID:             dbBeer.ID,
		Name:           dbBeer.Name,
		Brewery:        dbBeer.Brewery,
		ABV:            dbBeer.ABV,
		Country:        dbBeer.Country,
		BeerStyle:      dbBeer.BeerStyle,
		Aroma:          m.splitAndGetArray(dbBeer.AromaList, maxAromaListLen),
		ImageURL:       m.splitAndGetArray(dbBeer.ImageURLList, maxImageURLListLen),
		ThumbnailImage: dbBeer.ThumbnailImage,
		RateAvg:        dbBeer.RateAvg,
		ReviewCount:    dbBeer.ReviewCount,
	}
}

func (m mapper) splitAndGetArray(str string, maxLen int) []string {
	list := strings.Split(str, listSplitChar)
	return list[0:util.Min(len(list), maxLen)]
}

func (m mapper) splitAndGetString(strList []string) string {
	ret := ""
	for idx, v := range strList {
		ret += v
		if idx != len(strList)-1 {
			ret += listSplitChar
		}
	}
	return ret
}
