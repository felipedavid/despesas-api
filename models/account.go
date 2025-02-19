package models

import "time"

type AccountType string

const (
	AccountTypeBankAccount AccountType = "BANK"
	AccountTypeCreditCard  AccountType = "CREDIT"
)

var validAccountTypeMap = map[AccountType]struct{}{
	AccountTypeBankAccount: {},
	AccountTypeCreditCard:  {},
}

func ValidAccountType(at string) bool {
	_, exists := validAccountTypeMap[AccountType(at)]
	return exists
}

type AccountSubtype string

const (
	AccountSubtypeCheckingAccount AccountSubtype = "CHECKING_ACCOUNT"
	AccountSubtypeSavingsAccount  AccountSubtype = "SAVINGS_ACCOUNT"
	AccountSubtypeCreditCard      AccountSubtype = "CREDIT_CARD"
)

var validAccountSubtypeMap = map[AccountSubtype]struct{}{
	AccountSubtypeCheckingAccount: {},
	AccountSubtypeSavingsAccount:  {},
	AccountSubtypeCreditCard:      {},
}

func ValidAccountSubtype(ast string) bool {
	_, exists := validAccountSubtypeMap[AccountSubtype(ast)]
	return exists
}

type Account struct {
	ID                  string          `db:"id" json:"id"`
	Type                AccountType     `db:"type" json:"type"`
	Name                string          `db:"name" json:"name"`
	Balance             int             `db:"balance" json:"balance"`
	CurrencyCode        string          `db:"currency_code" json:"currency_code"`
	UserID              string          `db:"user_id" json:"user_id"`
	ExternalID          *string         `db:"external_id" json:"-"`
	Subtype             *AccountSubtype `db:"subtype" json:"subtype"`
	Number              *string         `db:"number" json:"-"`
	Owner               *string         `db:"owner" json:"-"`
	TaxNumber           *string         `db:"tax_number" json:"-"`
	BankAccountDataID   *string         `db:"bank_account_data_id" json:"-"`
	CreditAccountDataID *string         `db:"credit_account_data_id" json:"-"`
	FiConnectionID      *string         `db:"fi_connection_id" json:"-"`

	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}

type AccountNullable struct {
	ID                  *string         `db:"id"`
	Type                *AccountType    `db:"type"`
	Name                *string         `db:"name"`
	Balance             *int            `db:"balance"`
	CurrencyCode        *string         `db:"currency_code"`
	UserID              *string         `db:"user_id"`
	ExternalID          *string         `db:"external_id"`
	Subtype             *AccountSubtype `db:"subtype"`
	Number              *string         `db:"number"`
	Owner               *string         `db:"owner"`
	TaxNumber           *string         `db:"tax_number"`
	BankAccountDataID   *string         `db:"bank_account_data_id"`
	CreditAccountDataID *string         `db:"credit_account_data_id"`
	FiConnectionID      *string         `db:"fi_connection_id"`

	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
