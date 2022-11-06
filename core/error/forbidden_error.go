package coreerror

type ForbiddenError struct {
	Msg string
	err error
}

func (err *NotFoundError) ForbiddenError() string {
	return err.Msg
}

func NewForbiddenError(msg string, err error) *ForbiddenError {
	return &ForbiddenError{
		Msg: msg,
		err: err,
	}
}
