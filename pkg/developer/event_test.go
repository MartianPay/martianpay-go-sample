// event_test.go contains unit tests for webhook event processing and signature verification.
package developer

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConstructEvent(t *testing.T) {
	event := Event{
		ID:         "evt_123",
		Object:     "event",
		APIVersion: "2023-10-16",
		Created:    time.Now().Unix(),
		Data: &EventData{
			Object: map[string]interface{}{
				"id":                     "pi_123",
				"object":                 "payment_intent",
				"amount":                 1000,
				"amount_received":        0,
				"canceled_at":            0,
				"cancellation_reason":    "",
				"client_secret":          "pi_123_secret_456",
				"created":                time.Now().Unix(),
				"currency":               "USD",
				"customer":               nil,
				"description":            "",
				"transactions":           []*TransactionDetails{},
				"livemode":               false,
				"metadata":               map[string]string{},
				"payment_method_type":    nil,
				"payment_method_options": nil,
				"receipt_email":          "",
				"status":                 "succeeded",
			},
		},
		Type: EventTypePaymentIntentSucceeded,
	}

	secret := "whsec_01c43aa3c8342e7cbe94c25ccf11ad709c180029fd83a217f8566751fd414327"

	payload, err := json.Marshal(event)
	assert.NoError(t, err)

	// Get valid signature
	_, signature, err := GetPayloadAndSignature(&event, secret)
	fmt.Println(signature)
	assert.NoError(t, err)

	// Test valid payload and signature
	constructedEvent, err := ConstructEvent(payload, signature, secret)
	assert.NoError(t, err)
	assert.Equal(t, event.ID, constructedEvent.ID)
	assert.Equal(t, event.Type, constructedEvent.Type)

	// Test invalid signature
	_, err = ConstructEvent(payload, "invalid_sig", secret)
	assert.Error(t, err)

	// Test expired timestamp
	event.Created = time.Now().Add(-1 * time.Hour).Unix()
	expiredPayload, _ := json.Marshal(event)
	_, expiredSig, _ := GetPayloadAndSignature(&event, secret)
	_, err = ConstructEvent(expiredPayload, expiredSig, secret)
	assert.Error(t, err)
	assert.Equal(t, ErrTooOld, err)

	// Test invalid payload
	_, err = ConstructEvent([]byte("invalid json"), signature, secret)
	assert.Error(t, err)
}
