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

	s._server = &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: s._engine,
	}

	log.Printf("# server up starts at port %+v ...", PORT)

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
	userUseCase := user.NewUseCase(userRepo, HOST, PORT_STR)

	beersvc.NewController(engine, beerUseCase, userUseCase)
	usersvc.NewController(engine, userUseCase, HOST)
}

func New() Server {
	return &server{}
}
