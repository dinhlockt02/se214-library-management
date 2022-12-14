package coreerror

type InternalServerError struct {
	Msg string
	err error
}

func (err InternalServerError) Error() string {
	return err.Msg
}

func (err InternalServerError) UnderlyingError() error {
	return err.err
}

func NewInternalServerError(msg string, err error) *InternalServerError {
	return &InternalServerError{
		Msg: msg,
		err: err,
	}
}
