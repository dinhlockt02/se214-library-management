package coreerror

type InternalServerError struct {
	Msg string
}

func (err *InternalServerError) Error() string {
	return err.Msg
}

func NewInternalServerError(msg string) *InternalServerError {
	return &InternalServerError{
		Msg: msg,
	}
}
