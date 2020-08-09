package repo

import (
	"log"

	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
)

// TODO *** Attach Real DB and use ORM.
// TODO *** Add Cache ?
type userRepo struct {
}

func New() *userRepo {
	return &userRepo{}
}

func (r *userRepo) GetUserByExternalID(externalID string) (*user.User, error) {
	log.Printf("UserRepo - GetUserByExternalID() - externalID %+v", externalID)

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

func (r *userRepo) CreateUser(user user.User) error {
	log.Printf("UserRepo - GetUserByExternalID() - user %+v", spew.Sdump(user))
	return nil
}
