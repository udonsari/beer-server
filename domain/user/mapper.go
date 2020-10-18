package user

import (
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
	// 10자 이내로 제한을 걸어달라는 요구사항이 생겨서 붙였지만, 사용자가 동일한 닉네임에 대해 헷갈리지 않을까 UX 적인 고민 필요
	return randomdata.SillyName()
	// return fmt.Sprintf("%s-%s", randomdata.SillyName(), strconv.FormatInt(rand.Int63(), 10))
}
