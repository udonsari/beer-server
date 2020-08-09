package controller

import (
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/labstack/echo"
)

type CustomContext interface {
	echo.Context
	User() (*user.User, error)
	UserMust() (*user.User, error)
}
