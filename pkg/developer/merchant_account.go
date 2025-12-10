package developer

import (
	"github.com/dchest/uniuri"
)

const (
	MerchantAccountTypeWallet = "wallet"
	MerchantAccountTypeBank   = "bank"

	MerchantAccountObject   = "merchant_account"
	MerchantAccountPrefix   = "ma_"
	MerchantAccountIDLength = 24
)

func GenerateMerchantAccountID() string {
	return MerchantAccountPrefix + uniuri.NewLen(MerchantAccountIDLength)
}

// MerchantAccount represents a merchant's payout account (bank or wallet)
type MerchantAccount struct {
	// ID is the unique identifier for the merchant account
	ID string `json:"id"`
	// Object is the type identifier, always "merchant_account"
	Object string `json:"object"`
	// CreatedAt is the Unix timestamp when the account was created
	CreatedAt int64 `json:"created_at"`
	// AccountType is the type of account ("wallet" or "bank")
	AccountType string `json:"account_type"`
	// Alias is the user-defined name for this account
	Alias string `json:"alias"`
	// IsDefault indicates whether this is the default payout account
	IsDefault bool `json:"is_default"`
	// AccountStatus is the current status of the account
	AccountStatus string `json:"account_status"`
	// BankAccount contains bank account details (only for bank account type)
	BankAccount *BankAccount `json:"bank_account,omitempty"`
}

// BankAccount represents bank account details for payouts
type BankAccount struct {
	// AccountHolderName is the name of the account holder
	AccountHolderName string `json:"account_holder_name"`
	// BankName is the name of the bank
	BankName string `json:"bank_name"`
	// BranchName is the name of the bank branch
	BranchName string `json:"branch_name"`
	// SwiftCode is the SWIFT/BIC code for international transfers
	SwiftCode string `json:"swift_code"`
	// AccountNumber is the bank account number
	AccountNumber string `json:"account_number"`
}
