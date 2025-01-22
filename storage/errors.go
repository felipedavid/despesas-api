package storage

import "errors"

var (
	ErrDuplicatedEmail = errors.New(`duplicated email`)
	ErrNoRows          = errors.New(`no rows`)
)
