package repo

import (
	"log"

	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
)

type mockUserRepo struct {
	db *gorm.DB
}

func NewMockUserRepo() *mockUserRepo {
	return &mockUserRepo{}
}

func (r *mockUserRepo) GetUserByExternalID(externalID string) (*user.User, error) {
	log.Printf("MockUserRepo - GetUserByExternalID() - externalID %+v", externalID)

	return &user.User{
		ID:         1,
		ExternalID: externalID,
		Properties: user.Properties{
			NickName:       "TEST_NICKNAME",
			ProfileImage:   "TEST_PROFILE_IMAGE",
			ThumbnailImage: "TEST_THUMBNAIL_IMAGE",
		},
	}, nil
}

func (r *mockUserRepo) CreateUser(user user.User) error {
	log.Printf("MockUserRepo - CreateUser() - user %+v", spew.Sdump(user))
	return nil
}
