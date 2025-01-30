package service

import (
	"context"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/validator"
)

type CreateCategoryParams struct {
	Name *string `json:"name"`

	*validator.Validator
}

func NewCreateCategoryParams(ctx context.Context) *CreateCategoryParams {
	return &CreateCategoryParams{
		Validator: validator.New(helpers.GetTranslator(ctx)),
	}
}

func (p *CreateCategoryParams) Validate() bool {
	p.Check(p.Name != nil, "name", "must be provided")

	return len(p.Errors) == 0
}

func (p *CreateCategoryParams) Model(userID *int) *models.Category {
	return &models.Category{
		Name:   *p.Name,
		UserID: userID,
	}
}
