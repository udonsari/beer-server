package server

import (
	"fmt"

	"github.com/UdonSari/beer-server/domain/user"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	UserUseCase user.UseCase
}

// TODO 여기서 ctx CustomContext를 ctx *CustomContext로 하면 Controller Test에서 controller.CustomContext 인터페이스를 만족시키지 못한다고 함. 함수 buildContextAndRecorder. 왜 그럴까 ?
func (ctx CustomContext) User() (*user.User, error) {
	accessTokens := ctx.Request().Header["Authorization"]
	if len(accessTokens) < 1 {
		return nil, nil
	}

	user, err := ctx.UserUseCase.GetUser(accessTokens[0])
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ctx CustomContext) UserMust() (*user.User, error) {
	user, err := ctx.User()
	if err != nil {
		return nil, err
	} else if user == nil || user.ID == 0 {
		return nil, fmt.Errorf("failed to find user")
	}
	return user, nil
}
