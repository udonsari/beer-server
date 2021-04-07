package beersvc

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
	engine.GET("/api/random-beers", cont.GetRandomBeers)
	engine.POST("/api/review", cont.AddReview)
	engine.GET("/api/review", cont.GetReview)
	engine.GET("/api/app-config", cont.GetAppConfig)
	engine.POST("/api/favorite", cont.AddFavorite)
	engine.GET("/api/favorite", cont.GetFavorites)
	engine.POST("/api/user-beer-config", cont.AddUserBeerConfig)
	engine.GET("/api/user-beer-config", cont.GetUserBeerConfig)
	engine.GET("/api/popular-beers", cont.GetPopularBeers)
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
	dtoBeer := cont.mapper.MapBeerToDTOBeer(*br, dtoReviews, dtoReviewOwner, favoriteMap[br.ID])

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

func (cont *Controller) GetRandomBeers(ctx echo.Context) error {
	log.Printf("Controller - GetRandomBeers() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	}
	log.Printf("Controller - GetRandomBeers() - User %+v", spew.Sdump(user))

	beerList, err := cont.beerUseCase.GetRandomBeers()
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

	// GetRandomBeers response's cursor is nil
	var res dto.GetBeersResponse
	for idx, br := range beerList {
		log.Printf("Controller - GetRandomBeers() Making dto for %+vth beer %+v", idx, spew.Sdump(br))
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

	log.Printf("Controller - GetRandomBeers() dto beer list %+v", res.ReducedBeer)

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
	// TODO Add Semantic Versioning
	version := ctx.Request().Header.Get("AppVersion")
	log.Printf("Controller - GetAppConfig() - AppVersion %+v", spew.Sdump(version))
	if version == "" {
		return ctx.JSON(
			http.StatusOK,
			map[string]interface{}{
				"result": cont.getAppConfigV1(),
			},
		)
	}
	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": cont.getAppConfigV2(),
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

func (cont *Controller) AddUserBeerConfig(ctx echo.Context) error {
	log.Printf("Controller - AddUserBeerConfig() - Controller")
	_ctx := ctx.(controller.CustomContext)
	usr, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if usr == nil || usr.ID == 0 {
		return user.UserNotFoundError{}
	}

	var req dto.AddUserBeerConfig
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - AddUserBeerConfig() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - AddUserBeerConfig() - Param %+v", spew.Sdump(req))

	err = cont.beerUseCase.AddUserBeerConfig(
		beer.UserBeerConfig{
			UserID: usr.ID,
			Aroma:  req.Aroma,
			Style:  req.Style,
		},
	)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (cont *Controller) GetUserBeerConfig(ctx echo.Context) error {
	log.Printf("Controller - GetUserBeerConfig() - Controller")
	_ctx := ctx.(controller.CustomContext)
	usr, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if usr == nil || usr.ID == 0 {
		return user.UserNotFoundError{}
	}

	userBeerConfig, err := cont.beerUseCase.GetUserBeerConfig(usr.ID)
	if err != nil {
		return err
	}
	dtoUserBeerConfig := cont.mapper.MapUserBeerConfigToDTOUserBeerConfig(*userBeerConfig)

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": dtoUserBeerConfig,
		},
	)
}

func (cont *Controller) GetPopularBeers(ctx echo.Context) error {
	// Cursor does not set in this api
	log.Printf("Controller - GetPopularBeers() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	}
	log.Printf("Controller - GetPopularBeers() - User %+v", spew.Sdump(user))

	var req dto.GetPopularBeersRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetPopularBeers() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - GetPopularBeers() - Param %+v", spew.Sdump(req))

	limit := int64(5)
	var startDate, endDate time.Time

	if req.Limit != nil {
		limit = int64(*req.Limit)
	}

	// TODO Add Timezone Concept
	if req.StartDate != nil && req.EndDate != nil {
		const timeLayout = "2006-01-01 15:04:05"
		startDate, err = time.Parse(timeLayout, *req.StartDate)
		if err != nil {
			return err
		}
		endDate, err = time.Parse(timeLayout, *req.EndDate)
		if err != nil {
			return err
		}
	} else {
		// If startDate, endDate are not set then, set period as current month
		// Maybe this can have problem in day 1. There can be no favorite for beer in day 1 so this would return no beers
		cur := time.Now()
		startDate = time.Date(cur.Year(), cur.Month(), 1, 0, 0, 0, 0, time.UTC)
		endDate = cur
	}

	beerList, err := cont.beerUseCase.GetPopularBeers(startDate, endDate, limit)
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
		log.Printf("Controller - GetPopularBeers() Making dto for %+vth beer %+v", idx, spew.Sdump(br))
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

	log.Printf("Controller - GetPopularBeers() dto beer list %+v", res.ReducedBeer)

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": res,
		},
	)
}

func (cont *Controller) getAppConfigV1() dto.AppConfigV1 {
	return dto.AppConfigV1{
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

func (cont *Controller) getAppConfigV2() dto.AppConfigV2 {
	return dto.AppConfigV2{
		AromaList: []string{
			"Malty", "Caramel", "Roast", "Coffee", "Grass", "Banana", "Apple", "Peach", "Mango", "Orange", "Spicy", "Vinegar", "Nutty", "Pineapple", "Melon", "Blackberry", "Chocolate", "Cherry", "Lemon", "Passion Fruit", "Grapefruit",
		},
		CountryList: []string{
			"USA", "Begium", "Genmany", "Korea", "UK", "Czech", "France",
		},
		BeerStyleList: []interface{}{
			map[string]interface{}{
				"big_name": "Ale",
				"mid_categories": []interface{}{
					map[string]interface{}{
						"mid_name":    "Ale",
						"description": "상면 발효 효모를 사용하여\n화려하고 풍부한 향이 나는 맥주",
						"small_categories": []string{
							"Ale", "Abbey Ale", "Amber Ale", "American Pale Ale", "Brown Belgian Strong Ale", "Blonde Ale", "Brown Ale", "Saison", "Golden Ale", "Hop Ale", "Irish Ale", "Light Ale", "Old Ale", "Pale Ale", "Quadrupel Ale", "Red Ale", "Sparkling Ale", "Summer Ale", "Trappist Ale", "Tripel Ale", "White Ale", "Wheat Ale", "Wit Ale", "Barley Wine", "Dubbel Ale", "Dark Ale", "Wild Ale", "Pumpkin Ale",
						},
					},
					map[string]interface{}{
						"mid_name":    "IPA",
						"description": "페일 에일에 다량의 홉을 넣은,\n홉의 쌉쌀한 향과 맛이 매력적인 맥주",
						"small_categories": []string{
							"IPA", "American IPA", "Black IPA", "Belgian IPA", "Double IPA", "Hazy IPA", "Imperial IPA", "Rye IPA", "Session IPA", "Sour IPA", "Smoothie IPA", "Wheat IPA",
						},
					},
					map[string]interface{}{
						"mid_name":    "Dark Beer",
						"description": "로스팅된 맥아를 사용한 어두운 색상의 맥주로\n풍부한 바디감이 특징인 맥주",
						"small_categories": []string{
							"Dark Beer", "Porter", "Stout", "Baltic Porter", "Bourbon County Stout", "Imperial Porter", "Imperial Stout", "Irish Stout", "Sweet Stout", "Schwarz", "Milk Stout",
						},
					},
					map[string]interface{}{
						"mid_name":    "Wheat Beer",
						"description": "밀 맥아를 높은 비율로 사용한 맥주로\n부드럽고 달콤한 향이 특징인 맥주",
						"small_categories": []string{
							"Wheat Beer", "Belgian White", "Hefeweizen", "Witbier", "Weizen", "Dunkel Weizen", "Weisse",
						},
					},
				},
			},
			map[string]interface{}{
				"big_name": "Larger",
				"mid_categories": []interface{}{
					map[string]interface{}{
						"mid_name":    "Larger",
						"description": "하면 발효 효모를 사용하여\n가벼운 풍미와 시원한 청량감이 매력적인 맥주",
						"small_categories": []string{
							"Lager", "Amber Lager", "Dark Lager", "Helles Lager", "India Pale Lager", "Pale Lager", "Rauchbier", "Kellerbier", "Marzen", "Dunkel",
						},
					},
					map[string]interface{}{
						"mid_name":    "Bock",
						"description": "다양한 원료와 긴 발효기간을 거쳐\n풍부한 맛과 높은 도수를 가진 맥주",
						"small_categories": []string{
							"Bock", "Weizen Bock", "Double Bock", "MaiBock",
						},
					},
				},
			},
			map[string]interface{}{
				"big_name": "Lambic",
				"mid_categories": []interface{}{
					map[string]interface{}{
						"mid_name":    "Lambic",
						"description": "상큼한 맛이 매력적인 자연 발효 맥주",
						"small_categories": []string{
							"Lambic", "Gueuze",
						},
					},
				},
			},
			map[string]interface{}{
				"big_name": "etc",
				"mid_categories": []interface{}{
					map[string]interface{}{
						"mid_name":    "etc",
						"description": "비어있다에서 다양한 맥주를 만나보세요",
						"small_categories": []string{
							"Radler", "Cider", "Gose", "Gluten Free", "Kolsch", "Low Alcohol", "Ginger Beer",
						},
					},
				},
			},
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
