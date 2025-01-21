package service

import (
	"context"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
)

type RegisterUserParams struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber *string `json:"phone_number"`

	*Validator
}

func NewRegisterUserParams(ctx context.Context) *RegisterUserParams {
	return &RegisterUserParams{
		Validator: NewValidator(helpers.GetTranslator(ctx)),
	}
}

func (p *RegisterUserParams) Valid() bool {
	p.Check(p.Name != nil, "name", "must be provided")
	p.Check(p.Email != nil, "email", "must be provided")
	p.Check(p.Password != nil, "password", "must be provided")

	if p.Name != nil {
		p.Check(len(*p.Name) >= 3, "name", "should be at least 3 characters long")
	}

	if p.Email != nil {
		// TODO: Add email correctness check
	}

	if p.Password != nil {
		p.Check(len(*p.Password) >= 8, "password", "should be at least 8 characters long")
	}

	return len(p.Errors) == 0
}

func (p *RegisterUserParams) Model() *models.User {
	return &models.User{
		Name:        *p.Name,
		Email:       *p.Email,
		Password:    *p.Password,
		PhoneNumber: p.PhoneNumber,
	}
}

type UserAuthParams struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	*Validator
}

func NewUserAuthParams(ctx context.Context) *UserAuthParams {
	return &UserAuthParams{
		Validator: NewValidator(helpers.GetTranslator(ctx)),
	}
}

func (p *UserAuthParams) Valid() bool {
	p.Check(p.Email != nil, "email", "must be provided")
	p.Check(p.Password != nil, "password", "must be provided")

	return len(p.Errors) == 0
}
