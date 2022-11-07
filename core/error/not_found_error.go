package coreerror

type NotFoundError struct {
	Msg string
	err error
}

func (err NotFoundError) Error() string {
	return err.Msg
}

func NewNotFoundError(msg string, err error) *NotFoundError {
	return &NotFoundError{
		Msg: msg,
		err: err,
	}
}
