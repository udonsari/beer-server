package usersvc

import (
	"log"
	"net/http"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/usersvc/dto"
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

type Controller struct {
	controller.Base
	userUseCase user.UseCase
}

func NewController(engine *echo.Echo, userUseCase user.UseCase) Controller {
	cont := Controller{
		userUseCase: userUseCase,
	}
	// TODO REST API 컨벤션 처리
	engine.GET("/api/kakao/signin", cont.SignInKakao)
	engine.GET("/api/token", cont.GetToken)
	engine.GET("/api/user", cont.GetUser)
	return cont
}

func (cont *Controller) SignInKakao(ctx echo.Context) error {
	log.Printf("Controller - SignInKakao() - Controller")

	redirectBaseURL := user.HOST + "/api/kakao/token"
	redirectURL := user.KakaoOauthURL + "?client_id=" + user.KakaoAppKey + "&redirect_uri=" + redirectBaseURL + "&response_type=code"

	log.Printf("Controller - SignInKakao() - redirectURL : %v", redirectURL)
	err := ctx.Redirect(http.StatusPermanentRedirect, redirectURL)
	if err != nil {
		log.Printf("Controller - SignInKakao() - Redirect Error %+v", err)
	}
	return err
}

func (cont *Controller) GetToken(ctx echo.Context) error {
	// TODO Check Token Expiration
	log.Printf("Controller - GetToken() - Controller Param %+v", spew.Sdump(ctx.Request()))

	code := ctx.QueryParam("code")
	token, err := cont.userUseCase.GetToken(code)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, token)
}

func (cont *Controller) GetUser(ctx echo.Context) error {
	// TODO Token이 Header에서 오게 수정
	var req dto.GetUserRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeers() - Failed to bind %+v", err)
		return err
	}
	user, err := cont.userUseCase.GetUser(req.AccessToken)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}