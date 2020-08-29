package repo

import "github.com/UdonSari/beer-server/domain/user"

type mapper struct{}

func (m mapper) mapDBUserToUser(dbu DBUser) user.User {
	return user.User{
		ID:         dbu.ID,
		ExternalID: dbu.ExternalID,
		Properties: user.Properties{
			NickName:       dbu.NickName,
			ProfileImage:   dbu.ProfileImage,
			ThumbnailImage: dbu.ThumbnailImage,
		},
	}
}

func (m mapper) mapUserToDBUser(u user.User) DBUser {
	return DBUser{
		ExternalID:     u.ExternalID,
		NickName:       u.NickName,
		ProfileImage:   u.ProfileImage,
		ThumbnailImage: u.ThumbnailImage,
	}
}
