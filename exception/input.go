package exception

type InputError struct {
	Status int
	Err    error
}

func NewInputError(status int, err error) NotFoundError {
	return NotFoundError{Status: status, Err: err}
}

func (error InputError) Error() string {
	return error.Err.Error()
}
