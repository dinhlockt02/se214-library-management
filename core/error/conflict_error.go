package coreerror

type ConflictError struct {
	Msg string
	err error
}

func (err ConflictError) Error() string {
	return err.Msg
}

func NewConflictError(msg string, err error) *ConflictError {
	return &ConflictError{
		Msg: msg,
		err: err,
	}
}
