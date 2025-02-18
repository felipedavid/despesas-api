package service

import (
	"context"

	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/internal/null"
	"github.com/felipedavid/saldop/internal/validator"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/storage"
)

type CreateAccountParams struct {
	Name         *string `json:"name"`
	Type         *string `json:"type"`
	Balance      *int    `json:"balance"`
	CurrencyCode *string `json:"currency_code"`
	UserID       *string

	*validator.Validator
}

func NewCreateAccountParams(ctx context.Context) *CreateAccountParams {
	return &CreateAccountParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func (p *CreateAccountParams) Validate() bool {
	p.Check(p.Name != nil, "name", "must be provided")
	p.Check(p.Type != nil, "type", "must be provided")
	p.Check(p.Balance != nil, "balance", "must be provided")
	p.Check(p.CurrencyCode != nil, "currency_code", "must be provided")

	if p.Type != nil {
		p.Check(models.ValidAccountType(*p.Type), "type", "can only be BANK or CREDIT")
	}

	return len(p.Errors) == 0
}

func CreateManualAccount(p *CreateAccountParams) (*models.Account, error) {
	if !p.Validate() {
		return nil, ErrFailedValidation
	}

	newAccount := &models.Account{
		UserID:       *p.UserID,
		Name:         *p.Name,
		Type:         models.AccountType(*p.Type),
		Balance:      *p.Balance,
		CurrencyCode: *p.CurrencyCode,
	}
	err := storage.InsertAccount(context.Background(), newAccount)
	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

type EditAccountParams struct {
	Name         null.Nullable[string]             `json:"name"`
	Type         null.Nullable[models.AccountType] `json:"type"`
	Balance      null.Nullable[int]                `json:"balance"`
	CurrencyCode null.Nullable[string]             `json:"currency_code"`

	*validator.Validator
}

func NewEditAccountParams(ctx context.Context) *EditAccountParams {
	return &EditAccountParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}
