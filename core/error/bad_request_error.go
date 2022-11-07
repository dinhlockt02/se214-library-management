package coreerror

type BadRequestError struct {
	Msg string
	err error
}

func (err BadRequestError) Error() string {
	return err.Msg
}

func NewBadRequestError(msg string, err error) *BadRequestError {
	return &BadRequestError{
		Msg: msg,
		err: err,
	}
}
