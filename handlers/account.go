package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func createAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountParams := service.NewCreateAccountParams(r.Context())
	err := readJSON(r, &createAccountParams)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	if !createAccountParams.Validate() {
		return ValidationError(createAccountParams.Errors)
	}

	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

	newAccount := createAccountParams.Model(user.ID)
	err = storage.InsertAccount(context.Background(), newAccount)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, newAccount)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) error {
	accountID, err := strconv.Atoi(r.PathValue("accountID"))
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

	err = storage.DeleteAccount(context.Background(), user.ID, accountID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func listUserAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := storage.ListUserAccounts(context.Background(), 1)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, accounts)
}
