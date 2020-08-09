package beersvc

import (
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/beersvc/dto"
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

// TODO * Add Logger

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
	engine.POST("/api/rate", cont.AddRate)
	engine.POST("/api/comment", cont.AddComment)
	return cont
}

func (cont *Controller) GetBeers(ctx echo.Context) error {
	log.Printf("Controller - GetBeers() - Controller")

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
	for _, beer := range beerList {
		comments, err := cont.beerUseCase.GetComments(beer.ID)
		if err != nil {
			return err
		}
		rates, err := cont.beerUseCase.GetRates(beer.ID)
		if err != nil {
			return err
		}
		dtoBeer := cont.mapper.MapBeerToDTOBeer(beer, comments, rates)
		res.Beers = append(res.Beers, dtoBeer)
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": res,
		},
	)
}

func (cont *Controller) GetBeer(ctx echo.Context) error {
	log.Printf("Controller - GetBeer() - Controller")

	var req dto.GetBeerRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeer() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - GetBeer() - Param %+v", spew.Sdump(req))

	beer, err := cont.beerUseCase.GetBeer(req.BeerID)
	if err != nil {
		return err
	}

	var res dto.GetBeerResponse
	comments, err := cont.beerUseCase.GetComments(beer.ID)
	if err != nil {
		return err
	}
	rates, err := cont.beerUseCase.GetRates(beer.ID)
	if err != nil {
		return err
	}
	dtoBeer := cont.mapper.MapBeerToDTOBeer(*beer, comments, rates)
	res.Beer = dtoBeer

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": res,
		},
	)
}

func (cont *Controller) AddRate(ctx echo.Context) error {
	log.Printf("Controller - AddRate() - Controller")
	accessTokens := ctx.Request().Header["Authorization"]
	if len(accessTokens) < 1 {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	user, err := cont.userUseCase.GetUser(accessTokens[0])
	if err != nil {
		return err
	}

	var req dto.AddRateRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - AddRate() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - AddRate() - Param %+v", spew.Sdump(req))

	err = cont.beerUseCase.AddRate(req.BeerID, req.Ratio, user.ID)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (cont *Controller) AddComment(ctx echo.Context) error {
	log.Printf("Controller - AddComment() - Controller")
	accessTokens := ctx.Request().Header["Authorization"]
	if len(accessTokens) < 1 {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	user, err := cont.userUseCase.GetUser(accessTokens[0])
	if err != nil {
		return err
	}

	var req dto.AddCommentRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - AddComment() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - AddComment() - Param %+v", spew.Sdump(req))

	err = cont.beerUseCase.AddComment(req.BeerID, req.Content, user.ID)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
