package controller

import (
	"github.com/labstack/echo"
)

const PORT = 8081
const STR_PORT = "8081"
const HOST = "http://127.0.0.1:" + STR_PORT

type Base struct {
}

func (con *Base) Bind(c echo.Context, model interface{}) error {
	if err := c.Bind(model); err != nil {
		return NewBindError(err.Error())
	}
	return nil
}
