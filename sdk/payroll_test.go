package martianpay

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateDirectPayroll(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Step 1: Query assets to get crypto-only assets (payroll only supports crypto)
	assetsResp, err := client.ListAssets()
	assert.NoError(t, err)
	assert.NotNil(t, assetsResp)

	// Filter crypto assets that are payable
	var cryptoAssets []*developer.Asset
	for _, asset := range assetsResp.Assets {
		if !asset.IsFiat && asset.Payable && asset.CryptoAssetParams != nil {
			cryptoAssets = append(cryptoAssets, asset)
		}
	}

	logrus.WithField("crypto_asset_count", len(cryptoAssets)).Info("Found crypto assets")

	// Step 2: Query merchant balance to find available coins
	balance, err := client.GetBalance()
	assert.NoError(t, err)
	assert.NotNil(t, balance)

	logrus.WithFields(logrus.Fields{
		"currency":          balance.Currency,
		"available_balance": balance.AvailableBalance,
		"total_balance":     balance.TotalBalance,
	}).Info("Merchant balance summary")

	// Step 3: Match crypto assets with available balance >= 0.1
	var selectedCoin, selectedNetwork string

	for _, asset := range cryptoAssets {
		// Match asset ID with balance details
		for _, detail := range balance.BalanceDetails {
			if detail.Currency == asset.Id && detail.AvailableBalance != "0" && detail.AvailableBalance != "" {
				// Parse available balance to check if >= 0.1
				availableFloat, err := strconv.ParseFloat(detail.AvailableBalance, 64)
				if err != nil {
					continue
				}
				if availableFloat >= 0.1 {
					// Use the asset's coin and network from CryptoAssetParams
					selectedCoin = asset.Coin
					selectedNetwork = asset.CryptoAssetParams.Network
					logrus.WithFields(logrus.Fields{
						"asset_id":          asset.Id,
						"coin":              selectedCoin,
						"network":           selectedNetwork,
						"available_balance": detail.AvailableBalance,
					}).Info("Selected crypto asset with sufficient balance")
					break
				}
			}
		}
		if selectedCoin != "" {
			break
		}
	}

	// Skip test if no sufficient balance available
	if selectedCoin == "" {
		t.Skip("Skipping test - no crypto balance >= 0.1 found")
		return
	}

	// Step 4: Generate unique external IDs using timestamp to avoid duplication
	timestamp := time.Now().UnixNano()
	externalID := fmt.Sprintf("ORDER-%d", timestamp)
	itemExternalID1 := fmt.Sprintf("ITEM-%d-001", timestamp)
	itemExternalID2 := fmt.Sprintf("ITEM-%d-002", timestamp)

	// Step 5: Create request with auto-approval using selected coin/network
	req := PayrollDirectCreateRequest{
		ExternalID:  externalID,
		AutoApprove: true,
		Items: []PayrollDirectItem{
			{
				ExternalID:    itemExternalID1,
				Name:          "John Doe",
				Email:         "john@example.com",
				Phone:         "+1234567890",
				Coin:          selectedCoin,
				Network:       selectedNetwork,
				Address:       "TN9RRaXkCFtTXRso2GdTZxSxxwufzxLQPP", // TODO: Use dynamic address
				Amount:        "0.1",
				PaymentMethod: "normal",
			},
			{
				ExternalID:    itemExternalID2,
				Coin:          selectedCoin,
				Network:       selectedNetwork,
				Address:       "TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb", // TODO: Use dynamic address
				Amount:        "0.1",
				PaymentMethod: "normal",
			},
		},
	}

	// Call API
	response, err := client.CreateDirectPayroll(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.NotEmpty(t, response.Payroll.ID)
		assert.Equal(t, req.ExternalID, response.Payroll.ExternalID)
		assert.Equal(t, int64(2), response.Payroll.TotalItemNum)

		logrus.WithFields(logrus.Fields{
			"payroll_id":        response.Payroll.ID,
			"external_id":       response.Payroll.ExternalID,
			"status":            response.Payroll.Status,
			"approval_status":   response.Payroll.ApprovalStatus,
			"total_items":       response.Payroll.TotalItemNum,
			"total_amount":      response.Payroll.TotalAmount,
			"total_service_fee": response.Payroll.TotalServiceFee,
		}).Info("Successfully created direct payroll")

		// Log individual items
		for _, item := range response.Items {
			logrus.WithFields(logrus.Fields{
				"item_id":        item.ID,
				"external_id":    item.ExternalID,
				"name":           item.Name,
				"amount":         item.Amount,
				"coin":           item.Coin,
				"network":        item.Network,
				"status":         item.Status,
				"payment_method": item.PaymentMethod,
			}).Info("Payroll item details")
		}
	}
}

func TestGetPayroll(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create request
	req := PayrollGetReq{
		ID: "payroll_I6mzarIW0qZ4itwGg7xCa9EA",
	}

	// Call API
	response, err := client.GetPayroll(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)
		assert.Equal(t, req.ID, response.Payroll.ID)

		logrus.WithFields(logrus.Fields{
			"id":              response.Payroll.ID,
			"external_id":     response.Payroll.ExternalID,
			"status":          response.Payroll.Status,
			"approval_status": response.Payroll.ApprovalStatus,
			"total_items":     response.Payroll.TotalItemNum,
			"total_amount":    response.Payroll.TotalAmount,
		}).Info("Successfully retrieved payroll")
	}
}

func TestListPayrolls(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	externalID := "ORDER-2024-001"

	// Create request
	req := PayrollListReq{
		Page:       0,
		PageSize:   10,
		ExternalID: &externalID, // Optional: filter by external ID
	}

	// Call API
	response, err := client.ListPayrolls(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)

		logrus.WithFields(logrus.Fields{
			"total_records": response.Total,
		}).Info("Successfully retrieved payrolls")

		// Log individual payrolls
		for _, payroll := range response.Payrolls {
			logrus.WithFields(logrus.Fields{
				"id":              payroll.ID,
				"external_id":     payroll.ExternalID,
				"status":          payroll.Status,
				"approval_status": payroll.ApprovalStatus,
				"total_items":     payroll.TotalItemNum,
				"total_amount":    payroll.TotalAmount,
			}).Info("Payroll details")
		}
	}
}

func TestListPayrollItems(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	externalID := "ORDER-2024-001"

	// Create request
	req := PayrollItemsListReq{
		Page:       0,
		PageSize:   10,
		ExternalID: &externalID, // Optional: filter by payroll external ID
	}

	// Call API
	response, err := client.ListPayrollItems(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, response)

		logrus.WithFields(logrus.Fields{
			"total_records": response.Total,
		}).Info("Successfully retrieved payroll items")

		// Log individual items
		for _, item := range response.PayrollItems {
			logrus.WithFields(logrus.Fields{
				"id":             item.ID,
				"external_id":    item.ExternalID,
				"payroll_id":     item.PayrollID,
				"name":           item.Name,
				"amount":         item.Amount,
				"coin":           item.Coin,
				"network":        item.Network,
				"status":         item.Status,
				"payment_method": item.PaymentMethod,
			}).Info("Payroll item details")
		}
	}
}
