package error

type BusinessError struct {
	msg string
}

func (err *BusinessError) Error() string {
	return err.msg
}

func NewBusinessError(msg string) *BusinessError {
	return &BusinessError{
		msg: msg,
	}
}
