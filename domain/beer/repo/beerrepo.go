package repo

import (
	"fmt"
	"log"
	"strings"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// TODO *** Attach ElasticSearch ? (Good for search)
// TODO *** Add Cache (Beer not change much). Currently only added duration (만약 BeerQueryArgs를 Key로 한다면 얼마나 메모리가 들어갈까 ?)
// Gorm 사용시 함수 인자로 구조체를 넣는것과, Key와 Value 넣는 것 구분해라. 동작이 이상한 경우 있음 (ex. Single Column 업데이트와 Batch Column 업데이트)
type beerRepo struct {
	db            *gorm.DB
	cacheDuration int64
	mapper
}

func New(db *gorm.DB, cacheDuration int64) *beerRepo {
	return &beerRepo{
		db:            db,
		cacheDuration: cacheDuration,
	}
}

func (r *beerRepo) Addbeer(beer beer.Beer) error {
	dbBeer := r.mapper.mapBeerToDBBeer(beer)
	res := r.db.Create(&dbBeer)
	return res.Error

}

func (r *beerRepo) GetBeer(beerID int64) (*beer.Beer, error) {
	query := DBBeer{ID: beerID}
	var dbBeer DBBeer
	if err := r.db.Where(&query).Find(&dbBeer).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get beer. beer id: %v", beerID))
		}
	}

	beer := r.mapper.mapDBBeerToBeer(dbBeer)
	return &beer, nil
}

func (r *beerRepo) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	log.Printf("beerRepo - GetBeers args : %+v", spew.Sdump(args))

	// TODO Optimize. Too Naiive Now
	var dbBeers []DBBeer
	if err := r.db.Find(&dbBeers).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	validFlag := make([]bool, len(dbBeers))
	for idx := range validFlag {
		validFlag[idx] = true
	}

	log.Printf("beerRepo - GetBeers base beer len : %+v", len(dbBeers))

	if args.ABVInterval != nil {
		for idx, dbBeer := range dbBeers {
			if dbBeer.ABV > args.ABVInterval.MaxABV || dbBeer.ABV < args.ABVInterval.MinABV {
				validFlag[idx] = false
			}
		}
	}

	if args.Name != nil {
		for idx, dbBeer := range dbBeers {
			if !strings.Contains(dbBeer.Name, *args.Name) {
				validFlag[idx] = false
			}
		}
	}

	if args.Country != nil {
		for idx, dbBeer := range dbBeers {
			countryContains := false
			for _, country := range args.Country {
				if strings.Contains(dbBeer.Country, country) {
					countryContains = true
				}
			}
			if !countryContains {
				validFlag[idx] = false
			}
		}
	}

	if args.BeerStyle != nil {
		for idx, dbBeer := range dbBeers {
			styleContains := false
			for _, beerStyle := range args.BeerStyle {
				if strings.Contains(dbBeer.BeerStyle, beerStyle) {
					styleContains = true
				}
			}
			if !styleContains {
				validFlag[idx] = false
			}
		}
	}

	if args.Aroma != nil {
		// Beer의 Aroma도 n개, 쿼리의 Aroma도 n개. 우선 String Concate로 해놓으니까 그냥 String Contains 가능
		for idx, dbBeer := range dbBeers {
			aromaContains := false
			for _, aroma := range args.Aroma {
				if strings.Contains(dbBeer.AromaList, aroma) {
					aromaContains = true
				}
			}
			if !aromaContains {
				validFlag[idx] = false
			}
		}
	}

	var beers []beer.Beer
	for idx := range dbBeers {
		if validFlag[idx] {
			beers = append(beers, r.mapDBBeerToBeer(dbBeers[idx]))
		}
	}
	log.Printf("beerRepo - GetBeers return beer len : %+v", len(beers))
	return beers, nil
}

func (r *beerRepo) UpdateBeerRateAvg(beerID int64, rateAvg float64) error {
	res := r.db.Model(&DBBeer{}).Where("id = ?", beerID).Update("rate_avg", rateAvg)
	return res.Error
}

func (r *beerRepo) AddReview(review beer.Review) error {
	dbReview := r.mapper.mapReviewToDBReview(review)
	res := r.db.Create(&dbReview)
	// https://github.com/go-gorm/gorm/issues/2903
	// Gorm V1에서는 Duplicate Error를 정의하지 않음
	if res.Error != nil && strings.Contains(res.Error.Error(), "Error 1062: Duplicate entry") {
		return errors.New("already added review")
	}
	return res.Error
}

func (r *beerRepo) GetReviews(beerID int64) ([]beer.Review, error) {
	query := DBReview{BeerID: beerID}
	var dbReviews []DBReview
	if err := r.db.Where(&query).Find(&dbReviews).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get review. beer id: %v", beerID))
		}
	}

	var reviews []beer.Review
	for _, dbReview := range dbReviews {
		reviews = append(reviews, r.mapper.mapDBReviewToReview(dbReview))
	}
	return reviews, nil
}

func (r *beerRepo) GetReviewCount(beerID int64) (int64, error) {
	var count int64
	res := r.db.Model(&DBReview{}).Where("beer_id = ?", beerID).Count(&count)
	return count, res.Error
}

func (r *beerRepo) GetReviewByBeerIDAndUserID(beerID int64, userID int64) (*beer.Review, error) {
	query := DBReview{BeerID: beerID, UserID: userID}
	var dbReview DBReview
	if err := r.db.Where(&query).First(&dbReview).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get review. beer id: %v, user id : %v", beerID, userID))
		}
	}
	review := r.mapper.mapDBReviewToReview(dbReview)
	return &review, nil
}
