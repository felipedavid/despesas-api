package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func handleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	var createUserParams service.RegisterUserParams
	err := readJSON(r, &createUserParams)
	if err != nil {
		return BadRequestError(err.Error())
	}

	if !createUserParams.Valid() {
		return ValidationError(createUserParams.Errors)
	}

	newUser := createUserParams.Model()
	err = storage.InsertUser(context.Background(), newUser)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicatedEmail) {
			return ErrorRes(http.StatusConflict, err.Error(), nil)
		}
		return err
	}

	return writeJSON(w, http.StatusCreated, newUser)
}

func handleUserAuthentication(w http.ResponseWriter, r *http.Request) error {
	var userAuthParams service.UserAuthParams
	err := readJSON(r, &userAuthParams)
	if err != nil {
		return BadRequestError(err.Error())
	}

	if !userAuthParams.Valid() {
		return ValidationError(userAuthParams.Errors)
	}

	token := map[string]any{}
	return writeJSON(w, http.StatusCreated, token)
}
