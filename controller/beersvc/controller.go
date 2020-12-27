package beersvc

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/beersvc/dto"
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

// UseCase는 아무것도 안하고, Controller가 뭔가를 많이 하는 것 처럼 보이나, 사실 Response 만들기에 불과하므로 우선 OK

type Controller struct {
	controller.Base
	beerUseCase beer.UseCase
	userUseCase user.UseCase
	mapper      dto.Mapper
}

func NewController(engine *echo.Echo, beerUseCase beer.UseCase, userUseCase user.UseCase) Controller {
	cont := Controller{
		beerUseCase: beerUseCase,
		userUseCase: userUseCase,
		mapper:      dto.NewMapper(),
	}
	engine.GET("/api/beers", cont.GetBeers)
	engine.GET("/api/beer", cont.GetBeer)
	engine.POST("/api/review", cont.AddReview)
	engine.GET("/api/review", cont.GetReview)
	engine.GET("/api/app-config", cont.GetAppConfig)
	engine.POST("/api/favorite", cont.AddFavorite)
	engine.GET("/api/favorite", cont.GetFavorites)
	return cont
}

func (cont *Controller) GetBeers(ctx echo.Context) error {
	log.Printf("Controller - GetBeers() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	}
	log.Printf("Controller - GetBeers() - User %+v", spew.Sdump(user))

	var req dto.GetBeersRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeers() - Failed to bind %+v", err)
		return err
	}
	if err = cont.isValidSortBy(req.SortBy); err != nil {
		return err
	}
	if (req.MinABV != nil && req.MaxABV == nil) || (req.MinABV == nil && req.MaxABV != nil) {
		return InvalidArgsError{Message: "MinABV and MaxABV should come together"}
	}

	log.Printf("Controller - GetBeers() - Param %+v", spew.Sdump(req))

	args, err := cont.mapper.MapGetBeersRequestToBeerQueryArgs(req)
	if err != nil {
		return err
	}

	beerList, err := cont.beerUseCase.GetBeers(*args)
	if err != nil {
		return err
	}

	favoriteList, err := cont.beerUseCase.GetFavorites(user.ID)
	if err != nil {
		return err
	}
	favoriteMap := make(map[int64]bool)
	for _, favorite := range favoriteList {
		// If none, flag would be initial false
		favoriteMap[favorite.BeerID] = favorite.Flag
	}

	var res dto.GetBeersResponse
	for idx, br := range beerList {
		log.Printf("Controller - GetBeers() Making dto for %+vth beer %+v", idx, spew.Sdump(br))
		var dtoReviews []dto.Review
		reviews, err := cont.beerUseCase.GetReviews(br.ID)
		if err != nil {
			return err
		}
		for _, review := range reviews {
			reviewUser, err := cont.userUseCase.GetUserByID(review.UserID)
			if err != nil {
				return err
			}
			beer, err := cont.beerUseCase.GetBeer(review.BeerID)
			if err != nil {
				return err
			}
			dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer, favoriteMap[br.ID])
			dtoReview := cont.mapper.MapReviewToDTOReview(review, reviewUser.NickName, dtoReducedBeer)
			dtoReviews = append(dtoReviews, *dtoReview)
		}

		dtoBeer := cont.mapper.MapBeerToDTReducedBeer(br, favoriteMap[br.ID])
		res.ReducedBeer = append(res.ReducedBeer, dtoBeer)
	}

	// TODO 지금 Cursor 설정이 Controller에도, Repo에도 분포되어 있는 느낌인데 괜찮을까 고찰
	if len(beerList) != 0 {
		args.MaxCount++
		lastPageCheckBeersList, err := cont.beerUseCase.GetBeers(*args)
		if err != nil {
			return err
		}

		// 만약 maxCount + 1개 로 제한을 걸고 쿼리했을때 이전과 같은 개수의 맥주가 나온다면 마지막 페이지
		if len(lastPageCheckBeersList) != len(beerList) {
			res.Cursor = &beerList[len(beerList)-1].ID
		}
	}

	log.Printf("Controller - GetBeers() dto beer list %+v", res.ReducedBeer)

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": res,
		},
	)
}

func (cont *Controller) GetBeer(ctx echo.Context) error {
	log.Printf("Controller - GetBeer() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	}
	log.Printf("Controller - GetBeer() - User %+v", spew.Sdump(user))

	var req dto.GetBeerRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeer() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - GetBeer() - Param %+v", spew.Sdump(req))

	br, err := cont.beerUseCase.GetBeer(req.BeerID)
	if err != nil {
		return err
	}

	favoriteList, err := cont.beerUseCase.GetFavorites(user.ID)
	if err != nil {
		return err
	}
	favoriteMap := make(map[int64]bool)
	for _, favorite := range favoriteList {
		// If none, flag would be initial false
		favoriteMap[favorite.BeerID] = favorite.Flag
	}

	var res dto.GetBeerResponse
	var dtoReviewOwner *dto.Review
	var dtoReviews []dto.Review
	reviews, err := cont.beerUseCase.GetReviews(br.ID)
	if err != nil {
		return err
	}
	for _, review := range reviews {
		reviewUser, err := cont.userUseCase.GetUserByID(review.UserID)
		if err != nil {
			return err
		}
		beer, err := cont.beerUseCase.GetBeer(review.BeerID)
		if err != nil {
			return err
		}
		dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer, favoriteMap[beer.ID])
		dtoReview := cont.mapper.MapReviewToDTOReview(review, reviewUser.NickName, dtoReducedBeer)
		dtoReviews = append(dtoReviews, *dtoReview)
	}

	if user != nil {
		reviewOwner, err := cont.beerUseCase.GetReviewByBeerIDAndUserID(br.ID, user.ID)
		if err != nil {
			return err
		}
		if reviewOwner != nil {
			reviewOwnerBeer, err := cont.beerUseCase.GetBeer(reviewOwner.BeerID)
			if err != nil {
				return err
			}
			dtoReviewOwnerBeer := cont.mapper.MapBeerToDTReducedBeer(*reviewOwnerBeer, favoriteMap[br.ID])
			dtoReviewOwner = cont.mapper.MapReviewToDTOReview(*reviewOwner, user.NickName, dtoReviewOwnerBeer)
		}
	}
	dtoBeer := cont.mapper.MapBeerToDTOBeer(*br, dtoReviews, dtoReviewOwner)

	relatedBeers, err := cont.beerUseCase.GetRelatedBeers(br.ID)
	if err != nil {
		return err
	}
	dtorRelatedBeers := cont.mapper.MapRelatedBeersToDTORelatedBeers(relatedBeers, favoriteMap)

	res.Beer = dtoBeer
	res.RelatedBeers = dtorRelatedBeers

	log.Printf("Controller - GetBeer() dto beer %+v", res)

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": res,
		},
	)
}

func (cont *Controller) AddReview(ctx echo.Context) error {
	log.Printf("Controller - AddReview() - Controller")
	_ctx := ctx.(controller.CustomContext)
	usr, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if usr == nil || usr.ID == 0 {
		return user.UserNotFoundError{}
	}

	var req dto.AddReviewRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - AddReview() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - AddReview() - Param %+v", spew.Sdump(req))

	err = cont.beerUseCase.AddReview(
		beer.Review{
			BeerID:  req.BeerID,
			Content: req.Content,
			Ratio:   req.Ratio,
			UserID:  usr.ID,
		},
	)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (cont *Controller) GetReview(ctx echo.Context) error {
	log.Printf("Controller - GetReview() - Controller")
	_ctx := ctx.(controller.CustomContext)
	usr, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if usr == nil || usr.ID == 0 {
		return user.UserNotFoundError{}
	}

	reviews, err := cont.beerUseCase.GetReviewsByUserID(usr.ID)
	if err != nil {
		return err
	}

	favoriteList, err := cont.beerUseCase.GetFavorites(usr.ID)
	if err != nil {
		return err
	}
	favoriteMap := make(map[int64]bool)
	for _, favorite := range favoriteList {
		// If none, flag would be initial false
		favoriteMap[favorite.BeerID] = favorite.Flag
	}

	var dtoReviews []dto.Review
	for _, review := range reviews {
		beer, err := cont.beerUseCase.GetBeer(review.BeerID)
		if err != nil {
			return err
		}
		dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer, favoriteMap[beer.ID])
		dtoReview := cont.mapper.MapReviewToDTOReview(review, usr.NickName, dtoReducedBeer)
		dtoReviews = append(dtoReviews, *dtoReview)
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": dtoReviews,
		},
	)
}

func (cont *Controller) GetAppConfig(ctx echo.Context) error {
	// TODO Currently dummy config, later change to all aroma, country, beer style in DB
	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": cont.getDummyAppConfig(),
		},
	)
}

func (cont *Controller) AddFavorite(ctx echo.Context) error {
	log.Printf("Controller - AddFavorite() - Controller")
	_ctx := ctx.(controller.CustomContext)
	usr, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if usr == nil || usr.ID == 0 {
		return user.UserNotFoundError{}
	}

	var req dto.AddFavoriteRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - AddFavorite() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - AddFavorite() - Param %+v", spew.Sdump(req))

	err = cont.beerUseCase.AddFavorite(
		beer.Favorite{
			BeerID: req.BeerID,
			Flag:   req.Flag,
			UserID: usr.ID,
		},
	)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (cont *Controller) GetFavorites(ctx echo.Context) error {
	log.Printf("Controller - GetFavorites() - Controller")
	_ctx := ctx.(controller.CustomContext)
	usr, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if usr == nil || usr.ID == 0 {
		return user.UserNotFoundError{}
	}

	favorites, err := cont.beerUseCase.GetFavorites(usr.ID)
	if err != nil {
		return err
	}

	var dtoFavorites []dto.Favorite
	for _, favorite := range favorites {
		if !favorite.Flag {
			continue
		}

		beer, err := cont.beerUseCase.GetBeer(favorite.BeerID)
		if err != nil {
			return err
		}
		dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer, favorite.Flag)
		dtoFavorite := cont.mapper.MapFavoriteToDTOFavorite(favorite, dtoReducedBeer)
		dtoFavorites = append(dtoFavorites, *dtoFavorite)
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": dtoFavorites,
		},
	)

}

// TODO Replac real aroma in db
func (cont *Controller) getDummyAppConfig() dto.AppConfig {
	return dto.AppConfig{
		AromaList: []string{
			"Malty", "Caramel", "Roast", "Coffee", "Grass", "Banana", "Apple", "Peach", "Mango", "Orange", "Spicy", "Vinegar", "Nutty", "Pineapple", "Melon", "Blackberry", "Chocolate", "Cherry", "Lemon", "Passion Fruit", "Grapefruit",
		},
		CountryList: []string{
			"USA", "Begium", "Genmany", "Korea", "UK", "Czech", "France",
		},
		BeerStyleList: []string{
			"Porter", "Stout", "Pilsener", "Light Lager", "Scotch Ale", "Saison", "Pale Ale", "Brown Ale", "India Pale Ale", "Gose", "Quadrupel", "Tripel", "Lambic", "Rye Amber", "Kolsch", "Witbier", "Red Ale", "New England IPA", "Sour Ale", "ETC",
		},
		MinABV: 0.0,
		MaxABV: 15.0,
	}
}

func (cont *Controller) isValidSortBy(val *string) error {
	if val == nil {
		return nil
	}
	if *val == beer.SortByRateAvgAsc || *val == beer.SortByRateAvgDesc || *val == beer.SortByReviewCountAsc || *val == beer.SortByReviewCountDesc {
		return nil
	}
	return InvalidArgsError{fmt.Sprintf("invalid sortBy %+v", *val)}
}
