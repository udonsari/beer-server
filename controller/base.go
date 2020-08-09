package controller

import (
	"github.com/labstack/echo"
)

type Base struct {
}

func (con *Base) Bind(c echo.Context, model interface{}) error {
	if err := c.Bind(model); err != nil {
		return NewBindError(err.Error())
	}
	return nil
}
