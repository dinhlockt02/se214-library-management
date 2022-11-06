package coreerror

type UnauthorizedError struct {
	Msg string
	err error
}

func (err *UnauthorizedError) Error() string {
	return err.Msg
}

func NewUnauthorizedError(msg string, err error) *UnauthorizedError {
	return &UnauthorizedError{
		Msg: msg,
		err: err,
	}
}
