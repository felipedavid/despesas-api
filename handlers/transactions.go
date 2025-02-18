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

func createTransaction(w http.ResponseWriter, r *http.Request) error {
	params := service.NewCreateTransactionParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)
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

func getUserTransaction(w http.ResponseWriter, r *http.Request) error {
	user := helpers.GetUserFromRequestContext(r)

	transactionID := r.PathValue("transactionID")

	transaction, err := storage.GetUserTransactionWithPopulatedFields(context.Background(), user.ID, transactionID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, transaction)
}

func listUserTransactions(w http.ResponseWriter, r *http.Request) error {
	user := helpers.GetUserFromRequestContext(r)

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
	transactionID := r.PathValue("transactionID")

	user := helpers.GetUserFromRequestContext(r)

	err := storage.DeleteTransaction(context.Background(), user.ID, transactionID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func editTransaction(w http.ResponseWriter, r *http.Request) error {
	transactionID := r.PathValue("transactionID")

	params := service.NewEditTransactionParams(r.Context())
	err := readJSON(r, &params)
	if err != nil {
		return BadRequestError(r.Context(), err.Error())
	}

	user := helpers.GetUserFromRequestContext(r)

	transaction, err := storage.GetUserTransaction(context.Background(), user.ID, transactionID)
	if err != nil {
		return err
	}

	service.PatchValue(&transaction.Description, params.Description)
	service.PatchValue(&transaction.AccountID, params.AccountID)
	service.PatchValue(&transaction.CategoryID, params.CategoryID)
	service.PatchValue(&transaction.Amount, params.Amount)
	service.PatchValue(&transaction.CurrencyCode, params.CurrencyCode)
	service.PatchValue(&transaction.TransactionDate, params.TransactionDate)

	err = storage.UpdateTransaction(context.Background(), transaction)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, transaction)
}
