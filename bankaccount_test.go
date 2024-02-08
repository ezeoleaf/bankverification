package bankverification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankAccountValid(t *testing.T) {
	cases := []struct {
		bankAccount   string
		expectedValid bool
	}{
		{
			bankAccount:   "8323-6 988.123.838-4",
			expectedValid: false,
		},
		{
			bankAccount:   "8381-6",
			expectedValid: false,
		},
		{
			bankAccount:   "832799030933684",
			expectedValid: true,
		},
		{
			bankAccount:   "832799030933",
			expectedValid: true,
		},
		{
			bankAccount:   "8105923484362",
			expectedValid: false,
		},
		{
			bankAccount:   "9420 - 4172385",
			expectedValid: true,
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			ba := NewBankAccount(c.bankAccount)
			assert.Equal(t, c.expectedValid, ba.IsValid())
		})
	}
}

func TestBankAccountWithNewClearingNumbers(t *testing.T) {
	cases := []struct {
		bankAccount   string
		expectedValid bool
	}{
		{
			bankAccount:   "8323-6 988.123.838-4",
			expectedValid: false,
		},
		{
			bankAccount:   "8381-6",
			expectedValid: false,
		},
		{
			bankAccount:   "832799030933684",
			expectedValid: false,
		},
		{
			bankAccount:   "832799030933",
			expectedValid: false,
		},
		{
			bankAccount:   "8105923484362",
			expectedValid: false,
		},
		{
			bankAccount:   "9420 - 4172385",
			expectedValid: true,
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			cn := []ClearingNumber{
				{Interval: "9400..9449", BankName: "Forex Bank", AccountType: AccountTypes[AccountType1]},
			}

			ba := NewBankAccount(c.bankAccount)
			ba.SetClearingNumbers(cn)

			assert.Equal(t, c.expectedValid, ba.IsValid())
		})
	}
}
