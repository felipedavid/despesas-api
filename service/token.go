package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"

	"github.com/felipedavid/saldop/internal/validator"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/storage"
)

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "should be 26 bytes long")
}

func generateToken(userID string, ttl time.Duration, scope models.TokenScope) (*models.Token, error) {
	token := &models.Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}

func CreateToken(userID string, ttl time.Duration, scope models.TokenScope) (*models.Token, error) {
	t, err := generateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = storage.InsertToken(context.Background(), t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
