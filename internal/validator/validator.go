package validator

import (
	"github.com/felipedavid/saldop/internal/translations"
)

type Validator struct {
	Errors map[string]string
	*translations.Translator
}

func New(t *translations.Translator) *Validator {
	return &Validator{
		Errors:     make(map[string]string),
		Translator: t,
	}
}

func (v *Validator) Check(valid bool, attr string, errMsg string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	if !valid {
		msg := errMsg
		if v.Translator != nil {
			msg = v.Translate(errMsg)
		}

		v.Errors[attr] = msg
	}
}

func (v *Validator) AddError(attr, error string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	msg := error
	if v.Translator != nil {
		msg = v.Translate(error)
	}

	v.Errors[attr] = msg
}
