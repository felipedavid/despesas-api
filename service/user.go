package service

import (
	"context"
	"errors"
	"time"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/storage"
	"github.com/felipedavid/saldop/validator"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserParams struct {
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	PhoneNumber *string `json:"phone_number"`

	*validator.Validator
}

func NewRegisterUserParams(ctx context.Context) *RegisterUserParams {
	return &RegisterUserParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
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

var ErrUnableToHashPassword = errors.New(`unable to hash passwords`)
var ErrDuplicatedEmail = errors.New(`duplicated email`)

func RegisterUser(params *RegisterUserParams) (*AuthenticationResponse, error) {
	if !params.Valid() {
		return nil, ErrFailedValidation
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, ErrUnableToHashPassword
	}

	newUser := &models.User{
		Name:        *params.Name,
		Email:       *params.Email,
		Password:    hashedPassword,
		PhoneNumber: params.PhoneNumber,
	}
	err = storage.InsertUser(context.Background(), newUser)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicatedEmail) {
			return nil, err
		}
		return nil, err
	}

	tk, err := CreateToken(newUser.ID, 24*time.Hour*30, models.TokenScopeAuthentication)
	if err != nil {
		return nil, err
	}

	res := &AuthenticationResponse{
		User:  newUser,
		Token: tk,
	}

	return res, nil
}

type UserAuthParams struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
	*validator.Validator
}

func NewUserAuthParams(ctx context.Context) *UserAuthParams {
	return &UserAuthParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func (p *UserAuthParams) Valid() bool {
	p.Check(p.Email != nil, "email", "must be provided")
	p.Check(p.Password != nil, "password", "must be provided")

	return len(p.Errors) == 0
}
