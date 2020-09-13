package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/UdonSari/beer-server/controller/beersvc"
	"github.com/UdonSari/beer-server/controller/usersvc"
	"github.com/UdonSari/beer-server/domain/beer"
	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/domain/user"
	userRepo "github.com/UdonSari/beer-server/domain/user/repo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var PORT int
var PORT_STR string
var HOST string

type Server interface {
	Init()
	Start()
}

type serverImpl struct {
	_engine     *echo.Echo
	_server     *http.Server
	beerUseCase beer.UseCase
	userUseCase user.UseCase
}

func (s *serverImpl) Init() {
	var ok bool
	PORT_STR, ok = os.LookupEnv("PORT")
	if !ok {
		log.Printf("failed to find port in env so set 8081")
		PORT = 8081
		PORT_STR = "8081"
	} else {
		var err error
		PORT, err = strconv.Atoi(PORT_STR)
		if err != nil {
			log.Printf("failed to parse port %+v err %+v", PORT, err)
			os.Exit(1)
		}
	}
	HOST = "http://127.0.0.1:" + PORT_STR

	log.Printf("# server initialization starts ...")
	engine := s.engine()
	s.registerRoute(engine)
}

func (s *serverImpl) Start() {
	s._server = &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: s._engine,
	}

	log.Printf("# server up starts at port %+v ...", PORT)

	if err := s._server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("# server failed with err : %+v", err)
	}
}

func (s *serverImpl) engine() *echo.Echo {
	d := NewDependency()

	beerRepo := beerRepo.New(d.MysqlDB(), d.BeerCacheDuration())
	s.beerUseCase = beer.NewUseCase(beerRepo)
	userRepo := userRepo.New(d.MysqlDB())
	s.userUseCase = user.NewUseCase(userRepo, HOST, PORT_STR)

	if s._engine != nil {
		return s._engine
	}
	s._engine = echo.New()
	s._engine.Use(middleware.Recover())
	s._engine.Use(middleware.CORS())
	s._engine.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) error {
				cc := &CustomContext{ctx, s.userUseCase}
				return next(cc)
			}
		},
	)

	s._engine.HTTPErrorHandler = func(err error, c echo.Context) {
		log.Printf(c.Path(), err.Error())
		retErr := echo.HTTPError{
			Code:    http.StatusInternalServerError, // Currently all error returns with 500
			Message: err.Error(),
		}
		s._engine.DefaultHTTPErrorHandler(&retErr, c)
	}

	return s._engine
}

func (s *serverImpl) registerRoute(engine *echo.Echo) {
	beersvc.NewController(engine, s.beerUseCase, s.userUseCase)
	usersvc.NewController(engine, s.userUseCase, HOST)
}

func New() Server {
	return &serverImpl{}
}
