package storage

import (
	"context"

	"github.com/felipedavid/saldop/models"
)

func InsertTransaction(ctx context.Context, t *models.Transaction) error {
	query := `
		INSERT INTO transaction (
		 	external_id,
			user_id,
			account_id,
			description,
			amount,
			currency_code,
			transaction_date,
			category_id,
			status,
			type,
			operation_type
		) VALUES (
		 	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		) RETURNING id, created_at, updated_at
	`

	err := conn.QueryRow(
		ctx, query,
		t.ExternalID,
		t.UserID,
		t.AccountID,
		t.Description,
		t.Amount,
		t.CurrencyCode,
		t.TransactionDate,
		t.CategoryID,
		t.Status,
		t.Type,
		t.OperationType,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func ListUserTransactions(ctx context.Context, userID int) ([]models.Transaction, error) {
	query := `
		SELECT
			id,
			external_id,
			user_id,
			account_id,
			description,
			amount,
			currency_code,
			transaction_date,
			category_id,
			status,
			type,
			operation_type,
			created_at,
			updated_at
		FROM transaction
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			&t.ID,
			&t.ExternalID,
			&t.UserID,
			&t.AccountID,
			&t.Description,
			&t.Amount,
			&t.CurrencyCode,
			&t.TransactionDate,
			&t.CategoryID,
			&t.Status,
			&t.Type,
			&t.OperationType,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func DeleteTransaction(ctx context.Context, userID, transactionID int) error {
	query := `
		UPDATE transaction
		SET deleted_at = NOW()
		WHERE user_id = $1 AND id = $2
	`

	_, err := conn.Exec(ctx, query, userID, transactionID)
	if err != nil {
		return err
	}

	return nil
}
