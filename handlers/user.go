package handlers

import (
	"errors"
	"net/http"

	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func registerUser(w http.ResponseWriter, r *http.Request) error {
	params := service.NewRegisterUserParams(r.Context())
	if err := readJSON(r, params); err != nil {
		return BadRequestError(r.Context(), err.Error())
	}
	t := params.Translator

	res, err := service.RegisterUser(params)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicatedEmail) {
			return ValidationError(map[string]string{
				"email": t.Translate("email already exists"),
			})
		}
		return err
	}

	return writeJSON(w, http.StatusCreated, res)
}

func authenticateUser(w http.ResponseWriter, r *http.Request) error {
	params := service.NewUserAuthParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	if !params.Valid() {
		return ValidationError(params.Errors)
	}

	token := map[string]any{}
	return writeJSON(w, http.StatusCreated, token)
}
