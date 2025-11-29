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
	PaymentIntentStatusFrozen        PaymentIntentStatus = "Frozen"         // Payment received but funds frozen due to AML rejection
	PaymentIntentStatusUnfrozen      PaymentIntentStatus = "Unfrozen"       // Previously frozen funds have been successfully unfrozen
	PaymentIntentStatusCancelled     PaymentIntentStatus = "Cancelled"      // Cancelled due to timeout: 20min for Created, 60min for Partially Paid/Paid, 120min for Completed
)

type ProductOptionSwatch struct {
	Type    string `json:"type"`               // "color" or "image"
	Value   string `json:"value,omitempty"`    // Hex color or display value
	MediaID string `json:"media_id,omitempty"` // Reference to media asset when type=image
}

type ProductOptionValue struct {
	Value     string               `json:"value"`                // Option value label
	SortOrder int                  `json:"sort_order,omitempty"` // Optional order for display
	Swatch    *ProductOptionSwatch `json:"swatch,omitempty"`     // Optional swatch metadata
	Metadata  map[string]string    `json:"metadata,omitempty"`   // Additional metadata
}

type ProductOption struct {
	Name      string                `json:"name"`                 // Option name, e.g. "Color"
	SortOrder int                   `json:"sort_order,omitempty"` // Optional order for display
	Values    []*ProductOptionValue `json:"values"`               // Allowed values for this option
	Metadata  map[string]string     `json:"metadata,omitempty"`   // Optional metadata
}

type ProductVariant struct {
	ID                  string                       `json:"id"`                               // Unique identifier for the variant
	OptionValues        map[string]string            `json:"option_values"`                    // Selected option value per option name
	Price               *AssetAmount                 `json:"price,omitempty"`                  // Variant price
	MediaOrder          []string                     `json:"media_order,omitempty"`            // Ordered media asset IDs for the variant
	InventoryQuantity   *int                         `json:"inventory_quantity,omitempty"`     // Optional inventory quantity
	Active              bool                         `json:"active"`                           // Variant is active/available
	Version             int64                        `json:"version"`                          // Catalog version that produced this variant
	SellingPlanGroupIDs []string                     `json:"selling_plan_group_ids,omitempty"` // Variant-level selling plan group associations
	SellingPlanGroups   []*SellingPlanGroupWithPlans `json:"selling_plan_groups,omitempty"`    // Expanded selling plan groups with plans (only populated when variant has override)
	SellingPlanPricing  []*SellingPlanPricing        `json:"selling_plan_pricing,omitempty"`   // Pre-calculated subscription pricing per selling plan
}

type SellingPlanGroupWithPlans struct{} // Placeholder for subscription feature
type SellingPlanPricing struct{}        // Placeholder for subscription feature

type Product struct {
	ID                     string                       `json:"id"`                               // Unique identifier for the product
	Name                   string                       `json:"name"`                             // Name of the product
	Price                  *AssetAmount                 `json:"price,omitempty"`                  // Legacy price field for simple products
	FixedPrice             *AssetAmount                 `json:"fixed_price,omitempty"`            // Explicit price for simple products
	Description            string                       `json:"description"`                      // Detailed description of the product
	TaxCode                string                       `json:"tax_code"`                         // Tax classification code for the product
	Metadata               map[string]string            `json:"metadata"`                         // Additional custom data for the product
	DefaultCurrency        string                       `json:"default_currency"`                 // Base currency for the product/variants
	MediaOrder             []string                     `json:"media_order,omitempty"`            // Ordered media asset IDs at product level
	CollectShippingAddress bool                         `json:"collect_shipping_address"`         // Collect shipping address during checkout
	CollectTaxAddress      bool                         `json:"collect_tax_address"`              // Collect billing/tax address during checkout
	RequiresSellingPlan    bool                         `json:"requires_selling_plan"`            // true=subscription only, false=one-time or subscription
	Options                []*ProductOption             `json:"options,omitempty"`                // Variant option definitions
	Variants               []*ProductVariant            `json:"variants,omitempty"`               // Variant combinations
	Active                 bool                         `json:"active"`                           // Indicates if the product is currently available
	UpdatedAt              int64                        `json:"updated_at"`                       // Last update timestamp in Unix format
	CreatedAt              int64                        `json:"created_at"`                       // Creation timestamp in Unix format
	Version                int64                        `json:"version"`                          // Catalog version associated with this product
	Includes               *PaymentLinkIncludes         `json:"includes,omitempty"`               // Included media assets with signed URLs
	SellingPlanGroupIDs    []string                     `json:"selling_plan_group_ids,omitempty"` // Product-level selling plan group associations
	SellingPlanGroups      []*SellingPlanGroupWithPlans `json:"selling_plan_groups,omitempty"`    // Expanded selling plan groups with plans (use expand=selling_plans query param)
}

type ProductItem struct {
	Product  *Product `json:"product"`  // Referenced product information
	Quantity int      `json:"quantity"` // Number of product items
}

type PaymentLinkVariant struct {
	VariantID   string          `json:"variant_id"`             // Identifier of the product variant
	Variant     *ProductVariant `json:"variant,omitempty"`      // Variant details returned for convenience
	Quantity    int             `json:"quantity"`               // Default quantity for this variant
	MinQuantity *int            `json:"min_quantity,omitempty"` // Optional minimum quantity for add-ons
	MaxQuantity *int            `json:"max_quantity,omitempty"` // Optional maximum quantity for add-ons
	IsPrimary   bool            `json:"is_primary"`             // Indicates if this variant is part of the primary selection
}

type PaymentLinkPriceRange struct {
	Min *AssetAmount `json:"min"` // Minimum price across primary variants
	Max *AssetAmount `json:"max"` // Maximum price across primary variants
}

type PaymentLinkMedia struct {
	ID          string `json:"id"`                     // Media identifier
	URL         string `json:"url"`                    // CDN URL for the asset
	Width       int    `json:"width,omitempty"`        // Width in pixels
	Height      int    `json:"height,omitempty"`       // Height in pixels
	AltText     string `json:"alt_text,omitempty"`     // Alternative text for accessibility
	ContentType string `json:"content_type,omitempty"` // MIME type of the asset
}

type PaymentLinkIncludes struct {
	Media []*PaymentLinkMedia `json:"media,omitempty"` // Media assets referenced by the product/variants
}

type PaymentLink struct {
	ID              string                 `json:"id"`                         // Unique identifier for the payment link
	ProductItems    []*ProductItem         `json:"product_items,omitempty"`    // Legacy flattened product list
	Product         *Product               `json:"product,omitempty"`          // Product associated with the payment link
	PrimaryVariants []*PaymentLinkVariant  `json:"primary_variants,omitempty"` // Primary variants selectable by the buyer
	AddonVariants   []*PaymentLinkVariant  `json:"addon_variants,omitempty"`   // Optional add-on variants
	VariantConfig   map[string]any         `json:"variant_config,omitempty"`   // UI and default selection metadata
	PriceRange      *PaymentLinkPriceRange `json:"price_range,omitempty"`      // Price range across primary variants
	TotalPrice      *AssetAmount           `json:"total_price,omitempty"`      // Computed total price (legacy behaviour)
	Active          bool                   `json:"active"`                     // Indicates if the payment link is currently active
	UpdatedAt       int64                  `json:"updated_at"`                 // Last update timestamp in Unix format
	CreatedAt       int64                  `json:"created_at"`                 // Creation timestamp in Unix format
	URL             *string                `json:"url"`                        // URL to access the payment link
	Includes        *PaymentLinkIncludes   `json:"includes,omitempty"`         // Included shared resources (e.g., media)
}

type PaymentLinkDetails struct {
	MerchantID   string       `json:"merchant_id"`   // ID of the merchant who created the payment link
	MerchantName string       `json:"merchant_name"` // Name of the merchant
	PublicKey    string       `json:"public_key"`    // Public API key for this merchant
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
	Fiat   *FiatOption   `json:"fiat"`   // Simplified fiat options used for payment method confirmation, containing only the asset ID
}

type FiatOption struct {
	Currency          *string `json:"currency" binding:"required"`   // Currency code for the fiat
	PaymentMethodID   *string `json:"payment_method_id,omitempty"`   // Saved payment method ID for Stripe (used with cards payment type)
	SavePaymentMethod *bool   `json:"save_payment_method,omitempty"` // Whether to save the payment method for future use (Stripe only)
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
	Amount                 string            `json:"amount" binding:"required" example:"1.0"`                      // Amount to be charged in the specified currency
	Currency               string            `json:"currency" binding:"required" example:"USD"`                    // ISO 4217 currency code
	Customer               *string           `json:"customer"`                                                     // Identifier of the customer this payment intent belongs to
	Description            *string           `json:"description"`                                                  // An arbitrary string attached to the object. Often useful for displaying to users.
	Metadata               map[string]string `json:"metadata"`                                                     // Set of key-value pairs that you can attach to the object
	MerchantOrderId        string            `json:"merchant_order_id" binding:"required" example:"sn-1234567890"` // Unique identifier for the order on the merchant's side
	ReceiptEmail           string            `json:"receipt_email" binding:"omitempty,email"`                      // Email address to send the receipt to
	ReturnURL              *string           `json:"return_url"`                                                   // URL to redirect the customer to after payment
	PaymentMethodID        *string           `json:"payment_method_id"`                                            // ID of a saved payment method to use for this payment (Stripe only)
	CompleteOnFirstPayment *bool             `json:"complete_on_first_payment" swaggerignore:"true"`               // Complete payment immediately when first transaction arrives, regardless of amount
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

// UnfreezeWithdraw represents an unfreeze withdrawal of previously frozen funds
type UnfreezeWithdraw struct {
	ID                 string       `json:"id"`                    // Unique identifier for the withdraw
	PayoutID           string       `json:"payout_id"`             // Associated payout ID
	AssetID            string       `json:"asset_id"`              // Asset ID of the withdrawn funds
	Amount             *AssetAmount `json:"amount"`                // Amount being withdrawn
	NetworkFee         *AssetAmount `json:"network_fee"`           // Network fee for the withdrawal
	Status             string       `json:"status"`                // Status of the withdrawal
	Address            string       `json:"address"`               // Destination address for withdrawal
	Type               string       `json:"type"`                  // Type of withdraw (unfreeze_reverse or unfreeze_release)
	Description        string       `json:"description"`           // Description from the associated payout
	OriginalFrozenTxID string       `json:"original_frozen_tx_id"` // Original AML-rejected transaction ID
	CreatedAt          int64        `json:"created_at"`            // Created timestamp
	UpdatedAt          int64        `json:"updated_at"`            // Last updated timestamp
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
	CompleteOnFirstPayment  bool                   `json:"complete_on_first_payment"`  // Complete payment immediately when first transaction arrives, regardless of amount
	PermanentDeposit        bool                   `json:"permanent_deposit"`          // Indicates if this payment is a permanent deposit
	PermanentDepositAssetId string                 `json:"permanent_deposit_asset_id"` // Asset ID for the permanent deposit
	ExpiredAt               uint64                 `json:"expired_at"`                 // Expired time
	UnfreezeWithdraws       []*UnfreezeWithdraw    `json:"unfreeze_withdraws"`         // List of unfreeze withdrawals for previously frozen funds

	// Association fields (populated on demand via expand parameters)
	Invoice             *string     `json:"invoice,omitempty"`              // Invoice ID if this payment is for an invoice (subscriptions)
	InvoiceDetails      interface{} `json:"invoice_details,omitempty"`      // Expanded invoice details (use expand=invoice)
	Subscription        *string     `json:"subscription,omitempty"`         // Subscription ID (resolved through invoice)
	SubscriptionDetails interface{} `json:"subscription_details,omitempty"` // Expanded subscription details (use expand=invoice.subscription)
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
