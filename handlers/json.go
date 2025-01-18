package handlers

import (
	"encoding/json"
	"net/http"
)

func readJSON(req http.Request, to any) error {
	dec := json.NewDecoder(req.Body)
	err := dec.Decode(&to)
	return err
}

func writeJSON(w http.ResponseWriter, status int, payload any) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(payload)
	return err
}
