package exception

type NotFoundError struct {
	Error string //"Error string" -> mengikuti interface nya
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}