package main

import (
	"encoding/json"
	"strings"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Warn("failed to get request body")
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "failed to get request body",
		})
		return
	}
	req.bodyRaw = body

	event, err := developer.ConstructEvent(req.bodyRaw, req.signature, webhookSecret)
	if err != nil {
		logrus.WithError(err).Warn("Error verifying webhook signature")
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "Error verifying webhook signature",
		})
		return
	}
	logrus.WithFields(logrus.Fields{
		"event_id":      event.ID,
		"event_type":    event.Type,
		"event_created": event.Created,
	}).Info("Webhook event received")
	if strings.HasPrefix(string(event.Type), "payment_intent") {
		paymentIntent := developer.PaymentIntent{}
		err = json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshalling event data")
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		logrus.WithFields(logrus.Fields{
			"payment_intent": paymentIntent,
		}).Info("Webhook received: Payment intent")
	} else if strings.HasPrefix(string(event.Type), "refund") {
		refund := developer.Refund{}
		err = json.Unmarshal(event.Data.Raw, &refund)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshalling event data")
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		logrus.WithFields(logrus.Fields{
			"refund": refund,
		}).Info("Webhook received: Refund")
	} else if strings.HasPrefix(string(event.Type), "payout") {
		payout := developer.Payout{}
		err = json.Unmarshal(event.Data.Raw, &payout)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshalling event data")
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		logrus.WithFields(logrus.Fields{
			"payout": payout,
		}).Info("Webhook received: Payout")
	} else if strings.HasPrefix(string(event.Type), "payroll_item") {
		// Handle payroll_item events (must check before payroll to avoid prefix conflict)
		payrollItem := developer.PayrollItems{}
		err = json.Unmarshal(event.Data.Raw, &payrollItem)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshalling event data")
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		logrus.WithFields(logrus.Fields{
			"payroll_item_id": payrollItem.ID,
			"payroll_id":      payrollItem.PayrollID,
			"status":          payrollItem.Status,
			"amount":          payrollItem.Amount,
		}).Info("Webhook received: Payroll item")
	} else if strings.HasPrefix(string(event.Type), "payroll") {
		payroll := developer.Payroll{}
		err = json.Unmarshal(event.Data.Raw, &payroll)
		if err != nil {
			logrus.WithError(err).Warn("Error unmarshalling event data")
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "Error unmarshalling event data",
			})
			return
		}
		logrus.WithFields(logrus.Fields{
			"payroll_id":   payroll.ID,
			"status":       payroll.Status,
			"total_amount": payroll.TotalAmount,
			"item_count":   payroll.TotalItemNum,
		}).Info("Webhook received: Payroll")
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func main() {
	r := gin.Default()

	// Add webhook handler
	r.POST("/v1/webhook_test", webhookHandler)

	r.Run(":8080")
}
