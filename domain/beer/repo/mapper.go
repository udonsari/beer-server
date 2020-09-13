package repo

import (
	"strings"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/util"
)

type mapper struct{}

func (m mapper) mapDBCommentToComment(dbComment DBComment) beer.Comment {
	return beer.Comment{
		ID:      dbComment.ID,
		BeerID:  dbComment.BeerID,
		Content: dbComment.Content,
		UserID:  dbComment.UserID,
	}
}

func (m mapper) mapCommentToDBComment(comment beer.Comment) DBComment {
	return DBComment{
		BeerID:  comment.BeerID,
		Content: comment.Content,
		UserID:  comment.UserID,
	}
}

func (m mapper) mapDBRateToRate(dbRate DBRate) beer.Rate {
	return beer.Rate{
		ID:     dbRate.ID,
		BeerID: dbRate.BeerID,
		Ratio:  dbRate.Ratio,
		UserID: dbRate.UserID,
	}
}

func (m mapper) mapRateToDBRate(rate beer.Rate) DBRate {
	return DBRate{
		BeerID: rate.BeerID,
		Ratio:  rate.Ratio,
		UserID: rate.UserID,
	}
}

func (m mapper) mapBeerToDBBeer(beer beer.Beer) DBBeer {
	return DBBeer{
		Name:         beer.Name,
		Brewery:      beer.Brewery,
		ABV:          beer.ABV,
		Country:      beer.Country,
		BeerStyle:    beer.BeerStyle,
		AromaList:    m.splitAndGetString(beer.Aroma),
		ImageURLList: m.splitAndGetString(beer.ImageURL),
		RateAvg:      beer.RateAvg,
	}
}

func (m mapper) mapDBBeerToBeer(dbBeer DBBeer) beer.Beer {
	return beer.Beer{
		ID:        dbBeer.ID,
		Name:      dbBeer.Name,
		Brewery:   dbBeer.Brewery,
		ABV:       dbBeer.ABV,
		Country:   dbBeer.Country,
		BeerStyle: dbBeer.BeerStyle,
		Aroma:     m.splitAndGetArray(dbBeer.AromaList, maxAromaListLen),
		ImageURL:  m.splitAndGetArray(dbBeer.ImageURLList, maxImageURLListLen),
		RateAvg:   dbBeer.RateAvg,
	}
}

func (m mapper) splitAndGetArray(str string, maxLen int) []string {
	list := strings.Split(str, listSplitChar)
	return list[0:util.Min(len(list), maxLen)]
}

func (m mapper) splitAndGetString(strList []string) string {
	ret := ""
	for _, v := range strList {
		ret += v + listSplitChar
	}
	return ret
}
