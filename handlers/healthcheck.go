package handlers

import (
	"fmt"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "healthy")
	return nil
}
