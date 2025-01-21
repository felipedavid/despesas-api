package service

import (
	"context"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/validator"
)

type CreateAccountParams struct {
	Name         *string `json:"name"`
	Type         *string `json:"type"`
	Balance      *int    `json:"balance"`
	CurrencyCode *string `json:"currency_code"`

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

	return len(p.Errors) == 0
}

func (p *CreateAccountParams) Model(userID int) *models.Account {
	return &models.Account{
		UserID:       userID,
		Name:         *p.Name,
		Type:         *p.Type,
		Balance:      *p.Balance,
		CurrencyCode: *p.CurrencyCode,
	}
}
