package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var createAccountParams service.CreateAccountParams
	err := readJSON(r, &createAccountParams)
	if err != nil {
		return BadRequestError(err.Error())
	}

	if !createAccountParams.Valid() {
		return ValidationError(createAccountParams.Errors)
	}

	newAccount := createAccountParams.Model(1)
	err = storage.InsertAccount(context.Background(), newAccount)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, newAccount)
}

func handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	accountID, err := strconv.Atoi(r.PathValue("accountID"))
	if err != nil {
		return BadRequestError(err.Error())
	}

	err = storage.DeleteAccount(context.Background(), 1, accountID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func handleListUserAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := storage.ListUserAccounts(context.Background(), 1)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, accounts)
}
