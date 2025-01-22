package handlers

import (
	"context"
	"net/http"
	"strconv"

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

	newAccount := createAccountParams.Model(1)
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

	err = storage.DeleteAccount(context.Background(), 1, accountID)
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
