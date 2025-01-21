package service

import (
	"context"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/translations"
)

type Validator struct {
	Errors map[string]string
	*translations.Translator
}

func NewValidator(ctx context.Context) *Validator {
	var v Validator
	v.Errors = make(map[string]string)
	v.Translator = helpers.GetTranslator(ctx)

	return &v
}

func (v *Validator) Check(valid bool, attr string, errMsg string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	if !valid {
		v.Errors[attr] = v.Translate(errMsg)
	}
}
