package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/felipedavid/saldop/filters"
	"github.com/felipedavid/saldop/helpers"
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
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

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
	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

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
	return nil
}
