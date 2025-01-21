package service

import (
	"context"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/validator"
)

type CredentialsAuthenticationParams struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	*validator.Validator
}

func NewCredentialsAuthenticationParams(ctx context.Context) *CredentialsAuthenticationParams {
	return &CredentialsAuthenticationParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func (p *CredentialsAuthenticationParams) Valid() bool {
	p.Check(p.Email != nil, "email", "must be provided")
	p.Check(p.Password != nil, "password", "must be provided")

	if p.Email != nil {
		p.Check(len(*p.Email) != 0, "email", "cannot be empty")
	}

	if p.Password != nil {
		p.Check(len(*p.Password) >= 8, "password", "should be at least 8 characters long")
	}

	return len(p.Errors) == 0
}
