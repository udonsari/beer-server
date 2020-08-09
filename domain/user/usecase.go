package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type UseCase interface {
	GetUserByExternalID(externalID string) (*User, error)
	GetUser(accessToken string) (*User, error)
	GetToken(code string) (*Token, error)
}

type useCase struct {
	userRepo UserRepo
	mapper   mapper
	port     string
	host     string
}

func NewUseCase(userRepo UserRepo, host string, port string) UseCase {
	return &useCase{
		userRepo: userRepo,
		host:     host,
		port:     port,
	}
}

func (u *useCase) GetUserByExternalID(externalID string) (*User, error) {
	return u.userRepo.GetUserByExternalID(externalID)
}

func (u *useCase) createUser(user User) error {
	return u.userRepo.CreateUser(user)
}

func (u *useCase) GetToken(code string) (*Token, error) {
	redirectBaseURL := u.host + "/api/token"

	url := KakaoTokenURL + "?grant_type=authorization_code" + "&client_id=" + KakaoAppKey + "&redirect_uri=" + redirectBaseURL + "&code=" + code

	log.Printf("UseCase - GetToken() - url : %+v", url)
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
	} else if kakaoToken.AccessToken == "" {
		// TODO * Make auth error definition
		return nil, errors.New("failed to get access token from kakao")
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

	var kakaoUser KakaoUser
	err = json.Unmarshal(respBytes, &kakaoUser)
	if err != nil {
		log.Printf("UseCase - GetUser() - unmarshal error %v", err)
		return nil, err
	} else if kakaoUser.ID == 0 {
		// TODO 실제 유저 연동 없이 동작할 때는 주석 처리한다
		// return nil, errors.New("failed to get user from kakao")
	}

	log.Printf("UseCase - GetUser() - respBytes : %+v", string(respBytes))
	log.Printf("UseCase - GetUser() - kakaoUser : %+v", spew.Sdump(kakaoUser))
	resp.Body.Close()

	user, err := u.GetUserByExternalID(strconv.FormatInt(kakaoUser.ID, 10))
	if err != nil {
		return nil, err
	}

	// Goroutine 처리 고민했지만 user를 반환 해야하므로 동기 처리
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
