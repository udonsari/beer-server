package dto

type MapperError struct {
	msg string
}

func NewMapperError(msg string) error {
	return MapperError{msg: msg}
}

func (e MapperError) Error() string {
	return e.msg
}
