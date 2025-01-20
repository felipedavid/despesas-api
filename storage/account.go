package storage

import (
	"context"

	"github.com/felipedavid/saldop/models"
)

func InsertAccount(ctx context.Context, account *models.Account) error {
	query := `
		INSERT INTO account (
			type,
			name,
			balance,
			currency_code,
			user_id,
			external_id,
			subtype,
			number,
			owner,
			tax_number,
			bank_account_data_id,
			credit_account_data_id,
			fi_connection_id
		) VALUES (
		 	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		) RETURNING id, created_at, updated_at
	`

	err := conn.QueryRow(
		ctx, query,
		account.Type,
		account.Name,
		account.Balance,
		account.CurrencyCode,
		account.UserID,
		account.ExternalID,
		account.Subtype,
		account.Number,
		account.Owner,
		account.TaxNumber,
		account.BankAccountDataID,
		account.CreditAccountDataID,
		account.FiConnectionID,
	).Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func ListUserAccounts(ctx context.Context, userID int) ([]models.Account, error) {
	query := `
		SELECT
			id,
			type,
			name,
			balance,
			currency_code,
			user_id,
			external_id,
			subtype,
			number,
			owner,
			tax_number,
			bank_account_data_id,
			credit_account_data_id,
			fi_connection_id,
			created_at,
			updated_at
		FROM account
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	rows, err := conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&account.ID,
			&account.Type,
			&account.Name,
			&account.Balance,
			&account.CurrencyCode,
			&account.UserID,
			&account.ExternalID,
			&account.Subtype,
			&account.Number,
			&account.Owner,
			&account.TaxNumber,
			&account.BankAccountDataID,
			&account.CreditAccountDataID,
			&account.FiConnectionID,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func DeleteAccount(ctx context.Context, userID, accountID int) error {
	query := `
		UPDATE account
		SET deleted_at = NOW()
		WHERE user_id = $1 AND id = $2
	`

	_, err := conn.Exec(ctx, query, userID, accountID)
	return err
}
