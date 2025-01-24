package handlers

import (
	"net/http"

	"github.com/felipedavid/saldop/service"
)

func registerUser(w http.ResponseWriter, r *http.Request) error {
	params := service.NewRegisterUserParams(r.Context())
	err := readJSON(r, params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	// NOTE: Maybe I can do the json parsing the "New" that creates the parameter. So I don't need
	// to call readJSON on the controller anymore

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
