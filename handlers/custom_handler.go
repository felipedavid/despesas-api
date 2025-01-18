package handlers

import (
	"log/slog"
	"net/http"
)

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
