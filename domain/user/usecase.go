package user

type UseCase interface {
	GetUserByExternalID(externalID string) (*User, error)
	CreateUser(user User) error
}

type useCase struct {
	userRepo UserRepo
}

func NewUseCase(userRepo UserRepo) UseCase {
	return &useCase{
		userRepo: userRepo,
	}
}

func (u *useCase) GetUserByExternalID(externalID string) (*User, error) {
	return u.userRepo.GetUserByExternalID(externalID)
}

func (u *useCase) CreateUser(user User) error {
	return u.userRepo.CreateUser(user)
}
