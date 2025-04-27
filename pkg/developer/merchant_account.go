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

type MerchantAccount struct {
	ID            string       `json:"id"`
	Object        string       `json:"object"`
	CreatedAt     int64        `json:"created_at"`
	AccountType   string       `json:"account_type"`
	Alias         string       `json:"alias"`
	IsDefault     bool         `json:"is_default"`
	AccountStatus string       `json:"account_status"`
	BankAccount   *BankAccount `json:"bank_account,omitempty"`
}

type BankAccount struct {
	AccountHolderName string `json:"account_holder_name"`
	BankName          string `json:"bank_name"`
	BranchName        string `json:"branch_name"`
	SwiftCode         string `json:"swift_code"`
	AccountNumber     string `json:"account_number"`
}
