package models

import (
	"time"
)

type TokenScope string

const (
	TokenScopeAuthentication = "authentication"
)

type Token struct {
	Plaintext string     `json:"value"`
	Hash      []byte     `json:"-"`
	UserID    int        `json:"-"`
	Expiry    time.Time  `json:"expiry"`
	Scope     TokenScope `json:"-"`
}
