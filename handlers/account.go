package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/felipedavid/saldop/internal/filters"
	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func createAccount(w http.ResponseWriter, r *http.Request) error {
	params := service.NewCreateAccountParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)

	params.UserID = &user.ID

	acc, err := service.CreateManualAccount(params)
	if err != nil {
		if errors.Is(err, service.ErrFailedValidation) {
			return ValidationError(params.Errors)
		}
		return err
	}

	return writeJSON(w, http.StatusCreated, acc)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) error {
	accountID := r.PathValue("accountID")

	user := helpers.GetUserFromRequestContext(r)

	err := storage.DeleteAccount(context.Background(), user.ID, accountID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func getUserAccount(w http.ResponseWriter, r *http.Request) error {
	user := helpers.GetUserFromRequestContext(r)

	accountID := r.PathValue("accountID")

	account, err := storage.GetUserAccount(context.Background(), user.ID, accountID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, account)
}

func listUserAccounts(w http.ResponseWriter, r *http.Request) error {
	user := helpers.GetUserFromRequestContext(r)

	filters := filters.NewQueryFilters(r)
	if !filters.Valid() {
		return QueryValidationError(filters.Errors)
	}

	accounts, err := storage.ListUserAccounts(context.Background(), user.ID, filters)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"metadata": filters.Metadata(),
		"accounts": accounts,
	})
}

func editAccount(w http.ResponseWriter, r *http.Request) error {
	accountID := r.PathValue("accountID")

	params := service.NewEditAccountParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)

	account, err := storage.GetUserAccount(context.Background(), user.ID, accountID)
	if err != nil {
		return err
	}

	service.PatchValue(&account.Name, params.Name)
	service.PatchValue(&account.Type, params.Type)
	service.PatchValue(&account.Balance, params.Balance)
	service.PatchValue(&account.CurrencyCode, params.CurrencyCode)

	err = storage.UpdateAccount(context.Background(), account)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, account)
}
