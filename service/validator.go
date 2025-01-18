package service

type Validator struct {
	Errors map[string]string
}

func (v *Validator) Check(valid bool, attr string, errMsg string) {
	if v.Errors == nil {
		v.Errors = make(map[string]string)
	}

	if !valid {
		v.Errors[attr] = errMsg
	}
}
