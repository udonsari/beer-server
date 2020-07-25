package beersvc

import (
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/beersvc/dto"
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/labstack/echo"
)

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
	log.Printf("GetBeers() - Controller")
	var req dto.GetBeersRequest
	if err := controller.Bind(ctx, &req); err != nil {
		log.Printf("GetBeers() - Failed to bind %+v", err)
		return err
	}
	log.Printf("GetBeers() - Param %+v", req)

	beerList, err := controller.beerUseCase.GetBeers()
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
