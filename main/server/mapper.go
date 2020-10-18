package server

import (
	"net/http"

	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/beersvc"
	"github.com/UdonSari/beer-server/controller/usersvc"
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/pkg/errors"
)

type errorMapper struct{}

func (m errorMapper) fromDomainErrorToStatusCode(err error) int {
	var code int
	switch errors.Cause(err).(type) {
	case
		usersvc.InvalidArgsError,
		beersvc.InvalidArgsError,
		controller.BindError:

		code = http.StatusBadRequest

	case
		user.InvalidTokenError:

		code = http.StatusUnauthorized

	case
		user.UserNotFoundError:

		code = http.StatusNotFound

	case
		user.ProviderError:

		code = http.StatusInternalServerError

	default:
		code = http.StatusInternalServerError
	}

	return code
}
