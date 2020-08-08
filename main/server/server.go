package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller"
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
		Addr:    fmt.Sprintf(":%d", controller.PORT),
		Handler: s._engine,
	}

	if err := s._server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("# server failed with err : %+v", err)
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
	beerRepo := beerRepo.New()
	beerUseCase := beer.NewUseCase(beerRepo)
	userRepo := userRepo.New()
	userUseCase := user.NewUseCase(userRepo)

	beersvc.NewController(engine, beerUseCase)
	usersvc.NewController(engine, userUseCase)
}

func New() Server {
	return &server{}
}
