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
	return cont
}

func (cont *Controller) GetBeers(ctx echo.Context) error {
	// TODO GetBeers, GetBeer 중복 제거
	log.Printf("Controller - GetBeers() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.User()
	if err != nil {
		return err
	}
	log.Printf("Controller - GetBeers() - User %+v", spew.Sdump(user))

	var req dto.GetBeersRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeers() - Failed to bind %+v", err)
		return err
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

	var res dto.GetBeersResponse
	for idx, br := range beerList {
		log.Printf("Controller - GetBeers() Making dto for %+vth beer %+v", idx, spew.Sdump(br))
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
			dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer)
			dtoReview := cont.mapper.MapReviewToDTOReview(&review, reviewUser.NickName, dtoReducedBeer)
			dtoReviews = append(dtoReviews, *dtoReview)
		}

		if user != nil {
			reviewOwner, err := cont.beerUseCase.GetReviewByBeerIDAndUserID(br.ID, user.ID)
			if err != nil {
				return err
			}
			reviewOwnerBeer, err := cont.beerUseCase.GetBeer(reviewOwner.BeerID)
			if err != nil {
				return err
			}
			dtoReviewOwnerBeer := cont.mapper.MapBeerToDTReducedBeer(*reviewOwnerBeer)
			dtoReviewOwner = cont.mapper.MapReviewToDTOReview(reviewOwner, user.NickName, dtoReviewOwnerBeer)
		}
		dtoBeer := cont.mapper.MapBeerToDTOBeer(br, dtoReviews, dtoReviewOwner)
		res.Beers = append(res.Beers, dtoBeer)
	}

	// TODO 지금 Cursor 설정이 Controller에도, Repo에도 분포되어 있는 느낌인데 괜찮을까 고찰
	if len(beerList) != 0 {
		res.Cursor = &(res.Beers[len(beerList)-1].ID)
	}

	log.Printf("Controller - GetBeers() dto beer list %+v", res.Beers)

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
	user, err := _ctx.User()
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
		dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer)
		dtoReview := cont.mapper.MapReviewToDTOReview(&review, reviewUser.NickName, dtoReducedBeer)
		dtoReviews = append(dtoReviews, *dtoReview)
	}

	if user != nil {
		reviewOwner, err := cont.beerUseCase.GetReviewByBeerIDAndUserID(br.ID, user.ID)
		if err != nil {
			return err
		}
		reviewOwnerBeer, err := cont.beerUseCase.GetBeer(reviewOwner.BeerID)
		if err != nil {
			return err
		}
		dtoReviewOwnerBeer := cont.mapper.MapBeerToDTReducedBeer(*reviewOwnerBeer)
		dtoReviewOwner = cont.mapper.MapReviewToDTOReview(reviewOwner, user.NickName, dtoReviewOwnerBeer)
	}
	dtoBeer := cont.mapper.MapBeerToDTOBeer(*br, dtoReviews, dtoReviewOwner)

	relatedBeers, err := cont.beerUseCase.GetRelatedBeers(br.ID)
	if err != nil {
		return err
	}
	dtorRelatedBeers := cont.mapper.MapRelatedBeersToDTORelatedBeers(relatedBeers)

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
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if user == nil || user.ID == 0 {
		return fmt.Errorf("user not found")
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
			UserID:  user.ID,
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
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	} else if user == nil || user.ID == 0 {
		return fmt.Errorf("user not found")
	}

	reviews, err := cont.beerUseCase.GetReviewsByUserID(user.ID)
	if err != nil {
		return err
	}

	var dtoReviews []dto.Review
	for _, review := range reviews {
		beer, err := cont.beerUseCase.GetBeer(review.BeerID)
		if err != nil {
			return err
		}
		dtoReducedBeer := cont.mapper.MapBeerToDTReducedBeer(*beer)
		dtoReview := cont.mapper.MapReviewToDTOReview(&review, user.NickName, dtoReducedBeer)
		dtoReviews = append(dtoReviews, *dtoReview)
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": dtoReviews,
		},
	)
}
