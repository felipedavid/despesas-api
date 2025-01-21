package storage

import (
	"context"

	"github.com/felipedavid/saldop/models"
)

func InsertToken(ctx context.Context, token *models.Token) error {
	query := `INSERT INTO token (hash, user_id, expiry, scope) VALUES ($1, $2, $3, $4)`

	_, err := conn.Exec(ctx, query, token.Hash, token.UserID, token.Expiry, token.Scope)
	return err
}

func DeleteAllUserTokens(ctx context.Context, userID int, scope models.TokenScope) error {
	query := `DELETE FROM token WHERE user_id = $1`

	_, err := conn.Exec(ctx, query, userID, scope)
	return err
}
