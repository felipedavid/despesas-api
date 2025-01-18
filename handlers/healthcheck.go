package handlers

import (
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) error {
	return newError(http.StatusInternalServerError, "something ent wrong")
}
