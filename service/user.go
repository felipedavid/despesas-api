package service

import (
	"github.com/felipedavid/saldop/models"
)

type RegisterUserParams struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber *string `json:"phone_number"`

	Validator
}

func (p *RegisterUserParams) Valid() bool {
	p.Check(p.Name != nil, "name", "cannot be empty")
	p.Check(p.Email != nil, "email", "cannot be empty")
	p.Check(p.Password != nil, "password", "cannot be empty")

	if p.Name != nil {
		p.Check(len(*p.Name) >= 3, "name", "should be longer then 8 characters")
	}

	if p.Email != nil {
		// TODO: Add email correctness check
	}

	if p.Password != nil {
		p.Check(len(*p.Password) >= 8, "password", "cannot be empty")
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
