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
	p.Check(p.Email != nil && len(*p.Email) != 0, "email", "cannot be empty")
	p.Check(p.Password != nil && len(*p.Password) >= 8, "password", "should be at least 8 characters long")

	return len(p.Errors) == 0
}

type AuthenticationResponse struct {
	User  *models.User  `json:"user"`
	Token *models.Token `json:"token"`
}

func CredentialsAuthentication(params *CredentialsAuthenticationParams) (*AuthenticationResponse, error) {
	if !params.Valid() {
		return nil, ErrFailedValidation
	}

	user, err := storage.FindUserByEmail(context.Background(), *params.Email)
	switch {
	case err == nil:
	case errors.Is(err, storage.ErrNoRows):
		return nil, ErrInvalidCredentials
	default:
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(*params.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := CreateToken(user.ID, 24*time.Hour, models.TokenScopeAuthentication)
	if err != nil {
		return nil, err
	}

	return &AuthenticationResponse{
		User:  user,
		Token: token,
	}, nil
}
