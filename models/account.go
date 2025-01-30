package models

import "time"

type Account struct {
	ID                  int     `db:"id" json:"id"`
	Type                string  `db:"type" json:"type"`
	Name                string  `db:"name" json:"name"`
	Balance             int     `db:"balance" json:"balance"`
	CurrencyCode        string  `db:"currency_code" json:"currency_code"`
	UserID              int     `db:"user_id" json:"user_id"`
	ExternalID          *string `db:"external_id" json:"-"`
	Subtype             *string `db:"subtype" json:"subtype"`
	Number              *string `db:"number" json:"-"`
	Owner               *string `db:"owner" json:"-"`
	TaxNumber           *string `db:"tax_number" json:"-"`
	BankAccountDataID   *int    `db:"bank_account_data_id" json:"-"`
	CreditAccountDataID *int    `db:"credit_account_data_id" json:"-"`
	FiConnectionID      *int    `db:"fi_connection_id" json:"-"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}

type AccountNullable struct {
	ID                  *int    `db:"id"`
	Type                *string `db:"type"`
	Name                *string `db:"name"`
	Balance             *int    `db:"balance"`
	CurrencyCode        *string `db:"currency_code"`
	UserID              *int    `db:"user_id"`
	ExternalID          *string `db:"external_id"`
	Subtype             *string `db:"subtype"`
	Number              *string `db:"number"`
	Owner               *string `db:"owner"`
	TaxNumber           *string `db:"tax_number"`
	BankAccountDataID   *int    `db:"bank_account_data_id"`
	CreditAccountDataID *int    `db:"credit_account_data_id"`
	FiConnectionID      *int    `db:"fi_connection_id"`

	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
