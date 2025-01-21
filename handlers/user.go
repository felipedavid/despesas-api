package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func registerUser(w http.ResponseWriter, r *http.Request) error {
	params := service.NewRegisterUserParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(err.Error())
	}

	if !params.Valid() {
		return ValidationError(params.Errors)
	}

	newUser := params.Model()
	err = storage.InsertUser(context.Background(), newUser)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicatedEmail) {
			return ErrorRes(http.StatusConflict, err.Error(), nil)
		}
		return err
	}

	return writeJSON(w, http.StatusCreated, newUser)
}

func authenticateUser(w http.ResponseWriter, r *http.Request) error {
	params := service.NewUserAuthParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(err.Error())
	}

	if !params.Valid() {
		return ValidationError(params.Errors)
	}

	token := map[string]any{}
	return writeJSON(w, http.StatusCreated, token)
}
