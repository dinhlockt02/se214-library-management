package coreerror

type BadRequestError struct {
	Msg string
}

func (err *BadRequestError) Error() string {
	return err.Msg
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		Msg: msg,
	}
}
