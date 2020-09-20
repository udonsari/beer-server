package user

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Pallinder/go-randomdata"
)

type mapper struct {
}

func NewMapper() mapper {
	return mapper{}
}

func (m *mapper) MapKakaoUserToUser(kakaoUser KakaoUser) User {
	return User{
		ExternalID: strconv.FormatInt(kakaoUser.ID, 10),
		Properties: Properties{
			NickName: m.getRandomNickName(),
			// NickName:       kakaoUser.Properties.NickName, // We use our own nickname not kakao
			ProfileImage:   kakaoUser.Properties.ProfileImage,
			ThumbnailImage: kakaoUser.Properties.ThumbnailImage,
		},
	}
}

func (m *mapper) getRandomNickName() string {
	return fmt.Sprintf("%s-%s", randomdata.SillyName(), strconv.FormatInt(rand.Int63(), 10))
}
