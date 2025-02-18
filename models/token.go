package models

import (
	"time"
)

type TokenScope string

const (
	TokenScopeAuthentication TokenScope = "authentication"
)

type Token struct {
	Plaintext string     `json:"value"`
	Hash      []byte     `json:"-"`
	UserID    string     `json:"-"`
	Expiry    time.Time  `json:"expiry"`
	Scope     TokenScope `json:"-"`
}
