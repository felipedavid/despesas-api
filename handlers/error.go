package handlers

import "errors"

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (e StatusError) Error() string {
	return e.Err.Error()
}

func (e StatusError) Status() int {
	return e.Code
}

func newError(status int, errMsg string) Error {
	return &StatusError{Code: status, Err: errors.New(errMsg)}
}
