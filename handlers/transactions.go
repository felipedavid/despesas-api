package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/felipedavid/saldop/filters"
	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func createTransaction(w http.ResponseWriter, r *http.Request) error {
	params := service.NewCreateTransactionParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	if !params.Valid() {
		return ValidationError(params.Errors)
	}

	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

	newTransaction := params.Model(user.ID)
	err = storage.InsertTransaction(context.Background(), newTransaction)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, newTransaction)
}

func listUserTransactions(w http.ResponseWriter, r *http.Request) error {
	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

	filters := filters.NewQueryFilters(r)
	if !filters.Valid() {
		return QueryValidationError(filters.Errors)
	}

	transactions, err := storage.ListUserTransactions(context.Background(), user.ID, filters)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, map[string]any{
		"metadata":     filters.Metadata(),
		"transactions": transactions,
	})
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) error {
	transactionID, err := strconv.Atoi(r.PathValue("transactionID"))
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	err = storage.DeleteTransaction(context.Background(), 1, transactionID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}
