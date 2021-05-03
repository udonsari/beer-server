package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller/beersvc"
	"github.com/UdonSari/beer-server/controller/usersvc"
	"github.com/UdonSari/beer-server/domain/beer"
	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/domain/user"
	userRepo "github.com/UdonSari/beer-server/domain/user/repo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server interface {
	Init()
	Start()
}

type serverImpl struct {
	_engine     *echo.Echo
	_server     *http.Server
	beerUseCase beer.UseCase
	userUseCase user.UseCase
	host        string
	port        string
	errorMapper errorMapper
}

func (s *serverImpl) Init() {
	log.Printf("# server initialization starts ...")
	engine := s.engine()
	s.registerRoute(engine)
}

func (s *serverImpl) Start() {
	s._server = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.port),
		Handler: s._engine,
	}

	log.Printf("# server up starts at port %+v ...", s.port)

	if err := s._server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("# server failed with err : %+v", err)
	}
}

func (s *serverImpl) engine() *echo.Echo {
	d := NewDependency()

	s.port = d.PortStr()
	s.host = fmt.Sprintf("%s:%s", d.Host(), s.port)

	beerRepo := beerRepo.New(d.MysqlDB(true), d.BeerCacheDuration())
	s.beerUseCase = beer.NewUseCase(beerRepo)
	userRepo := userRepo.New(d.MysqlDB(true))
	s.userUseCase = user.NewUseCase(userRepo, s.host, s.port, d.ServerEnv(), d.DevToken())

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
	s._engine.Static("/static", "static")

	// Error Middleware. Care all domain errors
	s._engine.HTTPErrorHandler = func(err error, c echo.Context) {
		log.Printf(c.Path(), err.Error())

		retErr := echo.HTTPError{
			Code:    s.errorMapper.fromDomainErrorToStatusCode(err),
			Message: err.Error(),
		}
		s._engine.DefaultHTTPErrorHandler(&retErr, c)
	}

	return s._engine
}

func (s *serverImpl) registerRoute(engine *echo.Echo) {
	beersvc.NewController(engine, s.beerUseCase, s.userUseCase)
	usersvc.NewController(engine, s.userUseCase, s.host)
}

func New() Server {
	return &serverImpl{}
}
