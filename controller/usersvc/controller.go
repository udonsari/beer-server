package usersvc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/usersvc/dto"
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

const kakaoOauthURL = "https://kauth.kakao.com/oauth/authorize"
const kakaoTokenURL = "https://kauth.kakao.com/oauth/token"
const kakaoUserURL = "https://kapi.kakao.com/v2/user/me"

// TODO Remove
const appKey = "99db370b453955b495832f7f69294dcc"

type Controller struct {
	controller.Base
	userUseCase user.UseCase
	mapper      dto.Mapper
}

func NewController(engine *echo.Echo, userUseCase user.UseCase) Controller {
	controller := Controller{
		userUseCase: userUseCase,
		mapper:      dto.NewMapper(),
	}
	// TODO REST API 컨벤션 처리
	engine.GET("/api/kakao/signin", controller.SignInKakao)
	engine.GET("/api/kakao/token", controller.GetKakaoToken)
	engine.GET("/api/kakao/user", controller.GetKakaoUser)
	return controller
}

func (cont *Controller) SignInKakao(ctx echo.Context) error {
	log.Printf("Controller - SignInKakao() - Controller")

	redirectBaseURL := controller.HOST + "/api/kakao/token"
	redirectURL := kakaoOauthURL + "?client_id=" + appKey + "&redirect_uri=" + redirectBaseURL + "&response_type=code"

	log.Printf("Controller - SignInKakao() - redirectURL : %v", redirectURL)
	err := ctx.Redirect(http.StatusPermanentRedirect, redirectURL)
	if err != nil {
		log.Printf("Controller - SignInKakao() - Redirect Error %+v", err)
	}
	return err
}

func (cont *Controller) GetKakaoToken(ctx echo.Context) error {
	// TODO Check token expiration

	log.Printf("Controller - GetKakaoToken() - Controller Param %+v", spew.Sdump(ctx.Request()))

	code := ctx.QueryParam("code")

	redirectBaseURL := controller.HOST + "/api/kakao/token"

	url := kakaoTokenURL + "?grant_type=authorization_code" + "&client_id=" + appKey + "&redirect_uri=" + redirectBaseURL + "&code=" + code
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Controller - GetKakaoToken() - Kakao token error %v", err)
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Controller - GetKakaoToken() - Kakao token read body error %v", err)
		return err
	}
	var kakaoToken dto.KakaoToken
	err = json.Unmarshal(respBytes, &kakaoToken)
	if err != nil {
		return err
	}

	log.Printf("Controller - GetKakaoToken() - respBytes : %+v", string(respBytes))
	log.Printf("Controller - GetKakaoToken() - Token : %+v", kakaoToken)
	resp.Body.Close()

	// Access Token 내려주고, 임의의 API는 모두 Access Token 필요하고, 받고 나서는 Kakao API 호출해서 external id 얻고 DB에서 대응하는 유저 찾아서 처리
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, kakaoToken)
}

func (cont *Controller) GetKakaoUser(ctx echo.Context) error {
	var req dto.GetKakaoUserRequest
	if err := cont.Bind(ctx, &req); err != nil {
		log.Printf("Controller - GetBeers() - Failed to bind %+v", err)
		return err
	}
	kakaoUser, err := cont.getKakaoUser(req.AccessToken)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, kakaoUser)
}

func (cont *Controller) getKakaoUser(accessToken string) (*dto.KakaoUser, error) {
	log.Printf("Controller - GetKakaoUser() - Controller Access Token %v", accessToken)

	req, err := http.NewRequest("GET", kakaoUserURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Controller - GetKakaoUser() - Kakao user error %v", err)
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Controller - GetKakaoUser() - Kakao user read body error %v", err)
		return nil, err
	}
	var kakaoUser dto.KakaoUser
	err = json.Unmarshal(respBytes, &kakaoUser)
	if err != nil {
		log.Printf("Controller - GetKakaoUser() - unmarshal error %v", err)
		return nil, err
	}

	log.Printf("Controller - GetKakaoUser() - respBytes : %+v", string(respBytes))
	log.Printf("Controller - GetKakaoUser() - kakaoUser : %+v", spew.Sdump(kakaoUser))
	resp.Body.Close()

	user, err := cont.userUseCase.GetUserByExternalID(strconv.FormatInt(kakaoUser.ID, 10))
	if err != nil {
		return nil, err
	}

	// TODO Go routine 처리 고민
	if user == nil {
		err = cont.userUseCase.CreateUser(
			cont.mapper.MapDTOKakaoUserToUser(kakaoUser),
		)
		user, err = cont.userUseCase.GetUserByExternalID(strconv.FormatInt(kakaoUser.ID, 10))
		if err != nil {
			return nil, err
		}
	}

	return &kakaoUser, nil
}
