package service

import (
	"context"
	"time"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
)

type CreateTransactionParams struct {
	Description     *string    `json:"description"`
	AccountID       *int       `json:"account_id"`
	Amount          *int       `json:"amount"`
	CurrencyCode    *string    `json:"currency_code"`
	TransactionDate *time.Time `json:"transaction_date"`
	CategoryID      *int       `json:"category_id"`

	*Validator
}

func NewCreateTransactionParams(ctx context.Context) *CreateTransactionParams {
	return &CreateTransactionParams{
		Validator: NewValidator(helpers.GetTranslator(ctx)),
	}
}

func (p *CreateTransactionParams) Valid() bool {
	p.Check(p.Amount != nil, "amount", "must be provided")
	p.Check(p.CurrencyCode != nil, "currency_code", "must be provided")
	p.Check(p.TransactionDate != nil, "transaction_date", "must be provided")

	return len(p.Errors) == 0
}

func (p *CreateTransactionParams) Model(userID int) *models.Transaction {
	return &models.Transaction{
		UserID:          userID,
		AccountID:       p.AccountID,
		CategoryID:      p.CategoryID,
		Description:     p.Description,
		Amount:          *p.Amount,
		CurrencyCode:    *p.CurrencyCode,
		TransactionDate: *p.TransactionDate,
	}
}
