package models

import (
	"time"
)

type TokenScope string

const (
	Authentication = "authentication"
)

type Token struct {
	Plaintext string     `json:"token"`
	Hash      []byte     `json:"-"`
	UserID    int        `json:"-"`
	Expiry    time.Time  `json:"expiry`
	Scope     TokenScope `json:"-"`
}
