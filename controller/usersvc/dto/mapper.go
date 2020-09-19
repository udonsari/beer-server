package dto

import "github.com/UdonSari/beer-server/domain/user"

type Mapper struct {
}

func NewMapper() Mapper {
	return Mapper{}
}

func (m *Mapper) MapUserToDTOUser(user user.User) User {
	return User{
		ID:         user.ID,
		ExternalID: user.ExternalID,
		Properties: Properties{
			NickName:       user.NickName,
			ProfileImage:   user.ProfileImage,
			ThumbnailImage: user.ThumbnailImage,
		},
	}
}
