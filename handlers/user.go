package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/service"
)

func registerUser(w http.ResponseWriter, r *http.Request) error {
	params := service.NewRegisterUserParams(r.Context())
	if err := readJSON(r, params); err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	res, err := service.RegisterUser(params)
	if err != nil {
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
