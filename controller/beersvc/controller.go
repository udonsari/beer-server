package beersvc

import (
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/beersvc/dto"
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

// TODO Add Logger

type Controller struct {
	controller.Base
	beerUseCase beer.UseCase
	mapper      dto.Mapper
}

func NewController(engine *echo.Echo, beerUseCase beer.UseCase) Controller {
	con := Controller{
		beerUseCase: beerUseCase,
		mapper:      dto.NewMapper(),
	}
	engine.GET("/api/beers", con.GetBeers)
	return con
}

func (controller *Controller) GetBeers(ctx echo.Context) error {
	log.Printf("Controller - GetBeers() - Controller")

	var req dto.GetBeersRequest
	if err := controller.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeers() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - GetBeers() - Param %+v", spew.Sdump(req))

	args, err := controller.mapper.MapGetBeersRequestToBeerQueryArgs(req)
	if err != nil {
		return err
	}

	beerList, err := controller.beerUseCase.GetBeers(*args)
	if err != nil {
		return err
	}

	var res dto.GetBeersResponse
	for _, beer := range beerList {
		dtoBeer := controller.mapper.MapBeerToDTOBeer(beer)
		res.Beers = append(res.Beers, dtoBeer)
	}

	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": res,
		},
	)
}
