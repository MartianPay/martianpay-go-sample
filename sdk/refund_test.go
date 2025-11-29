package martianpay

import (
	"testing"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateRefund(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	reason := string(developer.PaymentIntentCancellationReasonRequestedByCustomer)
	paymentIntentId := "pi_m3cL1pixO1D3c8A5iC18iGZo"
	req := RefundCreateRequest{
		RefundParams: developer.RefundParams{
			PaymentIntent: &paymentIntentId,
			Amount:        "1",
			Reason:        &reason,
		},
	}

	// Call API
	response, err := client.CreateRefund(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.NotEmpty(t, response.Refunds)
		logrus.WithFields(logrus.Fields{
			"refund_id":      response.Refunds[0].ID,
			"amount":         response.Refunds[0].Amount,
			"payment_intent": response.Refunds[0].PaymentIntent,
		}).Info("Successfully created refund")
	}
}

func TestGetRefund(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := RefundGetRequest{
		ID: "rf_ZTqX7HuMtGxaJ1KnizURLU9m", // Replace with a valid refund ID
	}

	// Call API
	response, err := client.GetRefund(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.ID, response.ID)
		logrus.WithFields(logrus.Fields{
			"id":     response.ID,
			"amount": response.Amount,
			"status": response.Status,
		}).Info("Successfully retrieved refund")
	}
}

func TestListRefunds(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := RefundListRequest{
		Page:     0,
		PageSize: 10,
	}

	// Call API
	response, err := client.ListRefunds(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.Page, response.Page)
		assert.Equal(t, req.PageSize, response.PageSize)

		logrus.WithFields(logrus.Fields{
			"total_records": response.Total,
			"page":          response.Page,
			"page_size":     response.PageSize,
		}).Info("Successfully retrieved refunds")

		for _, refund := range response.Refunds {
			logrus.WithFields(logrus.Fields{
				"id":             refund.ID,
				"amount":         refund.Amount,
				"status":         refund.Status,
				"payment_intent": refund.PaymentIntent,
			}).Info("Refund details")
		}
	}
}
