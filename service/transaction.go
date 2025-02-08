package service

import (
	"context"
	"errors"
	"time"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/nullable"
	"github.com/felipedavid/saldop/storage"
	"github.com/felipedavid/saldop/validator"
)

type CreateTransactionParams struct {
	Description     *string    `json:"description"`
	AccountID       *int       `json:"account_id"`
	CategoryID      *int       `json:"category_id"`
	Amount          *int       `json:"amount"`
	CurrencyCode    *string    `json:"currency_code"`
	TransactionDate *time.Time `json:"transaction_date"`
	UserID          *int

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
	Description     nullable.Nullable[*string]   `json:"description"`
	AccountID       nullable.Nullable[*int]      `json:"account_id"`
	CategoryID      nullable.Nullable[*int]      `json:"category_id"`
	Amount          nullable.Nullable[int]       `json:"amount"`
	CurrencyCode    nullable.Nullable[string]    `json:"currency_code"`
	TransactionDate nullable.Nullable[time.Time] `json:"transaction_date"`

	*validator.Validator
}

func NewEditTransactionParams(ctx context.Context) *EditTransactionParams {
	return &EditTransactionParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func PatchValue[T any](v *T, attr nullable.Nullable[T]) {
	attrValue, err := attr.Get()
	if err != nil {
		switch {
		case errors.Is(err, nullable.ErrValueIsNull):
			var zero T
			*v = zero
		}
		return
	}

	*v = attrValue
}
