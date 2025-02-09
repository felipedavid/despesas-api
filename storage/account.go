package storage

import (
	"context"

	"github.com/felipedavid/saldop/internal/filters"
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

func GetUserAccount(ctx context.Context, userID, accountID int) (*models.Account, error) {
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
		WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL
	`

	var a models.Account
	err := conn.QueryRow(ctx, query, accountID, userID).Scan(
		&a.ID,
		&a.Type,
		&a.Name,
		&a.Balance,
		&a.CurrencyCode,
		&a.UserID,
		&a.ExternalID,
		&a.Subtype,
		&a.Number,
		&a.Owner,
		&a.TaxNumber,
		&a.BankAccountDataID,
		&a.CreditAccountDataID,
		&a.FiConnectionID,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func ListUserAccounts(ctx context.Context, userID int, f *filters.Filters) ([]models.Account, error) {
	query := `
		SELECT
            count(*) OVER(),
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
        LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(ctx, query, userID, f.Limit(), f.Offset())
	if err != nil {
		return nil, err
	}

	totalRecords := 0
	accounts := make([]models.Account, 0, f.PageSize)
	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&totalRecords,
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

	f.SetTotalRecords(totalRecords)

	return accounts, nil
}

func UpdateAccount(ctx context.Context, a *models.Account) error {
	query := `
		UPDATE account
		SET
            type = $1,
            name = $2,
            balance = $3,
            currency_code = $4,
            user_id = $5,
            external_id = $6,
            subtype = $7,
            number = $8,
            owner = $9,
            tax_number = $10,
            bank_account_data_id = $11,
            credit_account_data_id = $12,
            fi_connection_id = $13
		WHERE user_id = $14 AND id = $15
    `

	_, err := conn.Exec(
		ctx, query,
		a.Type,
		a.Name,
		a.Balance,
		a.CurrencyCode,
		a.UserID,
		a.ExternalID,
		a.Subtype,
		a.Number,
		a.Owner,
		a.TaxNumber,
		a.BankAccountDataID,
		a.CreditAccountDataID,
		a.FiConnectionID,
		a.UserID,
		a.ID,
	)
	if err != nil {
		return err
	}

	return nil
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
