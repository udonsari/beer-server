package repo_test

import (
	"testing"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/util"
	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

const (
	listSplitChar = "___"
)

func TestRepoTestSuite(t *testing.T) {
	suite.Run(t, new(repoTestSuite))
}

type repoTestSuite struct {
	*util.GormHelper
	suite.Suite
	mock sqlmock.Sqlmock
	db   *gorm.DB
	repo beer.BeerRepo
}

func (ts *repoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	ts.NoError(err)
	ts.mock = mock

	ts.db, err = gorm.Open("mysql", db)
	ts.NoError(err)
	ts.db.LogMode(true)
	ts.db = ts.db.Set("gorm:update_column", true)

	ts.repo = repo.New(ts.db, 1)
}

func (ts *repoTestSuite) Test_AddBeer() {
	// TODO 왜 Begin을 걸어야하지 확인

	var input beer.Beer
	faker.FakeData(&input)
	input.ImageURL = []string{"http://naver.com", "http://google.com"}
	input.Aroma = []string{"A", "B", "C"}
	parsedImageURLList := "http://naver.com___http://google.com"
	parsedAromaList := "A___B___C"

	expectedSQL := "INSERT INTO `beer_info` (`name`,`brewery`,`abv`,`country`,`beer_style`,`aroma_list`,`image_url_list`,`thumbnail_image`,`rate_avg`,`review_count`) VALUES (?,?,?,?,?,?,?,?,?,?)"

	ts.mock.ExpectBegin()
	ts.mock.ExpectExec(ts.FixedFullRe(expectedSQL)).WithArgs(input.Name, input.Brewery, input.ABV, input.Country, input.BeerStyle, parsedAromaList, parsedImageURLList, input.ThumbnailImage, float64(0), 0).WillReturnResult(sqlmock.NewResult(1, 1))
	ts.mock.ExpectCommit()

	err := ts.repo.AddBeer(input)
	ts.NoError(err)
}

func (ts *repoTestSuite) Test_GetBeer() {
	var beer beer.Beer
	faker.FakeData(&beer)
	beer.ImageURL = []string{"http://naver.com", "http://google.com"}
	beer.Aroma = []string{"A", "B", "C"}

	expectedSQL := "SELECT * FROM `beer_info` WHERE (`beer_info`.`id` = ?)"
	ts.mock.ExpectQuery(ts.FixedFullRe(expectedSQL)).WithArgs(beer.ID).WillReturnRows((&util.MockRowBuilder{}).Add(ts.mapBeerToDBBeer(beer)).Build())

	res, err := ts.repo.GetBeer(beer.ID)
	ts.NoError(err)
	ts.NotNil(res)
	ts.True(ts.equalBeer(beer, *res))
}

func (ts *repoTestSuite) mapBeerToDBBeer(beer beer.Beer) repo.DBBeer {
	return repo.DBBeer{
		Name:           beer.Name,
		Brewery:        beer.Brewery,
		ABV:            beer.ABV,
		Country:        beer.Country,
		BeerStyle:      beer.BeerStyle,
		AromaList:      ts.splitAndGetString(beer.Aroma),
		ImageURLList:   ts.splitAndGetString(beer.ImageURL),
		ThumbnailImage: beer.ThumbnailImage,
		RateAvg:        beer.RateAvg,
		ReviewCount:    beer.ReviewCount,
	}
}

func (ts *repoTestSuite) splitAndGetString(strList []string) string {
	ret := ""
	for idx, v := range strList {
		ret += v
		if idx != len(strList)-1 {
			ret += listSplitChar
		}
	}
	return ret
}

func (Ts *repoTestSuite) equalBeer(b1 beer.Beer, b2 beer.Beer) bool {
	ret := b1.Name == b2.Name &&
		b1.Brewery == b2.Brewery &&
		b1.ABV == b2.ABV &&
		b1.Country == b2.Country &&
		b1.BeerStyle == b2.BeerStyle &&
		b1.ThumbnailImage == b2.ThumbnailImage &&
		b1.RateAvg == b2.RateAvg &&
		b1.ReviewCount == b2.ReviewCount &&
		len(b1.Aroma) == len(b2.Aroma) &&
		len(b1.ImageURL) == len(b2.ImageURL)

	for idx := range b1.Aroma {
		ret = ret && (b1.Aroma[idx] == b2.Aroma[idx])
	}

	for idx := range b1.ImageURL {
		ret = ret && (b1.ImageURL[idx] == b2.ImageURL[idx])
	}

	return ret
}
