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
	return cont
}

func (cont *Controller) GetBeers(ctx echo.Context) error {
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
		var reviewOwner *beer.Review
		reviews, err := cont.beerUseCase.GetReviews(br.ID)
		if err != nil {
			return err
		}
		if user != nil {
			reviewOwner, err = cont.beerUseCase.GetReviewByBeerIDAndUserID(br.ID, user.ID)
			if err != nil {
				return err
			}
		}
		dtoBeer := cont.mapper.MapBeerToDTOBeer(br, reviews, reviewOwner)
		res.Beers = append(res.Beers, dtoBeer)
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
	var reviewOwner *beer.Review
	reviews, err := cont.beerUseCase.GetReviews(br.ID)
	if err != nil {
		return err
	}
	if user != nil {
		reviewOwner, err = cont.beerUseCase.GetReviewByBeerIDAndUserID(br.ID, user.ID)
		if err != nil {
			return err
		}
	}
	dtoBeer := cont.mapper.MapBeerToDTOBeer(*br, reviews, reviewOwner)

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
