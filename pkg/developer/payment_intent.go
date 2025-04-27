package developer

const (
	PaymentIntentObject = "payment_intent"
)

type PaymentIntentStatus string

const (
	PaymentIntentStatusCreated       PaymentIntentStatus = "Created"        // Just created, waiting for user to select payment method
	PaymentIntentStatusWaiting       PaymentIntentStatus = "Waiting"        // Payment method selected, waiting for payment
	PaymentIntentStatusPartiallyPaid PaymentIntentStatus = "Partially Paid" // At least one blockchain tx with status Submitted/Completed/Confirmed, but total amount insufficient
	PaymentIntentStatusPaid          PaymentIntentStatus = "Paid"           // At least one blockchain tx with status Submitted/Completed/Confirmed, with sufficient total amount
	PaymentIntentStatusCompleted     PaymentIntentStatus = "Completed"      // At least one blockchain tx with status Completed/Confirmed, with sufficient total amount
	PaymentIntentStatusConfirmed     PaymentIntentStatus = "Confirmed"      // At least one blockchain tx with status Confirmed, with sufficient total amount
	PaymentIntentStatusCancelled     PaymentIntentStatus = "Cancelled"      // Cancelled due to timeout: 20min for Created, 60min for Partially Paid/Paid, 120min for Completed
)

type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Price       *AssetAmount      `json:"price"`
	Description string            `json:"description"`
	PictureURL  string            `json:"picture_url"`
	TaxCode     string            `json:"tax_code"`
	Metadata    map[string]string `json:"metadata"`
	Active      bool              `json:"active"`
	UpdatedAt   int64             `json:"updated_at"`
	CreatedAt   int64             `json:"created_at"`
}

type ProductItem struct {
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
}

type PaymentLink struct {
	ID           string         `json:"id"`
	ProductItems []*ProductItem `json:"product_items"`
	TotalPrice   *AssetAmount   `json:"total_price"`
	Active       bool           `json:"active"`
	UpdatedAt    int64          `json:"updated_at"`
	CreatedAt    int64          `json:"created_at"`
}

type PaymentLinkDetails struct {
	MerchantID   string       `json:"merchant_id"`
	MerchantName string       `json:"merchant_name"`
	PaymentLink  *PaymentLink `json:"payment_link"`
}

type TransactionDetails struct {
	TxId               string  `json:"tx_id"`
	SourceAddress      string  `json:"source_address"`
	DestinationAddress string  `json:"destination_address"`
	TxHash             string  `json:"tx_hash"`
	Amount             string  `json:"amount"`
	Decimals           int     `json:"decimals"`
	AssetId            string  `json:"asset_id"` // Asset ID for the token
	Token              string  `json:"token"`    // Token name, e.g. USDT, BTC, ETH, etc.
	Network            string  `json:"network"`  // Network type for USDT transfer, e.g. TRC20, ERC20, etc.
	Type               string  `json:"type"`     // 0: deposit, 1: refund
	CreatedAt          int64   `json:"created_at"`
	Status             string  `json:"status"`
	AmlStatus          string  `json:"aml_status"` //'', 'approved', 'rejected'.
	AmlInfo            *string `json:"aml_info"`   //AML detail
	ChargeId           string  `json:"charge_id"`
	RefundId           string  `json:"refund_id"`    //only valid for type==1(refund transaction)
	FeeInfo            *string `json:"fee_info"`     //Details of the transaction's fee.
	FeeCurrency        string  `json:"fee_currency"` //The asset type used to pay the fee (ETH for ERC-20 tokens, BTC for Omni, XLM for Stellar tokens, etc.)
}

type PaymentMethodType string

const (
	PaymentMethodTypeCrypto     PaymentMethodType = "crypto"
	PaymentMethodTypeVisa       PaymentMethodType = "visa"
	PaymentMethodTypeMastercard PaymentMethodType = "mastercard"
	PaymentMethodTypeApplePay   PaymentMethodType = "apple pay"
	PaymentMethodTypeGooglePay  PaymentMethodType = "google pay"
)

type PaymentMethodOptions struct {
	Crypto *Crypto `json:"crypto"`
}

type Crypto struct {
	Amount         *string `json:"amount"`          // Amount of crypto to be paid
	Token          *string `json:"token"`           // Token name, e.g. USDT, BTC, ETH, etc.
	AssetId        *string `json:"asset_id"`        // Asset ID for the token
	Network        *string `json:"network"`         // Network type for USDT transfer, e.g. TRC20, ERC20, etc.
	Decimals       *int    `json:"decimals"`        // Decimals for the token
	ExchangeRate   *string `json:"exchange_rate"`   // Exchange rate for USDT transfer
	DepositAddress *string `json:"deposit_address"` // Deposit address for USDT transfer
	ExpiredAt      *int64  `json:"expired_at"`      // Expired time
}

type PaymentIntentParams struct {
	Amount            string                `json:"amount" binding:"required"`
	Currency          string                `json:"currency" binding:"required"`
	Customer          *string               `json:"customer"`
	Description       *string               `json:"description"` // An arbitrary string attached to the object. Often useful for displaying to users.
	Metadata          map[string]string     `json:"metadata"`
	MerchantOrderId   string                `json:"merchant_order_id" binding:"required"`
	PaymentMethodType *PaymentMethodType    `json:"payment_method_type"` // crypto, visa, mastercard, apple pay, google pay, etc.
	PaymentMethodData *PaymentMethodOptions `json:"payment_method_options"`
	ReceiptEmail      *string               `json:"receipt_email"`
	ReturnURL         *string               `json:"return_url"`
	OneTimePayment    *bool                 `json:"one_time_payment"`
}

type PaymentDetails struct {
	Settled        bool                    `json:"-"`
	AmountCaptured *AssetAmount            `json:"amount_captured"`
	AmountRefunded *AssetAmount            `json:"amount_refunded"`
	TxFee          *AssetAmount            `json:"tx_fee"`
	TaxFee         *AssetAmount            `json:"tax_fee"`
	FrozenAmount   *AssetAmount            `json:"frozen_amount"`
	NetAmount      *AssetAmount            `json:"net_amount"`
	GasFee         map[string]*AssetAmount `json:"gas_fee"`
	NetworkFee     *AssetAmount            `json:"network_fee"`
}

type PaymentIntent struct {
	ID                  string                 `json:"id"`
	Object              string                 `json:"object"` // payment_intent
	Amount              *AssetAmount           `json:"amount"`
	PaymentDetails      *PaymentDetails        `json:"payment_details"`
	CanceledAt          int64                  `json:"canceled_at"`
	CancellationReason  string                 `json:"cancellation_reason"`
	ClientSecret        string                 `json:"client_secret"`
	Created             int64                  `json:"created"`
	Updated             int64                  `json:"updated"`
	Currency            string                 `json:"currency"`
	Customer            *Customer              `json:"customer"`
	Description         string                 `json:"description"`
	Livemode            bool                   `json:"livemode"` // Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Metadata            map[string]interface{} `json:"metadata"`
	PaymentLinkDetails  *PaymentLinkDetails    `json:"payment_link_details"`
	MerchantOrderId     string                 `json:"merchant_order_id"`
	PaymentMethodType   PaymentMethodType      `json:"payment_method_type"`
	Charges             []*Charge              `json:"charges"`
	ReceiptEmail        string                 `json:"receipt_email"`
	Status              string                 `json:"status"`
	ProcessingStatus    string                 `json:"-"`
	PaymentIntentStatus PaymentIntentStatus    `json:"payment_intent_status"`
	OneTimePayment      bool                   `json:"one_time_payment"`
}

type PaymentIntentCancellationReason string

const (
	PaymentIntentCancellationReasonAbandoned           PaymentIntentCancellationReason = "abandoned"
	PaymentIntentCancellationReasonAutomatic           PaymentIntentCancellationReason = "automatic"
	PaymentIntentCancellationReasonDuplicate           PaymentIntentCancellationReason = "duplicate"
	PaymentIntentCancellationReasonFailedInvoice       PaymentIntentCancellationReason = "failed_invoice"
	PaymentIntentCancellationReasonFraudulent          PaymentIntentCancellationReason = "fraudulent"
	PaymentIntentCancellationReasonRequestedByCustomer PaymentIntentCancellationReason = "requested_by_customer"
	PaymentIntentCancellationReasonVoidInvoice         PaymentIntentCancellationReason = "void_invoice"
)
