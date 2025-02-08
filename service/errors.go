package service

import "errors"

var ErrFailedValidation = errors.New(`failed validation`)
var ErrInvalidCredentials = errors.New(`invalid credentials`)
