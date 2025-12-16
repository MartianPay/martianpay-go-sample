// deposit.go contains types for managing cryptocurrency deposit accounts and transactions.
// It includes structures for tracking deposit addresses, blockchain transactions, and AML screening results.
package developer

// DepositAccount represents a cryptocurrency deposit account for receiving payments.
// It is used to track deposits to specific blockchain addresses, whether temporary (for single payments)
// or permanent (for recurring deposits with unique identifiers).
type DepositAccount struct {
	// MerchantId is the unique identifier of the merchant
	MerchantId string `json:"merchant_id"`
	// UserId is the unique identifier of the user
	UserId string `json:"user_id"`
	// ChargeId is the unique identifier of the associated charge
	ChargeId string `json:"charge_id"`
	// AssetId is the unique identifier of the cryptocurrency asset
	AssetId string `json:"asset_id"`
	// DepositAddress is the blockchain address for receiving deposits
	DepositAddress string `json:"deposit_address"`
	// DepositedAmount is the total amount deposited to this address
	DepositedAmount string `json:"deposited_amount"`
	// Transactions is the list of blockchain transactions on this deposit address
	Transactions []*Transaction `json:"transactions,omitempty"`
	// IsPermanent indicates whether this is a permanent deposit address
	IsPermanent bool `json:"is_permanent,omitempty"`
	// UniqueId is the unique identifier for permanent deposit addresses (only valid when IsPermanent is true)
	UniqueId string `json:"unique_id,omitempty"`
}

// AmlInfo contains Anti-Money Laundering (AML) screening information
type AmlInfo struct {
	// Score is the AML risk score from 0-10 (10 is the highest risk)
	Score float64 `json:"score,omitempty"`
	// RuleNames contains the list of AML rules that were triggered
	RuleNames []string `json:"rule_names,omitempty"`
}

// FeeInfo contains fee information for a blockchain transaction
type FeeInfo struct {
	// NetworkFee is the fee paid to the blockchain network
	NetworkFee string `json:"network_fee,omitempty"`
	// ServiceFee is the total fee deducted by the platform (serviceFee = amount - netAmount)
	ServiceFee string `json:"service_fee,omitempty"`
}

// Transaction represents a blockchain transaction
type Transaction struct {
	// SourceAddress is the address sending the funds
	SourceAddress string `json:"source_address,omitempty"`
	// DestinationAddress is the address receiving the funds
	DestinationAddress string `json:"destination_address,omitempty"`
	// TxHash is the transaction hash on the blockchain
	TxHash string `json:"tx_hash,omitempty"`
	// Amount is the transaction amount
	Amount string `json:"amount,omitempty"`
	// AssetId is the unique identifier of the asset (maps to token symbol and network)
	AssetId string `json:"asset_id,omitempty"`
	// Type indicates the transaction type (0: deposit, 1: refund)
	Type int32 `json:"type,omitempty"`
	// CreatedAt is the Unix timestamp when the transaction was created
	CreatedAt int64 `json:"created_at,omitempty"`
	// Status is the transaction status ("submitted", "failed", "completed", "confirmed")
	// Transaction is successful when status is "confirmed" and AmlStatus is "approved"
	Status string `json:"status,omitempty"`
	// AmlStatus is the AML screening status ('', 'approved', 'rejected')
	AmlStatus string `json:"aml_status,omitempty"`
	// AmlInfo contains detailed AML screening information
	AmlInfo *AmlInfo `json:"aml_info,omitempty"`
	// ChargeId is the unique identifier of the associated charge
	ChargeId string `json:"charge_id,omitempty"`
	// RefundId is the unique identifier of the associated refund (only valid when Type==1)
	RefundId string `json:"refund_id,omitempty"`
	// FeeInfo contains detailed fee information for the transaction
	FeeInfo *FeeInfo `json:"fee_info,omitempty"`
	// FeeCurrency is the asset used to pay the transaction fee (e.g., ETH for ERC-20 tokens)
	FeeCurrency string `json:"fee_currency,omitempty"`
}
