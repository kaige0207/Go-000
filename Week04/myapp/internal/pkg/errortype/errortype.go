package errortype

// New returns an error that formats as the given text.
func New(code int, msg string) error {
	return &ErrUserNotFound{code, msg}
}

// errorString is a trivial implementation of error.
type ErrUserNotFound struct {
	code int
	msg  string
}

func (e *ErrUserNotFound) Error() string {
	return e.msg
}
