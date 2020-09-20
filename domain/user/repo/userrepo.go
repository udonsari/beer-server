package repo

import (
	"fmt"
	"log"

	"github.com/UdonSari/beer-server/domain/user"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

type userRepo struct {
	db *gorm.DB
	mapper
}

func New(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserByExternalID(externalID string) (*user.User, error) {
	log.Printf("UserRepo - GetUserByExternalID() - externalID %+v", externalID)

	query := DBUser{ExternalID: externalID}
	var dbUser DBUser
	if err := r.db.Where(&query).First(&dbUser).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get user. external id: %v", externalID))
		}
	}
	user := r.mapper.mapDBUserToUser(dbUser)
	return &user, nil
}

func (r *userRepo) GetUserByID(userID int64) (*user.User, error) {
	query := DBUser{ID: userID}
	var dbUser DBUser
	if err := r.db.Where(&query).First(&dbUser).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, nil
		default:
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get user. user id: %v", userID))
		}
	}
	user := r.mapper.mapDBUserToUser(dbUser)
	return &user, nil
}

func (r *userRepo) CreateUser(user user.User) error {
	log.Printf("UserRepo - CreateUser() - user %+v", spew.Sdump(user))

	dbUser := r.mapper.mapUserToDBUser(user)
	res := r.db.Create(&dbUser)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *userRepo) UpdateNickName(userID int64, nickName string) error {
	res := r.db.Model(&DBUser{}).Where("id = ?", userID).Update("nick_name", nickName)
	return res.Error
}
