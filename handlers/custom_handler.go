package handlers

import (
	"errors"
	"log/slog"
	"net/http"
)

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

type customHandler func(w http.ResponseWriter, r *http.Request) error

func handleErrors(h customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			switch e := err.(type) {
			case Error:
				slog.Error(e.Error())
				http.Error(w, e.Error(), e.Status())
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	}
}
