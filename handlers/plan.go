package handlers

import (
	"fmt"
	"net/http"
)

func getActivePlan(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "not implemented yet")
	return nil
}

func listAvailablePlans(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "not implemented yet")
	return nil
}
