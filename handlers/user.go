package handlers

import (
	"context"
	"net/http"

	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) error {
	var createUserParams service.RegisterUserParams
	err := readJSON(r, &createUserParams)
	if err != nil {
		return newError(http.StatusBadRequest, "bad request")
	}

	if !createUserParams.Valid() {
		return newError(http.StatusBadRequest, "invalid request")
	}

	newUser := createUserParams.Model()
	err = storage.InsertUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, newUser)
}
