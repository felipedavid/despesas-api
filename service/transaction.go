package service

import (
	"context"
	"errors"
	"time"

	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/internal/null"
	"github.com/felipedavid/saldop/internal/validator"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/storage"
)

type CreateTransactionParams struct {
	Description     *string    `json:"description"`
	AccountID       *int       `json:"account_id"`
	CategoryID      *int       `json:"category_id"`
	Amount          *int       `json:"amount"`
	CurrencyCode    *string    `json:"currency_code"`
	TransactionDate *time.Time `json:"transaction_date"`
	UserID          *string

	*validator.Validator
}

func NewCreateTransactionParams(ctx context.Context) *CreateTransactionParams {
	return &CreateTransactionParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func (p *CreateTransactionParams) Valid() bool {
	p.Check(p.Amount != nil, "amount", "must be provided")
	p.Check(p.CurrencyCode != nil, "currency_code", "must be provided")
	p.Check(p.TransactionDate != nil, "transaction_date", "must be provided")

	return len(p.Errors) == 0
}

func CreateTransaction(p *CreateTransactionParams) (*models.Transaction, error) {
	if !p.Valid() {
		return nil, ErrFailedValidation
	}

	newTransaction := &models.Transaction{
		UserID:          *p.UserID,
		AccountID:       p.AccountID,
		CategoryID:      p.CategoryID,
		Description:     p.Description,
		Amount:          *p.Amount,
		CurrencyCode:    *p.CurrencyCode,
		TransactionDate: *p.TransactionDate,
	}
	err := storage.InsertTransaction(context.Background(), newTransaction)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

type EditTransactionParams struct {
	Description     null.Nullable[*string]   `json:"description"`
	AccountID       null.Nullable[*int]      `json:"account_id"`
	CategoryID      null.Nullable[*int]      `json:"category_id"`
	Amount          null.Nullable[int]       `json:"amount"`
	CurrencyCode    null.Nullable[string]    `json:"currency_code"`
	TransactionDate null.Nullable[time.Time] `json:"transaction_date"`

	*validator.Validator
}

func NewEditTransactionParams(ctx context.Context) *EditTransactionParams {
	return &EditTransactionParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func PatchValue[T any](v *T, attr null.Nullable[T]) {
	attrValue, err := attr.Get()
	if err != nil {
		switch {
		case errors.Is(err, null.ErrValueIsNull):
			var zero T
			*v = zero
		}
		return
	}

	*v = attrValue
}
