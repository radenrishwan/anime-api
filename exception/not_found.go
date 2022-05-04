package exception

type NotFoundError struct {
	Status int
	Err    error
}

func NewNotFoundError(status int, err error) NotFoundError {
	return NotFoundError{Status: status, Err: err}
}

func (error NotFoundError) Error() string {
	return error.Err.Error()
}
