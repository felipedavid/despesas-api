package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/felipedavid/saldop/service"
	"github.com/felipedavid/saldop/storage"
)

func createTransaction(w http.ResponseWriter, r *http.Request) error {
	var createTransactionParams service.CreateTransactionParams
	err := readJSON(r, &createTransactionParams)
	if err != nil {
		return BadRequestError(err.Error())
	}

	if !createTransactionParams.Valid() {
		return ValidationError(createTransactionParams.Errors)
	}

	newTransaction := createTransactionParams.Model(1)
	err = storage.InsertTransaction(context.Background(), newTransaction)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, newTransaction)
}

func listUserTransactions(w http.ResponseWriter, r *http.Request) error {
	transactions, err := storage.ListUserTransactions(context.Background(), 1)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, transactions)
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) error {
	transactionID, err := strconv.Atoi(r.PathValue("transactionID"))
	if err != nil {
		return BadRequestError(err.Error())
	}

	err = storage.DeleteTransaction(context.Background(), 1, transactionID)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusNoContent, nil)
}
