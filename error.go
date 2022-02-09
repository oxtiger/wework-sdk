package wework

import "fmt"

type weworkError struct {
	err string
}

func (e *weworkError) Error() string {
	return e.err
}

func NewWeworkError(format string, a ...interface{}) weworkError {
	return weworkError{err: fmt.Sprintf(format, a...)}
}

func IsWeworkError(err interface{}) bool {
	_, b := err.(*weworkError)
	return b
}
