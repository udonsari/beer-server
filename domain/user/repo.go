package user

type UserRepo interface {
	GetUserByExternalID(externalID string) (*User, error)
	CreateUser(user User) error
}
