package user

import "fmt"

type InvalidTokenError struct {
	Message string
}

func (e InvalidTokenError) Error() string {
	return fmt.Sprintf("invalid token error : %s", e.Message)
}

type UserNotFoundError struct{}

func (e UserNotFoundError) Error() string {
	return "user not found error"
}

type ProviderError struct {
	Message string
}

func (e ProviderError) Error() string {
	return fmt.Sprintf("auth provider error : %s", e.Message)
}
