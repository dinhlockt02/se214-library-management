package coreerror

type BusinessError struct {
	Msg string
}

func (err *BusinessError) Error() string {
	return err.Msg
}

func NewBusinessError(msg string) *BusinessError {
	return &BusinessError{
		Msg: msg,
	}
}
