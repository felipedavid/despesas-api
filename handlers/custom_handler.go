package handlers

import "net/http"

type customHandler func(w http.ResponseWriter, r *http.Request) error

func httpHandler(h customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
		}
	}
}
