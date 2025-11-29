package martianpay

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateMerchantAddress(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Create a merchant address
	// Note: Replace with your actual blockchain address
	req := MerchantAddressCreateRequest{
		Network: "Ethereum Sepolia",                          // Test network
		Address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", // Your Ethereum address
	}

	resp, err := client.CreateMerchantAddress(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.ID)
		assert.Equal(t, req.Network, resp.Network)
		assert.Equal(t, req.Address, resp.Address)
		assert.NotEmpty(t, resp.Status) // Should be "created" initially

		logrus.WithFields(logrus.Fields{
			"address_id": resp.ID,
			"network":    resp.Network,
			"address":    resp.Address,
			"status":     resp.Status,
		}).Info("Successfully created merchant address")

		// If verification is returned, log it
		if resp.Verification != nil {
			logrus.WithFields(logrus.Fields{
				"verification_id": resp.Verification.ID,
				"status":          resp.Verification.Status,
				"asset_id":        resp.Verification.AssetID,
			}).Info("Address verification details")
		}
	}
}

func TestGetMerchantAddress(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Get a merchant address by ID
	// Replace with an actual address ID from your account
	addressID := "ma_xxxxxxxxxxxxxxxxxxxxxx"

	resp, err := client.GetMerchantAddress(addressID)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		assert.Equal(t, addressID, resp.ID)

		logrus.WithFields(logrus.Fields{
			"address_id": resp.ID,
			"network":    resp.Network,
			"address":    resp.Address,
			"status":     resp.Status,
			"alias":      resp.Alias,
			"created_at": resp.CreatedAt,
		}).Info("Successfully retrieved merchant address")

		if resp.Verification != nil {
			logrus.WithFields(logrus.Fields{
				"verification_status": resp.Verification.Status,
				"tried_times":         resp.Verification.TriedTimes,
				"aml_status":          resp.Verification.AmlStatus,
			}).Info("Verification details")
		}
	}
}

func TestUpdateMerchantAddress(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Update a merchant address
	// Replace with an actual address ID from your account
	addressID := "ma_xxxxxxxxxxxxxxxxxxxxxx"
	alias := "My Main Wallet"

	req := MerchantAddressUpdateRequest{
		ID:    addressID,
		Alias: &alias,
	}

	resp, err := client.UpdateMerchantAddress(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		assert.Equal(t, addressID, resp.ID)
		assert.Equal(t, alias, resp.Alias)

		logrus.WithFields(logrus.Fields{
			"address_id": resp.ID,
			"alias":      resp.Alias,
			"network":    resp.Network,
			"address":    resp.Address,
		}).Info("Successfully updated merchant address")
	}
}

func TestVerifyMerchantAddress(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Verify a merchant address
	// Replace with an actual address ID from your account
	// You need to send a small test transaction to this address first
	addressID := "ma_xxxxxxxxxxxxxxxxxxxxxx"

	req := MerchantAddressVerifyRequest{
		ID:     addressID,
		Amount: "0.01", // The amount you sent to verify ownership
	}

	resp, err := client.VerifyMerchantAddress(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)
		assert.Equal(t, addressID, resp.ID)

		logrus.WithFields(logrus.Fields{
			"address_id": resp.ID,
			"status":     resp.Status,
		}).Info("Successfully initiated address verification")

		if resp.Verification != nil {
			logrus.WithFields(logrus.Fields{
				"verification_id":     resp.Verification.ID,
				"verification_status": resp.Verification.Status,
				"asset_id":            resp.Verification.AssetID,
				"tried_times":         resp.Verification.TriedTimes,
			}).Info("Verification details")
		}
	}
}

func TestListMerchantAddresses(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// List all merchant addresses
	req := MerchantAddressListRequest{
		Page:     0,
		PageSize: 10,
		// Network: Optional - filter by specific network
	}

	resp, err := client.ListMerchantAddresses(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)

		logrus.WithFields(logrus.Fields{
			"total":     resp.Total,
			"page":      resp.Page,
			"page_size": resp.PageSize,
			"count":     len(resp.MerchantAddresses),
		}).Info("Successfully listed merchant addresses")

		// Log each address
		for i, addr := range resp.MerchantAddresses {
			logrus.WithFields(logrus.Fields{
				"index":      i,
				"address_id": addr.ID,
				"network":    addr.Network,
				"address":    addr.Address,
				"status":     addr.Status,
				"alias":      addr.Alias,
			}).Info("Merchant address details")
		}
	}
}

func TestListMerchantAddressesByNetwork(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// List merchant addresses filtered by network
	network := "Ethereum Sepolia"
	req := MerchantAddressListRequest{
		Page:     0,
		PageSize: 10,
		Network:  &network,
	}

	resp, err := client.ListMerchantAddresses(req)

	// Verify response
	if assert.NoError(t, err) {
		assert.NotNil(t, resp)

		logrus.WithFields(logrus.Fields{
			"network": network,
			"total":   resp.Total,
			"count":   len(resp.MerchantAddresses),
		}).Info("Successfully listed merchant addresses for network")

		// Verify all addresses are from the specified network
		for _, addr := range resp.MerchantAddresses {
			assert.Equal(t, network, addr.Network)
		}
	}
}

func TestDeleteMerchantAddress(t *testing.T) {
	// Initialize the client with your API key
	client := NewClient(apiKey)

	// Delete a merchant address
	// Replace with an actual address ID from your account
	// WARNING: This will permanently delete the address
	addressID := "ma_xxxxxxxxxxxxxxxxxxxxxx"

	err := client.DeleteMerchantAddress(addressID)

	// Verify response
	if assert.NoError(t, err) {
		logrus.WithField("address_id", addressID).Info("Successfully deleted merchant address")

		// Verify the address is deleted by trying to get it
		_, getErr := client.GetMerchantAddress(addressID)
		assert.Error(t, getErr) // Should return an error since address is deleted
	}
}

func TestCreateAndVerifyMerchantAddress(t *testing.T) {
	// This test demonstrates the complete workflow:
	// 1. Create an address
	// 2. Send a test transaction to it (manual step)
	// 3. Verify ownership with the transaction amount

	client := NewClient(apiKey)

	// Step 1: Create a merchant address
	createReq := MerchantAddressCreateRequest{
		Network: "Ethereum Sepolia",
		Address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
	}

	createResp, err := client.CreateMerchantAddress(createReq)
	assert.NoError(t, err)
	assert.NotNil(t, createResp)

	addressID := createResp.ID
	logrus.WithFields(logrus.Fields{
		"address_id": addressID,
		"network":    createResp.Network,
		"address":    createResp.Address,
	}).Info("Step 1: Created merchant address")

	// Step 2: Manual step - send a small test transaction
	logrus.Info("Step 2: Send a test transaction to the address to verify ownership")
	logrus.WithFields(logrus.Fields{
		"address": createResp.Address,
		"network": createResp.Network,
		"note":    "Send a small amount (e.g., 0.01 USDC) and note the exact amount",
	}).Warn("MANUAL ACTION REQUIRED")

	// For testing purposes, we'll skip the actual verification
	// In a real scenario, after sending the transaction, you would call:
	/*
		verifyReq := MerchantAddressVerifyRequest{
			ID:     addressID,
			Amount: "0.01", // The exact amount you sent
		}
		verifyResp, err := client.VerifyMerchantAddress(verifyReq)
		assert.NoError(t, err)
		logrus.WithField("status", verifyResp.Status).Info("Step 3: Verification initiated")
	*/

	t.Skip("Skipping verification step - requires manual transaction")
}
