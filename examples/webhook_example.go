// Package main provides examples for handling MartianPay webhook events.
// Webhooks allow MartianPay to send real-time notifications about payment events
// to your server. This is essential for tracking payment status changes.
//
// Webhook Events:
//   - payment_intent.* - Payment lifecycle events (created, succeeded, failed, etc.)
//   - refund.* - Refund events
//   - payout.* - Payout events
//   - payroll.* - Payroll batch events
//   - payroll_item.* - Individual payroll payment events
//
// Security:
// All webhook requests are signed with a secret key. You MUST verify the signature
// to ensure the request actually came from MartianPay and wasn't tampered with.
package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/gin-gonic/gin"
)

// webhookSecret is the secret key used to verify webhook signatures.
// In production, this should be:
// 1. Retrieved from MartianPay dashboard (Settings > Webhooks)
// 2. Stored securely (environment variable, secret manager)
// 3. Never committed to version control
const webhookSecret = "whsec_51b0f300f018596566c3dbd86ef2c8adee14673f6fde28f20be1ccc82e171f2f"

// WebhookTestRequest represents a webhook request with signature verification.
type WebhookTestRequest struct {
	signature string // HMAC signature from MartianPay-Signature header
	bodyRaw   []byte // Raw request body (needed for signature verification)
}

// webhookHandler processes incoming webhook events from MartianPay.
// This function demonstrates the complete webhook handling workflow:
// 1. Read raw request body and signature header
// 2. Verify signature to ensure authenticity
// 3. Parse event type and data
// 4. Process event based on type
// 5. Return 200 OK to acknowledge receipt
//
// Security Notes:
//   - ALWAYS verify the signature before processing
//   - Use the raw body for signature verification (not parsed JSON)
//   - Return 200 only after successful verification and processing
//   - MartianPay will retry failed webhooks (non-200 responses)
func webhookHandler(c *gin.Context) {
	var req WebhookTestRequest
	var body []byte
	var err error

	// Step 1: Extract signature from request header
	// MartianPay includes a signature in the "MartianPay-Signature" header
	req.signature = c.GetHeader(developer.MartianPaySignature)

	// Step 2: Read raw request body (needed for signature verification)
	// Important: Must use raw body, not parsed JSON
	if body, err = c.GetRawData(); err != nil {
		fmt.Printf("✗ Failed to get request body: %v\n", err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "failed to get request body",
		})
		return
	}
	req.bodyRaw = body

	// Step 3: Verify signature and construct event
	// ConstructEvent verifies the HMAC signature and parses the event
	// This ensures the request is authentic and from MartianPay
	event, err := developer.ConstructEvent(req.bodyRaw, req.signature, webhookSecret)
	if err != nil {
		fmt.Printf("✗ Error verifying webhook signature: %v\n", err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Error verifying webhook signature",
		})
		return
	}

	// Step 4: Log event details
	fmt.Printf("✓ Webhook event received:\n")
	fmt.Printf("  Event ID: %s\n", event.ID)
	fmt.Printf("  Event Type: %s\n", event.Type)
	fmt.Printf("  Created: %d\n", event.Created)

	// Step 5: Process event based on type
	// Each event type contains different data structures
	// Parse and handle accordingly
	if strings.HasPrefix(string(event.Type), "payment_intent") {
		// Handle payment intent events (payment_intent.succeeded, payment_intent.failed, etc.)
		paymentIntent := developer.PaymentIntent{}
		err = json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Printf("✗ Error unmarshalling payment intent data: %v\n", err)
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		fmt.Printf("\n  Payment Intent Details:\n")
		fmt.Printf("    ID: %s\n", paymentIntent.ID)
		fmt.Printf("    Status: %s\n", paymentIntent.Status)
		fmt.Printf("    Amount: %s %s\n", paymentIntent.Amount.Amount, paymentIntent.Amount.AssetId)
		if paymentIntent.Customer != nil {
			fmt.Printf("    Customer: %s (%s)\n", paymentIntent.Customer.ID, paymentIntent.Customer.Email)
		}
	} else if strings.HasPrefix(string(event.Type), "refund") {
		refund := developer.Refund{}
		err = json.Unmarshal(event.Data.Raw, &refund)
		if err != nil {
			fmt.Printf("✗ Error unmarshalling refund data: %v\n", err)
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		fmt.Printf("\n  Refund Details:\n")
		fmt.Printf("    ID: %s\n", refund.ID)
		fmt.Printf("    Status: %s\n", refund.Status)
		fmt.Printf("    Amount: %s %s\n", refund.Amount.Amount, refund.Amount.AssetId)
		fmt.Printf("    Payment Intent: %s\n", refund.PaymentIntent)
	} else if strings.HasPrefix(string(event.Type), "payout") {
		payout := developer.Payout{}
		err = json.Unmarshal(event.Data.Raw, &payout)
		if err != nil {
			fmt.Printf("✗ Error unmarshalling payout data: %v\n", err)
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		fmt.Printf("\n  Payout Details:\n")
		fmt.Printf("    ID: %s\n", payout.ID)
		fmt.Printf("    Status: %s\n", payout.Status)
		fmt.Printf("    Source Amount: %s %s\n", payout.SourceAmount.String(), payout.SourceCoin)
		fmt.Printf("    Receive Amount: %s %s\n", payout.ReceiveAmount.String(), payout.ReceiveCoin)
	} else if strings.HasPrefix(string(event.Type), "payroll_item") {
		// Handle payroll_item events (must check before payroll to avoid prefix conflict)
		payrollItem := developer.PayrollItems{}
		err = json.Unmarshal(event.Data.Raw, &payrollItem)
		if err != nil {
			fmt.Printf("✗ Error unmarshalling payroll item data: %v\n", err)
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		fmt.Printf("\n  Payroll Item Details:\n")
		fmt.Printf("    ID: %s\n", payrollItem.ID)
		fmt.Printf("    Payroll ID: %s\n", payrollItem.PayrollID)
		fmt.Printf("    Status: %s\n", payrollItem.Status)
		fmt.Printf("    Amount: %s\n", payrollItem.Amount)
		fmt.Printf("    Coin: %s\n", payrollItem.Coin)
		fmt.Printf("    Network: %s\n", payrollItem.Network)
	} else if strings.HasPrefix(string(event.Type), "payroll") {
		payroll := developer.Payroll{}
		err = json.Unmarshal(event.Data.Raw, &payroll)
		if err != nil {
			fmt.Printf("✗ Error unmarshalling payroll data: %v\n", err)
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		fmt.Printf("\n  Payroll Details:\n")
		fmt.Printf("    ID: %s\n", payroll.ID)
		fmt.Printf("    Status: %s\n", payroll.Status)
		fmt.Printf("    Total Amount: %s %s\n", payroll.TotalAmount, payroll.Currency)
		fmt.Printf("    Item Count: %d\n", payroll.TotalItemNum)
	}

	// Step 6: Return 200 OK to acknowledge receipt
	// MartianPay considers the webhook delivered when it receives a 2xx response
	// If you return an error (4xx, 5xx), MartianPay will retry the webhook
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

// startWebhookServer starts a local HTTP server to receive webhook events.
// This example server listens on port 8080 and processes webhook events.
//
// Setup Instructions:
// 1. Run this server (it will listen on http://localhost:8080)
// 2. Use ngrok or similar tool to expose localhost to the internet:
//    ngrok http 8080
// 3. Copy the ngrok URL (e.g., https://abc123.ngrok.io)
// 4. In MartianPay dashboard, add webhook endpoint: https://abc123.ngrok.io/v1/webhook_test
// 5. Test by creating a payment intent or triggering other events
//
// For production:
//   - Deploy to a proper server (not localhost)
//   - Use HTTPS (required by MartianPay)
//   - Implement proper error handling and logging
//   - Store webhook secret securely (environment variable)
//   - Add retry logic for failed processing
//   - Consider using a queue for async processing
func startWebhookServer() {
	fmt.Println("\n=== Webhook Event Receiver ===")
	fmt.Println("\nStarting webhook server on http://localhost:8080...")
	fmt.Println("\nServer is listening for webhook events at POST /v1/webhook_test")
	fmt.Println("See examples/README.md for the curl command to test webhook events")
	fmt.Println("\nPress Ctrl+C to stop the server\n")

	r := gin.Default()
	r.POST("/v1/webhook_test", webhookHandler)
	r.Run(":8080")
}
