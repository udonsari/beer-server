package user

type UserRepo interface {
	GetUserByExternalID(externalID string) (*User, error)
	GetUserByID(userID int64) (*User, error)
	CreateUser(user User) error
	UpdateNickName(userID int64, nickName string) error
}
