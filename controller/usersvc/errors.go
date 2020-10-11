package usersvc

import "fmt"

type InvalidArgsError struct {
	Message string
}

func (e InvalidArgsError) Error() string {
	return fmt.Sprintf("invaid argument error : %s", e.Message)
}
