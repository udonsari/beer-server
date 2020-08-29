package user

import "strconv"

type mapper struct {
}

func NewMapper() mapper {
	return mapper{}
}

func (m *mapper) MapKakaoUserToUser(kakaoUser KakaoUser) User {
	return User{
		ExternalID: strconv.FormatInt(kakaoUser.ID, 10),
		Properties: Properties{
			NickName:       kakaoUser.Properties.NickName,
			ProfileImage:   kakaoUser.Properties.ProfileImage,
			ThumbnailImage: kakaoUser.Properties.ThumbnailImage,
		},
	}
}
