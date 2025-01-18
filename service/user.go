package service

import "github.com/felipedavid/saldop/models"

type RegisterUserParams struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber *string `json:"phone_number"`
}

func (p RegisterUserParams) Validate() error {
	return nil
}

func (p RegisterUserParams) Valid() bool {
	return true
}

func (p RegisterUserParams) Model() *models.User {
	return &models.User{
		Name:        *p.Name,
		Email:       *p.Email,
		Password:    *p.Password,
		PhoneNumber: p.PhoneNumber,
	}
}
