package developer

type ChargeFraudMartianReport string

const (
	ChargeFraudMartianReportFraudulent ChargeFraudMartianReport = "fraudulent"
)

type ChargeFraudUserReport string

const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

type ChargeFraudDetails struct {
	MartianReport ChargeFraudMartianReport `json:"martian_report"`
	UserReport    ChargeFraudUserReport    `json:"user_report"`
}

type ReviewClosedReason string

const (
	ReviewClosedReasonApproved        ReviewClosedReason = "approved"
	ReviewClosedReasonDisputed        ReviewClosedReason = "disputed"
	ReviewClosedReasonRedacted        ReviewClosedReason = "redacted"
	ReviewClosedReasonRefunded        ReviewClosedReason = "refunded"
	ReviewClosedReasonRefundedAsFraud ReviewClosedReason = "refunded_as_fraud"
)

type ReviewIPAddressLocation struct {
	// The city where the payment originated.
	City string `json:"city"`
	// Two-letter ISO code representing the country where the payment originated.
	Country string `json:"country"`
	// The geographic latitude where the payment originated.
	Latitude float64 `json:"latitude"`
	// The geographic longitude where the payment originated.
	Longitude float64 `json:"longitude"`
	// The state/county/province/region where the payment originated.
	Region string `json:"region"`
}

type ReviewOpenedReason string

// List of values that ReviewOpenedReason can take
const (
	ReviewOpenedReasonManual ReviewOpenedReason = "manual"
	ReviewOpenedReasonRule   ReviewOpenedReason = "rule"
)

type ReviewReason string

// List of values that ReviewReason can take
const (
	ReviewReasonApproved        ReviewReason = "approved"
	ReviewReasonDisputed        ReviewReason = "disputed"
	ReviewReasonManual          ReviewReason = "manual"
	ReviewReasonRefunded        ReviewReason = "refunded"
	ReviewReasonRefundedAsFraud ReviewReason = "refunded_as_fraud"
	ReviewReasonRedacted        ReviewReason = "redacted"
	ReviewReasonRule            ReviewReason = "rule"
)

type ReviewSession struct {
	// The browser used in this browser session (e.g., `Chrome`).
	Browser string `json:"browser"`
	// Information about the device used for the browser session (e.g., `Samsung SM-G930T`).
	Device string `json:"device"`
	// The platform for the browser session (e.g., `Macintosh`).
	Platform string `json:"platform"`
	// The version for the browser session (e.g., `61.0.3163.100`).
	Version string `json:"version"`
}

// 争议交易客户的补充信息
type Review struct {
	ID                string                   `json:"id"`
	Object            string                   `json:"object"` // review
	Charge            *string                  `json:"charge"`
	ClosedReason      ReviewClosedReason       `json:"closed_reason"`
	Created           int64                    `json:"created"`
	IPAddress         string                   `json:"ip_address"`
	IPAddressLocation *ReviewIPAddressLocation `json:"ip_address_location"`
	Livemode          bool                     `json:"livemode"`
	Open              bool                     `json:"open"`
	// The reason the review was opened. One of `rule` or `manual`.
	OpenedReason ReviewOpenedReason `json:"opened_reason"`
	// The PaymentIntent ID associated with this review, if one exists.
	PaymentIntent *string `json:"payment_intent"` // paymentintent id
	// The reason the review is currently open or closed. One of `rule`, `manual`, `approved`, `refunded`, `refunded_as_fraud`, `disputed`, or `redacted`.
	Reason ReviewReason `json:"reason"`
	// Information related to the browsing session of the user who initiated the payment.
	Session *ReviewSession `json:"session"`
}

type ChargeStatus string

// List of values that ChargeStatus can take
const (
	ChargeStatusFailed    ChargeStatus = "failed"
	ChargeStatusPending   ChargeStatus = "pending"
	ChargeStatusSucceeded ChargeStatus = "succeeded"
)

const (
	ChargeObject = "charge"
)

type Charge struct {
	ID     string       `json:"id"`
	Object string       `json:"object"` // charge
	Amount *AssetAmount `json:"amount"`

	PaymentDetails *PaymentDetails `json:"payment_details"`
	ExchangeRate   string          `json:"exchange_rate"`

	CalculatedStatementDescriptor string              `json:"calculated_statement_descriptor"` // webesim.com
	Captured                      bool                `json:"captured"`                        // true
	Created                       int64               `json:"created"`
	Customer                      string              `json:"customer"`
	Description                   string              `json:"description"`
	Disputed                      bool                `json:"disputed"` // false, 是否有争议
	FailureCode                   string              `json:"failure_code"`
	FailureMessage                string              `json:"failure_message"`
	FraudDetails                  *ChargeFraudDetails `json:"fraud_details"`
	Livemode                      bool                `json:"livemode"`
	Metadata                      map[string]string   `json:"metadata"`

	// Outcome                       *ChargeOutcome        `json:"outcome"`

	Paid          bool    `json:"paid"`
	PaymentIntent *string `json:"payment_intent"` // paymentintent id

	PaymentMethodType    PaymentMethodType     `json:"payment_method_type"`
	PaymentMethodOptions *PaymentMethodOptions `json:"payment_method_options"`
	Transactions         []*TransactionDetails `json:"transactions"` // 区块链交易列表

	ReceiptEmail string `json:"receipt_email"`
	ReceiptURL   string `json:"receipt_url"` // This is the URL to view the receipt for this charge. The receipt is kept up-to-date to the latest state of the charge, including any refunds. If the charge is for an Invoice, the receipt will be stylized as an Invoice receipt.

	Refunded bool `json:"refunded"`

	Refunds []*Refund `json:"refunds"` // list of refund ids
	Review  *Review   `json:"review"`  // Review associated with this charge if one exists.
}
