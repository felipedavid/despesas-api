package service

import (
	"context"

	"github.com/felipedavid/saldop/localizer"
	"github.com/felipedavid/saldop/models"
)

type CreateAccountParams struct {
	Name         *string `json:"name"`
	Type         *string `json:"type"`
	Balance      *int    `json:"balance"`
	CurrencyCode *string `json:"currency_code"`

	Validator
}

func NewCreateAccountParams(ctx context.Context) *CreateAccountParams {
	var params CreateAccountParams
	params.Validator.Localizer, _ = localizer.Get(ctx.Value("language").(string))

	return &params
}

func (p *CreateAccountParams) Validate(ctx context.Context) bool {
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
