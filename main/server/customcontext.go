package server

import (
	"fmt"

	"github.com/UdonSari/beer-server/domain/user"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
	userUseCase user.UseCase
}

func (ctx *CustomContext) User() (*user.User, error) {
	accessTokens := ctx.Request().Header["Authorization"]
	if len(accessTokens) < 1 {
		return nil, nil
	}

	user, err := ctx.userUseCase.GetUser(accessTokens[0])
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ctx *CustomContext) UserMust() (*user.User, error) {
	user, err := ctx.User()
	if err != nil {
		return nil, err
	} else if user == nil || user.ID == 0 {
		return nil, fmt.Errorf("failed to find user")
	}
	return user, nil
}
