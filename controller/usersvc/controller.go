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
	mapper      dto.Mapper
	host        string
}

func NewController(engine *echo.Echo, userUseCase user.UseCase, host string) Controller {
	cont := Controller{
		userUseCase: userUseCase,
		host:        host,
	}
	engine.GET("/api/kakao/signin", cont.SignInKakao)
	engine.GET("/api/token", cont.GetToken)
	engine.GET("/api/user", cont.GetUser)
	engine.POST("/api/user/update", cont.UpdateNickName)
	return cont
}

func (cont *Controller) SignInKakao(ctx echo.Context) error {
	log.Printf("Controller - SignInKakao() - Controller")

	redirectBaseURL := cont.host + "/api/token"
	redirectURL := user.KakaoOauthURL + "?client_id=" + user.KakaoAppKey + "&redirect_uri=" + redirectBaseURL + "&response_type=code"

	log.Printf("Controller - SignInKakao() - redirectURL : %v", redirectURL)
	err := ctx.Redirect(http.StatusPermanentRedirect, redirectURL)
	if err != nil {
		log.Printf("Controller - SignInKakao() - Redirect Error %+v", err)
	}
	return err
}

func (cont *Controller) GetToken(ctx echo.Context) error {
	// TODO * Check Token Expiration
	log.Printf("Controller - GetToken() - Controller Param %+v", spew.Sdump(ctx.Request()))

	code := ctx.QueryParam("code")
	token, err := cont.userUseCase.GetToken(code)
	if err != nil {
		return err
	}
	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": token,
		},
	)
}

func (cont *Controller) GetUser(ctx echo.Context) error {
	log.Printf("Controller - GetUser() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	}
	return ctx.JSON(
		http.StatusOK,
		map[string]interface{}{
			"result": cont.mapper.MapUserToDTOUser(*user),
		},
	)
}

func (cont *Controller) UpdateNickName(ctx echo.Context) error {
	log.Printf("Controller - UpdateNickName() - Controller")
	_ctx := ctx.(controller.CustomContext)
	user, err := _ctx.UserMust()
	if err != nil {
		return err
	}

	var req dto.UpdateNickNameRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - UpdateNickName() - Failed to bind %+v", err)
		return err
	}
	log.Printf("Controller - UpdateNickName() - Body %+v", spew.Sdump(req))

	err = cont.userUseCase.UpdateNickName(user.ID, req.NickName)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
