package storage

import (
	"context"
	"testing"

	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/test"
)

func createTestUser(t *testing.T) *models.User {
	t.Helper()

	user := &models.User{
		Name:     "User " + t.Name(),
		Email:    t.Name() + "@gmail.com",
		Password: []byte(t.Name()),
	}

	err := InsertUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Error inserting user: %v", err)
	}

	return user
}

func TestInsertAccount(t *testing.T) {
	user := createTestUser(t)

	newAccount := &models.Account{
		UserID:              user.ID,
		Type:                "CREDIT",
		Name:                "Nubank",
		Balance:             1212,
		CurrencyCode:        "BRL",
		ExternalID:          nil,
		Subtype:             nil,
		Number:              nil,
		Owner:               nil,
		TaxNumber:           nil,
		BankAccountDataID:   nil,
		CreditAccountDataID: nil,
		FiConnectionID:      nil,
	}

	err := InsertAccount(context.Background(), newAccount)
	if err != nil {
		t.Errorf("Error inserting account: %v", err)
	}

	createdAccount, err := GetUserAccount(context.Background(), user.ID, newAccount.ID)
	if err != nil {
		t.Errorf("Error getting account: %v", err)
	}

	test.Equal(t, newAccount.UserID, createdAccount.UserID)
	test.Equal(t, newAccount.Type, createdAccount.Type)
	test.Equal(t, newAccount.Name, createdAccount.Name)
	test.Equal(t, newAccount.Balance, createdAccount.Balance)
	test.Equal(t, newAccount.CurrencyCode, createdAccount.CurrencyCode)
	test.Equal(t, newAccount.ExternalID, createdAccount.ExternalID)
	test.Equal(t, newAccount.Subtype, createdAccount.Subtype)
	test.Equal(t, newAccount.Number, createdAccount.Number)
	test.Equal(t, newAccount.Owner, createdAccount.Owner)
	test.Equal(t, newAccount.TaxNumber, createdAccount.TaxNumber)
	test.Equal(t, newAccount.BankAccountDataID, createdAccount.BankAccountDataID)
	test.Equal(t, newAccount.CreditAccountDataID, createdAccount.CreditAccountDataID)
	test.Equal(t, newAccount.FiConnectionID, createdAccount.FiConnectionID)
}

//func TestGetUserAccount(t *testing.T) {
//	t.Parallel()
//}
//
//func TestListUserAccounts(t *testing.T) {
//	t.Parallel()
//}
//
//func TestUpdateAccount(t *testing.T) {
//	t.Parallel()
//}
//
//func TestDeleteAccount(t *testing.T) {
//	t.Parallel()
//}
