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

func createTransaction(w http.ResponseWriter, r *http.Request) error {
	params := service.NewCreateTransactionParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}
	params.UserID = &user.ID

	newTransaction, err := service.CreateTransaction(params)
	if err != nil {
		if errors.Is(err, service.ErrFailedValidation) {
			return ValidationError(params.Errors)
		}
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

	transactions, err := storage.ListUserTransactionsWithPopulatedFields(context.Background(), user.ID, filters)
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

	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

	err = storage.DeleteTransaction(context.Background(), user.ID, transactionID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func editTransaction(w http.ResponseWriter, r *http.Request) error {
	transactionID, err := strconv.Atoi(r.PathValue("transactionID"))
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)
	if user == nil {
		return UnauthenticatedError(r.Context())
	}

	transaction, err := storage.GetUserTransaction(context.Background(), user.ID, transactionID)
	if err != nil {
		return err
	}

	params := service.NewEditTransactionParams(r.Context())
	err = readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	params.PatchModel(transaction)

	err = storage.UpdateTransaction(context.Background(), transaction)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, transaction)
}
