package models

import "time"

type Transaction struct {
	ID              string    `db:"id" json:"id"`
	ExternalID      *int      `db:"external_id" json:"-"`
	UserID          string    `db:"user_id" json:"user_id"`
	AccountID       *string   `db:"account_id" json:"account_id"`
	Description     *string   `db:"description" json:"description"`
	Amount          int       `db:"amount" json:"amount"`
	CurrencyCode    string    `db:"currency_code" json:"currency_code"`
	TransactionDate time.Time `db:"transaction_date" json:"transaction_date"`
	CategoryID      *int      `db:"category_id" json:"category_id"`
	Status          *string   `db:"status" json:"status"`
	Type            *string   `db:"type" json:"type"`
	OperationType   *string   `db:"operation_type" json:"operation_type"`

	Account  *Account  `json:"account,omitempty"`
	Category *Category `json:"category,omitempty"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}
