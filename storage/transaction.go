package storage

import (
	"context"

	"github.com/felipedavid/saldop/filters"
	"github.com/felipedavid/saldop/models"
)

func GetUserTransaction(ctx context.Context, userID, transactionID int) (*models.Transaction, error) {
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
		WHERE id = $1 AND user_id = $2 AND deleted_at IS NULL
    `

	var t models.Transaction
	err := conn.QueryRow(
		ctx, query,
		transactionID,
		userID,
	).Scan(
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

	return &t, nil
}

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

func ListUserTransactions(ctx context.Context, userID int, f *filters.Filters) ([]models.Transaction, error) {
	query := `
		SELECT
            count(*) OVER(),
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
        LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(ctx, query, userID, f.Limit(), f.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecords := 0
	transactions := make([]models.Transaction, 0, f.PageSize)
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(
			&totalRecords,
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

	f.SetTotalRecords(totalRecords)

	return transactions, nil
}

func ListUserTransactionsWithPopulatedFields(ctx context.Context, userID int, f *filters.Filters) ([]models.Transaction, error) {
	query := `
		SELECT
            count(*) OVER(),
			t.id,
			t.external_id,
			t.user_id,
			t.account_id,
			t.description,
			t.amount,
			t.currency_code,
			t.transaction_date,
			t.category_id,
			t.status,
			t.type,
			t.operation_type,
			t.created_at,
			t.updated_at,
			c.id,
            c.name,
            c.default_category,
            c.user_id,
			c.created_at,
			c.updated_at,
			c.deleted_at,
			a.id,
			a.type,
			a.name,
			a.balance,
			a.currency_code,
			a.user_id,
			a.external_id,
			a.subtype,
			a.number,
			a.owner,
			a.tax_number,
			a.bank_account_data_id,
			a.credit_account_data_id,
			a.fi_connection_id,
			a.created_at,
			a.updated_at
		FROM transaction t
        LEFT JOIN category c ON c.id = t.category_id
        LEFT JOIN account a ON a.id = t.account_id
		WHERE t.user_id = $1 AND t.deleted_at IS NULL
        LIMIT $2 OFFSET $3
	`

	rows, err := conn.Query(ctx, query, userID, f.Limit(), f.Offset())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecords := 0
	transactions := make([]models.Transaction, 0, f.PageSize)
	for rows.Next() {
		var t models.Transaction
		var c models.CategoryNullable
		var a models.AccountNullable

		err := rows.Scan(
			&totalRecords,
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
			&c.ID,
			&c.Name,
			&c.DefaultCategory,
			&c.UserID,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.DeletedAt,
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

		if c.ID != nil {
			t.Category = &models.Category{
				ID:              *c.ID,
				Name:            *c.Name,
				DefaultCategory: *c.DefaultCategory,
				UserID:          c.UserID,
				CreatedAt:       *c.CreatedAt,
				UpdatedAt:       *c.UpdatedAt,
				DeletedAt:       c.DeletedAt,
			}
		}

		if a.ID != nil {
			t.Account = &models.Account{
				ID:                  *a.ID,
				Type:                *a.Type,
				Name:                *a.Name,
				Balance:             *a.Balance,
				CurrencyCode:        *a.CurrencyCode,
				UserID:              *a.UserID,
				ExternalID:          a.ExternalID,
				Subtype:             a.Subtype,
				Number:              a.Number,
				Owner:               a.Owner,
				TaxNumber:           a.TaxNumber,
				BankAccountDataID:   a.BankAccountDataID,
				CreditAccountDataID: a.CreditAccountDataID,
				FiConnectionID:      a.FiConnectionID,
				CreatedAt:           *a.CreatedAt,
				UpdatedAt:           *a.UpdatedAt,
				DeletedAt:           a.DeletedAt,
			}
		}

		transactions = append(transactions, t)
	}

	f.SetTotalRecords(totalRecords)

	return transactions, nil
}

func UpdateTransaction(ctx context.Context, t *models.Transaction) error {
	query := `
		UPDATE transaction
		SET
            description      = $1,
            account_id       = $2,
            category_id      = $3,
            amount           = $4,
            currency_code    = $5,
            transaction_date = $6
		WHERE user_id = $7 AND id = $8
	`

	_, err := conn.Exec(
		ctx, query,
		t.Description,
		t.AccountID,
		t.CategoryID,
		t.Amount,
		t.CurrencyCode,
		t.TransactionDate,
		t.UserID,
		t.ID)
	if err != nil {
		return err
	}

	return nil
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
