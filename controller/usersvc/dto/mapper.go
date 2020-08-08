package dto

import (
	"strconv"

	"github.com/UdonSari/beer-server/domain/user"
)

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) MapDTOKakaoUserToUser(kakaoUser KakaoUser) user.User {
	return user.User{
		ExternalID: strconv.FormatInt(kakaoUser.ID, 10),
		Properties: user.Properties{
			NickName:       kakaoUser.Properties.NickName,
			ProfileImage:   kakaoUser.Properties.ProfileImage,
			ThumbnailImage: kakaoUser.Properties.ThumbnailImage,
		},
	}
}
