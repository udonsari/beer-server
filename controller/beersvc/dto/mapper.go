package dto

import (
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/util"
)

const (
	DefaultMaxCount = int64(20)
	DefaultCursor   = int64(0)
)

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) MapBeerToDTOBeer(beer beer.Beer, dtoReviews []Review, dtoReviewOwner *Review, favoriteFlag bool) Beer {
	// TODO 이미 밖에서 DTO를 만들어서 다시 Map한다는게 이상
	return Beer{
		ID:             beer.ID,
		Name:           beer.Name,
		Brewery:        beer.Brewery,
		ABV:            util.Floor(beer.ABV, 2),
		Country:        beer.Country,
		BeerStyle:      beer.BeerStyle,
		Aroma:          beer.Aroma,
		ImageURL:       beer.ImageURL,
		ThumbnailImage: beer.ThumbnailImage,
		Reviews:        dtoReviews,
		RateAvg:        util.Floor(beer.RateAvg, 2),
		ReviewOwner:    dtoReviewOwner,
		ReviewCount:    beer.ReviewCount,
		FavoriteFlag:   favoriteFlag,
	}
}

func (m *Mapper) MapBeerToDTReducedBeer(beer beer.Beer, favoriteFlag bool) ReducedBeer {
	return ReducedBeer{
		ID:             beer.ID,
		Name:           beer.Name,
		Brewery:        beer.Brewery,
		ABV:            util.Floor(beer.ABV, 2),
		Country:        beer.Country,
		BeerStyle:      beer.BeerStyle,
		Aroma:          beer.Aroma,
		ThumbnailImage: beer.ThumbnailImage,
		RateAvg:        util.Floor(beer.RateAvg, 2),
		ReviewCount:    beer.ReviewCount,
		FavoriteFlag:   favoriteFlag,
	}
}

// TODO 여기에 Favorite Flag를 넣는게 이상할 수 있다. 어디는 mapper 밖에서 처리하고, 어디는 넘겨서 mapper에서 처리하고
// favoriteMap : [beer, favoriteFlag]
func (m *Mapper) MapRelatedBeersToDTORelatedBeers(relatedBeer *beer.RelatedBeers, favoriteMap map[int64]bool) *RelatedBeers {
	if relatedBeer == nil {
		return nil
	}
	var dtoRelatedBeers RelatedBeers
	for _, b := range relatedBeer.AromaRelatedBeer {
		dtoRelatedBeers.AromaRelatedBeer = append(dtoRelatedBeers.AromaRelatedBeer, m.MapBeerToDTReducedBeer(b, favoriteMap[b.ID]))
	}
	for _, b := range relatedBeer.StyleRelatedBeer {
		dtoRelatedBeers.StyleRelatedBeer = append(dtoRelatedBeers.StyleRelatedBeer, m.MapBeerToDTReducedBeer(b, favoriteMap[b.ID]))
	}
	for _, b := range relatedBeer.RandomlyRelatedBeer {
		dtoRelatedBeers.RandomlyRelatedBeer = append(dtoRelatedBeers.RandomlyRelatedBeer, m.MapBeerToDTReducedBeer(b, favoriteMap[b.ID]))
	}
	return &dtoRelatedBeers
}

func (m *Mapper) MapReviewToDTOReview(review beer.Review, nickName string, beer ReducedBeer) *Review {
	return &Review{
		ReducedBeer: beer,
		Content:     review.Content,
		Ratio:       util.Floor(review.Ratio, 2),
		UserID:      review.UserID,
		NickName:    nickName,
		CreatedAt:   review.CreatedAt,
	}
}

func (m *Mapper) MapFavoriteToDTOFavorite(favorite beer.Favorite, beer ReducedBeer) *Favorite {
	return &Favorite{
		ReducedBeer: beer,
		UserID:      favorite.UserID,
		BeerID:      favorite.BeerID,
	}
}

func (m *Mapper) MapUserBeerConfigToDTOUserBeerConfig(userBeerConfig beer.UserBeerConfig) UserBeerConfig {
	return UserBeerConfig{
		Aroma: userBeerConfig.Aroma,
		Style: userBeerConfig.Style,
	}
}

func (m *Mapper) MapGetBeersRequestToBeerQueryArgs(req GetBeersRequest) (*beer.BeerQueryArgs, error) {
	var args beer.BeerQueryArgs
	if req.MinABV != nil {
		args.ABVInterval = &beer.ABVInterval{
			MinABV: *req.MinABV,
			MaxABV: *req.MaxABV,
		}
	}

	args.Name = req.Name
	args.Country = req.Country
	args.BeerStyle = req.BeerStyle
	args.Aroma = req.Aroma

	if req.Cursor == nil {
		args.Cursor = DefaultCursor
	} else {
		args.Cursor = *req.Cursor
	}

	if req.MaxCount == nil {
		args.MaxCount = DefaultMaxCount
	} else {
		args.MaxCount = *req.MaxCount
	}
	args.SortBy = req.SortBy

	return &args, nil
}
