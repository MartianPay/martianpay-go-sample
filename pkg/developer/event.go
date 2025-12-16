// event.go contains types and utilities for webhook events and signature verification.
// It provides event type constants, event data structures, and HMAC signature validation
// for secure webhook processing.
package developer

import (
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// EventType describes the type of webhook event (for example, `invoice.created` or `payment_intent.succeeded`).
type EventType string

// Webhook event types
const (
	// Payment Intent events

	// EventTypePaymentIntentCreated is sent when a new payment intent is created
	EventTypePaymentIntentCreated EventType = "payment_intent.created"
	// EventTypePaymentIntentSucceeded is sent when a payment intent has been fully paid and completed successfully
	EventTypePaymentIntentSucceeded EventType = "payment_intent.succeeded"
	// EventTypePaymentIntentPaymentFailed is sent when a payment attempt for a payment intent fails
	EventTypePaymentIntentPaymentFailed EventType = "payment_intent.payment_failed"
	// EventTypePaymentIntentProcessing is sent when a payment intent is being processed (e.g., waiting for blockchain confirmation)
	EventTypePaymentIntentProcessing EventType = "payment_intent.processing"
	// EventTypePaymentIntentPartiallyPaid is sent when a payment intent has received partial payment but is not yet fully paid
	EventTypePaymentIntentPartiallyPaid EventType = "payment_intent.partially_paid"
	// EventTypePaymentIntentCanceled is sent when a payment intent is canceled
	EventTypePaymentIntentCanceled EventType = "payment_intent.canceled"

	// Refund events

	// EventTypeRefundCreated is sent when a new refund is created
	EventTypeRefundCreated EventType = "refund.created"
	// EventTypeRefundSucceeded is sent when a refund has been successfully processed and funds returned to customer
	EventTypeRefundSucceeded EventType = "refund.succeeded"
	// EventTypeRefundUpdated is sent when a refund's details are updated
	EventTypeRefundUpdated EventType = "refund.updated"
	// EventTypeRefundFailed is sent when a refund attempt fails
	EventTypeRefundFailed EventType = "refund.failed"

	// Payout events

	// EventTypePayoutCreated is sent when a new payout is created
	EventTypePayoutCreated EventType = "payout.created"
	// EventTypePayoutSucceeded is sent when a payout has been successfully transferred to the recipient
	EventTypePayoutSucceeded EventType = "payout.succeeded"
	// EventTypePayoutUpdated is sent when a payout's details are updated
	EventTypePayoutUpdated EventType = "payout.updated"
	// EventTypePayoutFailed is sent when a payout attempt fails
	EventTypePayoutFailed EventType = "payout.failed"

	// Payroll events

	// EventTypePayrollCreated is sent when a new payroll batch is created
	EventTypePayrollCreated EventType = "payroll.created"
	// EventTypePayrollApproved is sent when a payroll batch has been approved for execution
	EventTypePayrollApproved EventType = "payroll.approved"
	// EventTypePayrollRejected is sent when a payroll batch approval is rejected
	EventTypePayrollRejected EventType = "payroll.rejected"
	// EventTypePayrollCanceled is sent when a payroll batch is canceled
	EventTypePayrollCanceled EventType = "payroll.canceled"
	// EventTypePayrollExecuting is sent when a payroll batch execution has started
	EventTypePayrollExecuting EventType = "payroll.executing"
	// EventTypePayrollCompleted is sent when all items in a payroll batch have been processed successfully
	EventTypePayrollCompleted EventType = "payroll.completed"
	// EventTypePayrollFailed is sent when a payroll batch execution fails
	EventTypePayrollFailed EventType = "payroll.failed"

	// Payroll item events

	// EventTypePayrollItemProcessing is sent when an individual payroll item is being processed
	EventTypePayrollItemProcessing EventType = "payroll_item.processing"
	// EventTypePayrollItemSucceeded is sent when an individual payroll item has been successfully paid
	EventTypePayrollItemSucceeded EventType = "payroll_item.succeeded"
	// EventTypePayrollItemFailed is sent when an individual payroll item payment fails
	EventTypePayrollItemFailed EventType = "payroll_item.failed"
	// EventTypePayrollItemAddressVerification is sent when address verification email has been sent to the recipient
	EventTypePayrollItemAddressVerification EventType = "payroll_item.address_verification_sent"
	// EventTypePayrollItemAddressVerified is sent when the recipient has verified their wallet address
	EventTypePayrollItemAddressVerified EventType = "payroll_item.address_verified"

	// Subscription events

	// EventTypeSubscriptionCreated is sent when a new subscription is created
	EventTypeSubscriptionCreated EventType = "subscription.created"
	// EventTypeSubscriptionUpdated is sent when a subscription's details are updated (e.g., plan, quantity, billing cycle)
	EventTypeSubscriptionUpdated EventType = "subscription.updated"
	// EventTypeSubscriptionDeleted is sent when a subscription is deleted or permanently canceled
	EventTypeSubscriptionDeleted EventType = "subscription.deleted"
	// EventTypeSubscriptionPaused is sent when a subscription is temporarily paused
	EventTypeSubscriptionPaused EventType = "subscription.paused"
	// EventTypeSubscriptionResumed is sent when a paused subscription is resumed
	EventTypeSubscriptionResumed EventType = "subscription.resumed"
	// EventTypeSubscriptionTrialWill is sent when a subscription's trial period is about to end (typically 3 days before)
	EventTypeSubscriptionTrialWill EventType = "subscription.trial_will_end"

	// Invoice events

	// EventTypeInvoiceCreated is sent when a new invoice is created (draft state)
	EventTypeInvoiceCreated EventType = "invoice.created"
	// EventTypeInvoiceFinalized is sent when an invoice is finalized and ready for payment
	EventTypeInvoiceFinalized EventType = "invoice.finalized"
	// EventTypeInvoicePaid is sent when an invoice has been fully paid
	EventTypeInvoicePaid EventType = "invoice.paid"
	// EventTypeInvoicePaymentSucceeded is sent when a payment attempt for an invoice succeeds
	EventTypeInvoicePaymentSucceeded EventType = "invoice.payment_succeeded"
	// EventTypeInvoicePaymentFailed is sent when a payment attempt for an invoice fails
	EventTypeInvoicePaymentFailed EventType = "invoice.payment_failed"
	// EventTypeInvoicePaymentActionRequired is sent when an invoice payment requires additional action from the customer
	EventTypeInvoicePaymentActionRequired EventType = "invoice.payment_action_required"
	// EventTypeInvoiceUpcoming is sent when an upcoming invoice will be generated soon (for subscriptions)
	EventTypeInvoiceUpcoming EventType = "invoice.upcoming"
	// EventTypeInvoiceUpdated is sent when an invoice's details are updated
	EventTypeInvoiceUpdated EventType = "invoice.updated"
	// EventTypeInvoiceVoided is sent when an invoice is voided and can no longer be paid
	EventTypeInvoiceVoided EventType = "invoice.voided"
)

const (
	// EventObject is the type identifier for event objects
	EventObject = "event"
	// MartianPaySignature is the HTTP header name for webhook signatures
	MartianPaySignature = "Martian-Pay-Signature"
)

// EventData contains the data payload of a webhook event
type EventData struct {
	// Object is a raw mapping of the API resource contained in the event
	Object map[string]interface{} `json:"-"`
	// PreviousAttributes contains the names of updated attributes and their values prior to the event (only included in events of type `*.updated`)
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
	// Raw is the raw JSON data of the event object
	Raw json.RawMessage `json:"object"`
}

// Event represents a webhook event sent to subscribed endpoints
type Event struct {
	// ID is the unique identifier for the event
	ID string `json:"id"`
	// Object is the type identifier, always "event"
	Object string `json:"object"`
	// APIVersion is the API version used for this event
	APIVersion string `json:"api_version"`
	// Created is the Unix timestamp when the event was created
	Created int64 `json:"created"`
	// Data contains the event payload
	Data *EventData `json:"data"`
	// Livemode indicates whether this event was created in live mode or test mode
	Livemode bool `json:"livemode"`
	// PendingWebhooks is the number of webhooks that haven't been successfully delivered
	PendingWebhooks int64 `json:"pending_webhooks"`
	// Type is the event type (e.g., "invoice.created" or "charge.refunded")
	Type EventType `json:"type"`
}

// GetPayloadAndSignature extracts the payload and signature from an event
func GetPayloadAndSignature(event *Event, secret string) ([]byte, string, error) {
	if event == nil {
		return nil, "", errors.New("event is nil")
	}

	// Marshal the event to get the payload
	payload, err := json.Marshal(event)
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal event: %v", err)
	}

	// Get timestamp from event
	timestamp := time.Unix(event.Created, 0)

	// Compute signature using developer.ComputeSignature
	signature := ComputeSignature(timestamp, payload, secret)

	// Format signature as "t=timestamp,v1=signature"
	formattedSignature := fmt.Sprintf("t=%d,v1=%s", event.Created, hex.EncodeToString(signature))

	return payload, formattedSignature, nil
}

const (
	// DefaultTolerance indicates that signatures older than this will be rejected by ConstructEvent.
	DefaultTolerance time.Duration = 3000000000 * time.Second
	// signingVersion represents the version of the signature we currently use.
	signingVersion string = "v1"
)

var (
	ErrInvalidHeader    = errors.New("webhook has invalid Martian-Pay-Signature header")
	ErrNoValidSignature = errors.New("webhook had no valid signature")
	ErrNotSigned        = errors.New("webhook has no Martian-Pay-Signature header")
	ErrTooOld           = errors.New("timestamp wasn't within tolerance")
)

// signedHeader contains parsed timestamp and signatures from the webhook signature header
type signedHeader struct {
	timestamp  time.Time
	signatures [][]byte
}

// parseSignatureHeader parses the Martian-Pay-Signature header into timestamp and signatures
func parseSignatureHeader(header string) (*signedHeader, error) {
	sh := &signedHeader{}

	if header == "" {
		return sh, ErrNotSigned
	}

	// Signed header looks like "t=1495999758,v1=ABC,v1=DEF,v0=GHI"
	pairs := strings.Split(header, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) != 2 {
			return sh, ErrInvalidHeader
		}

		switch parts[0] {
		case "t":
			timestamp, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return sh, ErrInvalidHeader
			}
			sh.timestamp = time.Unix(timestamp, 0)

		case signingVersion:
			sig, err := hex.DecodeString(parts[1])
			if err != nil {
				continue // Ignore invalid signatures
			}

			sh.signatures = append(sh.signatures, sig)

		default:
			continue // Ignore unknown parts of the header
		}
	}

	if len(sh.signatures) == 0 {
		return sh, ErrNoValidSignature
	}

	return sh, nil
}

// validatePayload verifies the webhook payload signature and checks timestamp tolerance
func validatePayload(payload []byte, sigHeader string, secret string) error {
	header, err := parseSignatureHeader(sigHeader)
	if err != nil {
		return err
	}

	expectedSignature := ComputeSignature(header.timestamp, payload, secret)
	expiredTimestamp := time.Since(header.timestamp) > DefaultTolerance
	if expiredTimestamp {
		return ErrTooOld
	}

	// Check all given v1 signatures, multiple signatures will be sent temporarily in the case of a rolled signature secret
	for _, sig := range header.signatures {
		if hmac.Equal(expectedSignature, sig) {
			return nil
		}
	}

	return ErrNoValidSignature
}

func ConstructEvent(payload []byte, sigHeader string, secret string) (Event, error) {
	e := Event{}

	if err := validatePayload(payload, sigHeader, secret); err != nil {
		return e, err
	}

	if err := json.Unmarshal(payload, &e); err != nil {
		return e, fmt.Errorf("failed to parse webhook body json: %s", err.Error())
	}

	return e, nil

}
