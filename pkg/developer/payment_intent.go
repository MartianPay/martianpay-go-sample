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
	ID          string            `json:"id"`          // Unique identifier for the product
	Name        string            `json:"name"`        // Name of the product
	Price       *AssetAmount      `json:"price"`       // Price of the product including currency information
	Description string            `json:"description"` // Detailed description of the product
	PictureURL  string            `json:"picture_url"` // URL to the product image
	TaxCode     string            `json:"tax_code"`    // Tax classification code for the product
	Metadata    map[string]string `json:"metadata"`    // Additional custom data for the product
	Active      bool              `json:"active"`      // Indicates if the product is currently available
	UpdatedAt   int64             `json:"updated_at"`  // Last update timestamp in Unix format
	CreatedAt   int64             `json:"created_at"`  // Creation timestamp in Unix format
}

type ProductItem struct {
	Product  *Product `json:"product"`  // Referenced product information
	Quantity int      `json:"quantity"` // Number of product items
}

type PaymentLink struct {
	ID           string         `json:"id"`            // Unique identifier for the payment link
	ProductItems []*ProductItem `json:"product_items"` // List of products included in this payment link
	TotalPrice   *AssetAmount   `json:"total_price"`   // Total price for all products in the payment link
	Active       bool           `json:"active"`        // Indicates if the payment link is currently active
	UpdatedAt    int64          `json:"updated_at"`    // Last update timestamp in Unix format
	CreatedAt    int64          `json:"created_at"`    // Creation timestamp in Unix format
	URL          *string        `json:"url"`           // URL to access the payment link
}

type PaymentLinkDetails struct {
	MerchantID   string       `json:"merchant_id"`   // ID of the merchant who created the payment link
	MerchantName string       `json:"merchant_name"` // Name of the merchant
	PaymentLink  *PaymentLink `json:"payment_link"`  // Details of the payment link
}

type TransactionDetails struct {
	TxId               string  `json:"tx_id"`               // Transaction ID
	SourceAddress      string  `json:"source_address"`      // Source address
	DestinationAddress string  `json:"destination_address"` // Destination address
	TxHash             string  `json:"tx_hash"`             // Transaction hash
	Amount             string  `json:"amount"`              // Amount
	Decimals           int     `json:"decimals"`            // Decimals
	AssetId            string  `json:"asset_id"`            // Asset ID for the token
	Token              string  `json:"token"`               // Token name, e.g. USDT, BTC, ETH, etc.
	Network            string  `json:"network"`             // Network type for USDT transfer, e.g. TRC20, ERC20, etc.
	Type               string  `json:"type"`                // 0: deposit, 1: refund
	CreatedAt          int64   `json:"created_at"`          // Created time
	Status             string  `json:"status"`              // Transaction status
	AmlStatus          string  `json:"aml_status"`          //'', 'approved', 'rejected'.
	AmlInfo            *string `json:"aml_info"`            //AML detail
	ChargeId           string  `json:"charge_id"`           // Charge ID
	RefundId           string  `json:"refund_id"`           //only valid for type==1(refund transaction)
	FeeInfo            *string `json:"fee_info"`            //Details of the transaction's fee.
	FeeCurrency        string  `json:"fee_currency"`        //The asset type used to pay the fee (ETH for ERC-20 tokens, BTC for Omni, XLM for Stellar tokens, etc.)
}

type PaymentMethodType string

const (
	PaymentMethodTypeCrypto        PaymentMethodType = "crypto"
	PaymentMethodTypeTron          PaymentMethodType = "tron"
	PaymentMethodTypeWalletConnect PaymentMethodType = "wallet_connect"
	PaymentMethodTypeCards         PaymentMethodType = "cards"
	PaymentMethodTypeOthers        PaymentMethodType = "others"
)

type PaymentMethodOptions struct {
	Crypto *Crypto `json:"crypto"` // Crypto payment method options containing details like amount, token, network, etc.
	Fiat   *Fiat   `json:"fiat"`   // Fiat payment method options containing details like amount, token, network, etc.
}

type PaymentMethodConfirmOptions struct {
	Crypto *CryptoOption `json:"crypto"` // Simplified crypto options used for payment method confirmation, containing only the asset ID
	Fiat   *FiatOption   `json:"fiat""`  // Simplified fiat options used for payment method confirmation, containing only the asset ID
}

type FiatOption struct {
	Currency *string `json:"currency" binding:"required"` // Currency code for the fiat
}

type CryptoOption struct {
	AssetId *string `json:"asset_id" binding:"required"` // Asset ID for the token
}

type Fiat struct {
	Amount       *string `json:"amount"`                      // Amount of fiat to be paid
	AssetId      *string `json:"asset_id" binding:"required"` // Asset ID for the fiat
	Currency     *string `json:"currency" binding:"required"` // Currency code for the fiat
	Decimals     *int    `json:"decimals"`                    // Decimals for the fiat
	ExchangeRate *string `json:"exchange_rate"`               // Exchange rate for the fiat
}

type Crypto struct {
	Amount         *string `json:"amount"`                      // Amount of crypto to be paid
	Token          *string `json:"token"`                       // Token name, e.g. USDT, BTC, ETH, etc.
	AssetId        *string `json:"asset_id" binding:"required"` // Asset ID for the token
	Network        *string `json:"network"`                     // Network type for USDT transfer, e.g. TRC20, ERC20, etc.
	Decimals       *int    `json:"decimals"`                    // Decimals for the token
	ExchangeRate   *string `json:"exchange_rate"`               // Exchange rate for USDT transfer
	DepositAddress *string `json:"deposit_address"`             // Deposit address for USDT transfer
}

type PaymentIntentParams struct {
	Amount          string            `json:"amount" binding:"required" example:"1.0"`                      // Amount to be charged in the specified currency
	Currency        string            `json:"currency" binding:"required" example:"USD"`                    // ISO 4217 currency code
	Customer        *string           `json:"customer"`                                                     // Identifier of the customer this payment intent belongs to
	Description     *string           `json:"description"`                                                  // An arbitrary string attached to the object. Often useful for displaying to users.
	Metadata        map[string]string `json:"metadata"`                                                     // Set of key-value pairs that you can attach to the object
	MerchantOrderId string            `json:"merchant_order_id" binding:"required" example:"sn-1234567890"` // Unique identifier for the order on the merchant's side
	ReceiptEmail    string            `json:"receipt_email" binding:"omitempty,email"`                      // Email address to send the receipt to
	ReturnURL       *string           `json:"return_url"`                                                   // URL to redirect the customer to after payment
	OneTimePayment  *bool             `json:"one_time_payment" swaggerignore:"true"`                        // Whether this payment can only be used once
}

type PaymentDetails struct {
	Settled        bool                    `json:"-"`               // Indicates whether the payment has been settled
	AmountCaptured *AssetAmount            `json:"amount_captured"` // The amount that has been captured from the customer
	AmountRefunded *AssetAmount            `json:"amount_refunded"` // The amount that has been refunded to the customer
	TxFee          *AssetAmount            `json:"tx_fee"`          // Transaction fee charged by the payment processor
	TaxFee         *AssetAmount            `json:"tax_fee"`         // Tax amount collected for the transaction
	FrozenAmount   *AssetAmount            `json:"frozen_amount"`   // Amount temporarily held/frozen for pending operations
	NetAmount      *AssetAmount            `json:"net_amount"`      // Net amount after deducting all fees
	GasFee         map[string]*AssetAmount `json:"gas_fee"`         // Gas fees for blockchain transactions by asset type
	NetworkFee     *AssetAmount            `json:"network_fee"`     // Network fee for processing the transaction
}

type PaymentIntent struct {
	ID                      string                 `json:"id"`                         // Unique identifier for the payment intent
	Object                  string                 `json:"object"`                     // payment_intent
	Amount                  *AssetAmount           `json:"amount"`                     // Amount to be collected by this payment intent
	PaymentDetails          *PaymentDetails        `json:"payment_details"`            // Details about the payment including settled amounts and fees
	CanceledAt              int64                  `json:"canceled_at"`                // Timestamp when the payment intent was canceled
	CancellationReason      string                 `json:"cancellation_reason"`        // Reason why the payment intent was canceled
	ClientSecret            string                 `json:"client_secret"`              // Client secret used for client-side confirmation
	Created                 int64                  `json:"created"`                    // Timestamp when the payment intent was created
	Updated                 int64                  `json:"updated"`                    // Timestamp when the payment intent was last updated
	Currency                string                 `json:"currency"`                   // Currency of the payment intent
	Customer                *Customer              `json:"customer"`                   // Customer associated with this payment intent
	Description             string                 `json:"description"`                // Description of the payment intent
	Livemode                bool                   `json:"livemode"`                   // Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Metadata                map[string]interface{} `json:"metadata"`                   // Set of key-value pairs attached to the payment intent
	PaymentLinkDetails      *PaymentLinkDetails    `json:"payment_link_details"`       // Details about the payment link if this payment intent was created through a payment link
	MerchantOrderId         string                 `json:"merchant_order_id"`          // Merchant-specified order ID associated with this payment intent
	Charges                 []*Charge              `json:"charges"`                    // List of charges associated with this payment intent
	ReceiptEmail            string                 `json:"receipt_email"`              // Email address to send the receipt to
	Status                  string                 `json:"status"`                     // Current status of the payment intent
	ProcessingStatus        string                 `json:"-"`                          // Internal processing status of the payment intent
	PaymentIntentStatus     PaymentIntentStatus    `json:"payment_intent_status"`      // Standardized status of the payment intent
	OneTimePayment          bool                   `json:"one_time_payment"`           // Indicates if this is a one-time payment
	PermanentDeposit        bool                   `json:"permanent_deposit"`          // Indicates if this payment is a permanent deposit
	PermanentDepositAssetId string                 `json:"permanent_deposit_asset_id"` // Asset ID for the permanent deposit
	ExpiredAt               uint64                 `json:"expired_at"`                 // Expired time
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
