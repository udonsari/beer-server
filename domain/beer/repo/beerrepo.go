package repo

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/util"
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

func (r *beerRepo) AddBeer(beer beer.Beer) error {
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

	// TODO Need test really
	var dbBeers []DBBeer
	baseQuery := r.db

	if args.ABVInterval != nil {
		baseQuery = baseQuery.Where("abv BETWEEN ? AND ? ", args.ABVInterval.MinABV, args.ABVInterval.MaxABV)
	}

	if args.Name != nil {
		baseQuery = baseQuery.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *args.Name))
	}

	if args.Country != nil {
		baseQuery = baseQuery.Where("country IN (?)", args.Country)
	}

	if args.BeerStyle != nil {
		baseQuery = baseQuery.Where("beer_style IN (?)", args.BeerStyle)
	}

	if args.Aroma != nil {
		for _, aroma := range args.Aroma {
			baseQuery = baseQuery.Where("aroma_list LIKE ?", fmt.Sprintf("%%%s%%", aroma))
		}
	}

	// Cursor is just id value
	cursor := args.Cursor
	maxCount := args.MaxCount
	// baseQuery = baseQuery.Limit(maxCount)

	// Sort by
	if args.SortBy == nil {
		baseQuery = baseQuery.Order("id asc")
	} else if *args.SortBy == beer.SortByRateAvgAsc {
		baseQuery = baseQuery.Order("rate_avg asc")
	} else if *args.SortBy == beer.SortByRateAvgDesc {
		baseQuery = baseQuery.Order("rate_avg desc")
	} else if *args.SortBy == beer.SortByReviewCountAsc {
		baseQuery = baseQuery.Order("review_count asc")
	} else if *args.SortBy == beer.SortByReviewCountDesc {
		baseQuery = baseQuery.Order("review_count desc")
	}

	if err := baseQuery.Find(&dbBeers).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	var limitedDBBeers []DBBeer
	var i int
	if cursor <= 0 {
		i = -1
	} else {
		for ; i < len(dbBeers); i++ {
			if dbBeers[i].ID == cursor {
				break
			}
		}
	}
	if i != len(dbBeers) {
		limitedDBBeers = dbBeers[i+1 : util.Min(i+1+int(maxCount), len(dbBeers))]
	}

	var beers []beer.Beer
	for _, limitedDBBeer := range limitedDBBeers {
		beers = append(beers, r.mapDBBeerToBeer(limitedDBBeer))
	}
	log.Printf("beerRepo - GetBeers return beer len : %+v", len(beers))
	return beers, nil
}

func (r *beerRepo) UpdateBeerRateAvg(beerID int64, rateAvg float64) error {
	res := r.db.Model(&DBBeer{}).Where("id = ?", beerID).Update("rate_avg", rateAvg)
	return res.Error
}

func (r *beerRepo) AddReview(review beer.Review) error {
	// Upsert Implementation
	// Update를 시도하고 RowsAffected로 분기 치면 안된다 (0일 경우 Create, 1일 경우 그냥 Return). Row 구성을 하나도 바꾸지 않을 경우 RowsAffected로 == 0으로 되어버림. RowsMatched가 없네 ...
	// 그렇다고 Update에서 isNotFound를 반환해주는 것도 아니라서 분기 치기 애매하다
	preReview, err := r.GetReviewByBeerIDAndUserID(review.BeerID, review.UserID)
	if err != nil {
		return err
	}

	if preReview == nil {
		// TODO 아래 쿼리들이 Transactional하게 이뤄져야함. 충돌나면 동시성 문제 생길 듯. Retry...
		dbReview := r.mapper.mapReviewToDBReview(review)
		res := r.db.Create(&dbReview)
		// https://github.com/go-gorm/gorm/issues/2903
		// Gorm V1에서는 Duplicate Error를 정의하지 않음
		if res.Error != nil && strings.Contains(res.Error.Error(), "Error 1062: Duplicate entry") {
			return errors.New("already added review")
		}
		res = r.db.Model(&DBBeer{}).Where("id = ?", review.BeerID).Update("review_count", gorm.Expr("review_count + ?", 1))
		return res.Error

	}
	res := r.db.Model(&DBReview{}).Where("user_id = ? AND beer_id = ?", review.UserID, review.BeerID).Updates(review)
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

func (r *beerRepo) GetReviewsByUserID(userID int64) ([]beer.Review, error) {
	query := DBReview{UserID: userID}
	var dbReviews []DBReview
	if err := r.db.Where(&query).Find(&dbReviews).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get review. user id : %v", userID))
		}
	}

	var reviews []beer.Review
	for _, dbReview := range dbReviews {
		reviews = append(reviews, r.mapper.mapDBReviewToReview(dbReview))
	}
	return reviews, nil
}

func (r *beerRepo) AddFavorite(favorite beer.Favorite) error {
	// Upsert Implementation
	preFavorite, err := r.getFavorite(favorite.BeerID, favorite.UserID)
	if err != nil {
		return err
	}

	if preFavorite == nil {
		dbFavorite := r.mapper.mapFavoriteToDBFavorite(favorite)
		res := r.db.Create(&dbFavorite)
		// https://github.com/go-gorm/gorm/issues/2903
		// Gorm V1에서는 Duplicate Error를 정의하지 않음
		if res.Error != nil && strings.Contains(res.Error.Error(), "Error 1062: Duplicate entry") {
			return errors.New("already added favorite")
		}
		return res.Error
	}
	res := r.db.Model(&DBFavorite{}).Where("beer_id = ? AND user_id = ?", favorite.BeerID, favorite.UserID).Update("flag", favorite.Flag)
	return res.Error
}

func (r *beerRepo) GetFavorites(userID int64) ([]beer.Favorite, error) {
	query := DBFavorite{UserID: userID}
	var dbFavorites []DBFavorite
	if err := r.db.Where(&query).Find(&dbFavorites).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get favorites. userID id: %v", userID))
		}
	}

	var favorites []beer.Favorite
	for _, dbFavorite := range dbFavorites {
		favorites = append(favorites, r.mapper.mapDBFavoriteToFavorite(dbFavorite))
	}
	return favorites, nil
}

func (r *beerRepo) getFavorite(beerID int64, userID int64) (*DBFavorite, error) {
	query := DBFavorite{BeerID: beerID, UserID: userID}
	var dbFavorite DBFavorite
	if err := r.db.Where(&query).First(&dbFavorite).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get favorites. beer id: %v, user id : %v", beerID, userID))
		}
	}
	return &dbFavorite, nil
}

func (r *beerRepo) AddUserBeerConfig(userBeerConfig beer.UserBeerConfig) error {
	// Upsert Implementation
	preUserBeerConfig, err := r.getUserBeerConfig(userBeerConfig.UserID)
	if err != nil {
		return err
	}

	dbUserBeerConfig := r.mapper.mapUserBeerConfigToDBUserBeerConfig(userBeerConfig)
	if preUserBeerConfig == nil {
		res := r.db.Create(&dbUserBeerConfig)
		// https://github.com/go-gorm/gorm/issues/2903
		// Gorm V1에서는 Duplicate Error를 정의하지 않음
		if res.Error != nil && strings.Contains(res.Error.Error(), "Error 1062: Duplicate entry") {
			return errors.New("already added favorite")
		}
		return res.Error
	}
	res := r.db.Model(&DBUserBeerConfig{}).Where("user_id = ?", dbUserBeerConfig.UserID).Updates(DBUserBeerConfig{AromaList: dbUserBeerConfig.AromaList, StyleList: dbUserBeerConfig.StyleList})
	return res.Error
}

func (r *beerRepo) GetBeerRanks(startDate time.Time, endDate time.Time, limit int64) ([]beer.BeerRank, error) {
	var beerRanks []beer.BeerRank
	res := r.db.Select("beer_id, count(*) as favorite_count").Table("beer_favorite").Where("created_at >= ? AND created_at <= ?", startDate, endDate).Group("beer_id").Order("count(*) desc").Limit(limit).Find(&beerRanks)
	log.Printf("rank data - %+v", spew.Sdump(beerRanks))

	if res.Error != nil {
		return nil, res.Error
	}
	return beerRanks, nil
}

func (r *beerRepo) getUserBeerConfig(userID int64) (*DBUserBeerConfig, error) {
	query := DBUserBeerConfig{UserID: userID}
	var dbUserBeerConfig DBUserBeerConfig
	if err := r.db.Where(&query).First(&dbUserBeerConfig).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get user beer config. user id : %v", userID))
		}
	}
	return &dbUserBeerConfig, nil

}

func (r *beerRepo) GetUserBeerConfig(userID int64) (*beer.UserBeerConfig, error) {
	dbUserBeerConfig, err := r.getUserBeerConfig(userID)
	if err != nil {
		return nil, err
	} else if dbUserBeerConfig == nil {
		return &beer.UserBeerConfig{}, nil // Returns DBUserConfig with default values
	}

	userBeerConfig := r.mapper.mapDBUserBeerConfigToUserBeerConfig(*dbUserBeerConfig)
	return &userBeerConfig, nil
}
