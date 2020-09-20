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

// TODO Kakao repo를 파서 통신 내리기

const (
	createRetryCount = 10
)

type UseCase interface {
	CreateUser(user User) error
	GetToken(code string) (*Token, error)
	GetUser(accessToken string) (*User, error)
	GetUserByID(userID int64) (*User, error)
	GetUserByExternalID(externalID string) (*User, error)
	UpdateNickName(userID int64, nickName string) error
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

func (u *useCase) CreateUser(user User) error {
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
		// [연동 참조] 아래 주석을 풀어줍니다
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
		// Unique 조건에서 걸릴수 있으니 Retry
		for i := 0; i < createRetryCount; i++ {
			err = u.CreateUser(
				u.mapper.MapKakaoUserToUser(kakaoUser),
			)
			if err != nil {
				continue
			} else {
				break
			}
		}
		if err != nil {
			return nil, fmt.Errorf("failed to create user with retry %+v", err)
		}
		user, err = u.GetUserByExternalID(strconv.FormatInt(kakaoUser.ID, 10))
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (u *useCase) GetUserByID(userID int64) (*User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *useCase) GetUserByExternalID(externalID string) (*User, error) {
	return u.userRepo.GetUserByExternalID(externalID)
}

func (u *useCase) UpdateNickName(userID int64, nickName string) error {
	return u.userRepo.UpdateNickName(userID, nickName)
}
