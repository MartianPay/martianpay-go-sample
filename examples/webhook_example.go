package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/gin-gonic/gin"
)

const webhookSecret = "whsec_51b0f300f018596566c3dbd86ef2c8adee14673f6fde28f20be1ccc82e171f2f"

type WebhookTestRequest struct {
	signature string
	bodyRaw   []byte
}

func webhookHandler(c *gin.Context) {
	var req WebhookTestRequest
	var body []byte
	var err error

	req.signature = c.GetHeader(developer.MartianPaySignature)
	if body, err = c.GetRawData(); err != nil {
		fmt.Printf("✗ Failed to get request body: %v\n", err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "failed to get request body",
		})
		return
	}
	req.bodyRaw = body

	event, err := developer.ConstructEvent(req.bodyRaw, req.signature, webhookSecret)
	if err != nil {
		fmt.Printf("✗ Error verifying webhook signature: %v\n", err)
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Error verifying webhook signature",
		})
		return
	}

	fmt.Printf("✓ Webhook event received:\n")
	fmt.Printf("  Event ID: %s\n", event.ID)
	fmt.Printf("  Event Type: %s\n", event.Type)
	fmt.Printf("  Created: %d\n", event.Created)

	if strings.HasPrefix(string(event.Type), "payment_intent") {
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

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

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
