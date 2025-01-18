package handlers

import (
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) error {
	res := map[string]any{
		"status": "healthy",
	}

	return writeJSON(w, http.StatusOK, res)
}
