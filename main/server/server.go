package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller/beersvc"
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TODO Use env. maybe viper ?
const port = 8081

type Server interface {
	Init()
	Start()
}

type server struct {
	_engine *echo.Echo
	_server *http.Server
}

func (s *server) Init() {
	log.Printf("# server initialization starts ...")
	engine := s.engine()
	s.registerRoute(engine)
}

func (s *server) Start() {
	log.Printf("# server up starts ...")

	s._server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s._engine,
	}

	if err := s._server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("server failed with err : %+v", err)
	}
}

func (s *server) engine() *echo.Echo {
	if s._engine != nil {
		return s._engine
	}
	s._engine = echo.New()
	s._engine.Use(middleware.Recover())
	s._engine.Use(middleware.CORS())

	return s._engine
}

func (s *server) registerRoute(engine *echo.Echo) {
	beerRepo := repo.New()
	beerUseCase := beer.NewUseCase(beerRepo)

	beersvc.NewController(engine, beerUseCase)
}

func New() Server {
	return &server{}
}
