package service

import "github.com/felipedavid/saldop/localizer"

type Validator struct {
	Localizer *localizer.Localizer
	Errors    map[string]string
}

func (v *Validator) Check(valid bool, attr string, errMsg string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	if !valid {
		v.Errors[attr] = v.Localizer.Translate(errMsg)
	}
}
