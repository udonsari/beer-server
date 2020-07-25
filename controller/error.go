package controller

type BindError struct {
	msg string
}

func NewBindError(msg string) error {
	return BindError{msg: msg}
}

func (e BindError) Error() string {
	return e.msg
}
