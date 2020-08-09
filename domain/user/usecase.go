package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

const PORT = 8081
const STR_PORT = "8081"
const HOST = "http://127.0.0.1:" + STR_PORT

type UseCase interface {
	GetUserByExternalID(externalID string) (*User, error)
	GetUser(accessToken string) (*User, error)
	GetToken(code string) (*Token, error)
}

type useCase struct {
	userRepo UserRepo
	mapper   mapper
}

func NewUseCase(userRepo UserRepo) UseCase {
	return &useCase{
		userRepo: userRepo,
	}
}

func (u *useCase) GetUserByExternalID(externalID string) (*User, error) {
	return u.userRepo.GetUserByExternalID(externalID)
}

func (u *useCase) createUser(user User) error {
	return u.userRepo.CreateUser(user)
}

func (u *useCase) GetToken(code string) (*Token, error) {
	redirectBaseURL := HOST + "/api/kakao/token"

	url := KakaoTokenURL + "?grant_type=authorization_code" + "&client_id=" + KakaoAppKey + "&redirect_uri=" + redirectBaseURL + "&code=" + code
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("UseCase - GetToken() - Kakao token error %v", err)
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("UseCase - GetToken() - Kakao token read body error %v", err)
		return nil, err
	}

	var kakaoToken Token
	err = json.Unmarshal(respBytes, &kakaoToken)
	if err != nil {
		return nil, err
	}
	log.Printf("UseCase - GetToken() - respBytes : %+v", string(respBytes))
	log.Printf("UseCase - GetToken() - Token : %+v", kakaoToken)
	resp.Body.Close()

	return &kakaoToken, nil
}

func (u *useCase) GetUser(accessToken string) (*User, error) {
	req, err := http.NewRequest("GET", KakaoUserURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", accessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("UseCase - GetUser() - Kakao user error %v", err)
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("UseCase - GetUser() - Kakao user read body error %v", err)
		return nil, err
	}
	// 여기서 respBytes로 {"msg":"this access token does not exist","code":-401}가 와도 에러가 없는데 처리 필요
	var kakaoUser KakaoUser
	err = json.Unmarshal(respBytes, &kakaoUser)
	if err != nil {
		log.Printf("UseCase - GetUser() - unmarshal error %v", err)
		return nil, err
	}

	log.Printf("UseCase - GetUser() - respBytes : %+v", string(respBytes))
	log.Printf("UseCase - GetUser() - kakaoUser : %+v", spew.Sdump(kakaoUser))
	resp.Body.Close()

	user, err := u.GetUserByExternalID(strconv.FormatInt(kakaoUser.ID, 10))
	if err != nil {
		return nil, err
	}

	// TODO Go routine 처리 고민
	if user == nil {
		err = u.createUser(
			u.mapper.MapKakaoUserToUser(kakaoUser),
		)
		user, err = u.GetUserByExternalID(strconv.FormatInt(kakaoUser.ID, 10))
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
