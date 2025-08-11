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

// Description of the event (for example, `invoice.created` or `charge.refunded`).
type EventType string

// List of values that EventType can take
const (
	EventTypePaymentIntentCreated       EventType = "payment_intent.created"
	EventTypePaymentIntentSucceeded     EventType = "payment_intent.succeeded"
	EventTypePaymentIntentPaymentFailed EventType = "payment_intent.payment_failed"
	EventTypePaymentIntentProcessing    EventType = "payment_intent.processing"
	EventTypePaymentIntentPartiallyPaid EventType = "payment_intent.partially_paid"
	EventTypePaymentIntentCanceled      EventType = "payment_intent.canceled"

	EventTypeRefundCreated   EventType = "refund.created"
	EventTypeRefundSucceeded EventType = "refund.succeeded"
	EventTypeRefundUpdated   EventType = "refund.updated"
	EventTypeRefundFailed    EventType = "refund.failed"

	EventTypePayoutCreated   EventType = "payout.created"
	EventTypePayoutSucceeded EventType = "payout.succeeded"
	EventTypePayoutUpdated   EventType = "payout.updated"
	EventTypePayoutFailed    EventType = "payout.failed"

	// Payroll events
	EventTypePayrollCreated   EventType = "payroll.created"
	EventTypePayrollApproved  EventType = "payroll.approved"
	EventTypePayrollRejected  EventType = "payroll.rejected"
	EventTypePayrollCanceled  EventType = "payroll.canceled"
	EventTypePayrollExecuting EventType = "payroll.executing"
	EventTypePayrollCompleted EventType = "payroll.completed"
	EventTypePayrollFailed    EventType = "payroll.failed"

	// Payroll item events
	EventTypePayrollItemProcessing          EventType = "payroll_item.processing"
	EventTypePayrollItemSucceeded           EventType = "payroll_item.succeeded"
	EventTypePayrollItemFailed              EventType = "payroll_item.failed"
	EventTypePayrollItemAddressVerification EventType = "payroll_item.address_verification_sent"
	EventTypePayrollItemAddressVerified     EventType = "payroll_item.address_verified"
)

const (
	EventObject         = "event"
	MartianPaySignature = "Martian-Pay-Signature"
)

type EventData struct {
	// Object is a raw mapping of the API resource contained in the event.
	// Although marked with json:"-", it's still populated independently by
	// a custom UnmarshalJSON implementation.
	// Object containing the API resource relevant to the event.
	Object map[string]interface{} `json:"-"`
	// Object containing the names of the updated attributes and their values prior to the event (only included in events of type `*.updated`). If an array attribute has any updated elements, this object contains the entire array.
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
	Raw                json.RawMessage        `json:"object"`
}

type Event struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	APIVersion string `json:"api_version"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created  int64      `json:"created"`
	Data     *EventData `json:"data"`
	Livemode bool       `json:"livemode"`
	// Number of webhooks that haven't been successfully delivered (for example, to return a 20x response) to the URLs you specify.
	PendingWebhooks int64 `json:"pending_webhooks"`
	// Information on the API request that triggers the event.
	// Request *EventRequest `json:"request"`
	// Description of the event (for example, `invoice.created` or `charge.refunded`).
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

type signedHeader struct {
	timestamp  time.Time
	signatures [][]byte
}

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
