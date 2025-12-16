// Package developer contains data structures and types used by the MartianPay SDK.
// It provides comprehensive types for managing payments, subscriptions, products,
// customers, and other e-commerce operations through the MartianPay API.
//
// The package includes:
//   - API authentication and webhook signature verification
//   - Payment intent and charge management
//   - Product catalog and variant management
//   - Subscription and selling plan configuration
//   - Customer management and authentication
//   - Payout and payroll operations
//   - Refund and settlement tracking
//
// For API documentation, visit: https://docs.martianpay.com
package developer

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dchest/uniuri"
)

const (
	// MartianPayApiVersion is the current API version used for requests and webhooks
	MartianPayApiVersion = "2025-01-22"

	// WebhookEndpointIDLength is the length of the webhook endpoint ID suffix (excluding prefix)
	WebhookEndpointIDLength = 24
	// WebhookEndpointIDPrefix is the prefix for webhook endpoint IDs
	WebhookEndpointIDPrefix = "wh_"
	// WebhookEndpointPrefixKey is the prefix for webhook endpoint secret keys
	WebhookEndpointPrefixKey = "whsec_"
	// WebhookEndpointObject is the object type identifier for webhook endpoints
	WebhookEndpointObject = "webhook_endpoint"

	// ApiKeyIDLength is the length of the API key ID suffix (excluding prefix)
	ApiKeyIDLength = 24
	// ApiKeyIDPrefix is the prefix for API key IDs
	ApiKeyIDPrefix = "ak_"
	// ApiKeyObject is the object type identifier for API keys
	ApiKeyObject = "api_key"

	// ApiKeyPrefixPublicTest is the prefix for publishable test API keys
	ApiKeyPrefixPublicTest = "pk_test_"
	// ApiKeyPrefixSecretTest is the prefix for secret test API keys
	ApiKeyPrefixSecretTest = "sk_test_"
	// ApiKeyPrefixPublicLive is the prefix for publishable live API keys
	ApiKeyPrefixPublicLive = "pk_live_"
	// ApiKeyPrefixSecretLive is the prefix for secret live API keys
	ApiKeyPrefixSecretLive = "sk_live_"
	// ApiKeyRandomPartLength is the length of the random part of API keys (in characters)
	ApiKeyRandomPartLength = 96
)

const (
	// DeveloperKeyTypePublic indicates a publishable/public API key that can be safely exposed in client-side code
	DeveloperKeyTypePublic string = "public"
	// DeveloperKeyTypeSecret indicates a secret API key that should only be used server-side
	DeveloperKeyTypeSecret string = "secret"
)

// GenerateHMACKey generates a random HMAC SHA256 key and returns it as a hex string
func GenerateHMACKey() (string, error) {
	// Generate 32 random bytes (256 bits) for the HMAC key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	// Convert to hex string
	return WebhookEndpointPrefixKey + hex.EncodeToString(key), nil
}

// ComputeSignature computes an HMAC SHA256 signature for webhook payload verification
// It combines the timestamp, a period separator, and the payload to generate the signature
func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%d", t.Unix())))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

// GenerateApiKey generates an API key of the specified type (public or secret) for test or live mode
func GenerateApiKey(keyType string, isLive bool) (key string, err error) {
	// Generate 96 random bytes for the key
	keyRandomBytes := make([]byte, ApiKeyRandomPartLength/2)
	if _, err := rand.Read(keyRandomBytes); err != nil {
		return "", err
	}

	// Convert to hex string
	keyRandomStr := hex.EncodeToString(keyRandomBytes)

	// Add appropriate prefixes based on mode
	if isLive {
		if keyType == DeveloperKeyTypePublic {
			key = ApiKeyPrefixPublicLive + keyRandomStr
		} else if keyType == DeveloperKeyTypeSecret {
			key = ApiKeyPrefixSecretLive + keyRandomStr
		}
	} else {
		if keyType == DeveloperKeyTypePublic {
			key = ApiKeyPrefixPublicTest + keyRandomStr
		} else if keyType == DeveloperKeyTypeSecret {
			key = ApiKeyPrefixSecretTest + keyRandomStr
		}
	}

	return key, nil
}

// EncryptData encrypts data using AES-256-CBC encryption with PKCS7 padding
// Returns the encrypted data as a base64-encoded string that includes the IV prefix
func EncryptData(data string, encryptionKey string) (string, error) {
	key := []byte(encryptionKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Use fixed IV instead of random
	iv := bytes.Repeat([]byte{0x00}, aes.BlockSize)

	// Pad data to be multiple of block size
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	data = data + string(padtext)

	// Encrypt
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, []byte(data))

	// Combine IV and ciphertext
	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

// DecryptData decrypts AES-256-CBC encrypted data with PKCS7 padding
// Takes a base64-encoded encrypted string and returns the original plaintext
func DecryptData(encryptedData string, encryptionKey string) (string, error) {
	key := []byte(encryptionKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Decode base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract IV
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Decrypt
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad
	padding := int(ciphertext[len(ciphertext)-1])
	if padding < 1 || padding > aes.BlockSize {
		return "", fmt.Errorf("invalid padding")
	}
	return string(ciphertext[:len(ciphertext)-padding]), nil
}

const (
	// DeveloperStatusActive indicates the resource is active and operational
	DeveloperStatusActive = "active"
	// DeveloperStatusInactive indicates the resource is inactive and not currently operational
	DeveloperStatusInactive = "inactive"
	// DeveloperStatusDeleted indicates the resource has been deleted
	DeveloperStatusDeleted = "deleted"
	// DeveloperStatusDeleting indicates the resource is in the process of being deleted
	DeveloperStatusDeleting = "deleting"
	// DeveloperStatusUnknown indicates the resource status is unknown
	DeveloperStatusUnknown = "unknown"
)

// WebhookEndpointParams contains parameters for creating or updating a webhook endpoint
type WebhookEndpointParams struct {
	// URL is the webhook endpoint URL
	URL *string `json:"url"`
	// Description is a description of the webhook endpoint
	Description *string `json:"description"`
	// EnabledEvents is the list of events to enable for this endpoint (use ['*'] to enable all events)
	EnabledEvents []string `json:"enabled_events"`
}

// WebhookEndpoint represents a webhook endpoint for receiving event notifications
type WebhookEndpoint struct {
	// ID is the unique identifier for the webhook endpoint
	ID string `json:"id"`
	// Object is the type identifier, always "webhook_endpoint"
	Object string `json:"object"`
	// URL is the URL where webhook events are sent
	URL string `json:"url"`
	// APIVersion is the API version used for webhook events
	APIVersion string `json:"api_version"`
	// Created is the Unix timestamp when the webhook was created
	Created int64 `json:"created"`
	// Deleted indicates whether the webhook has been deleted
	Deleted bool `json:"deleted"`
	// Description is a description of the webhook endpoint
	Description string `json:"description"`
	// EnabledEvents is the list of event types this endpoint receives
	EnabledEvents []string `json:"enabled_events"`
	// Livemode indicates if this is a live mode webhook (true) or test mode (false)
	Livemode bool `json:"livemode"`
	// Metadata is a set of key-value pairs for storing additional information
	Metadata map[string]string `json:"metadata"`
	// Secret is the secret key for verifying webhook signatures
	Secret string `json:"secret"`
	// Status is the current status of the webhook endpoint
	Status string `json:"status"`
}

// GenerateWebhookEndpointID generates a new unique webhook endpoint identifier with the 'wh_' prefix
func GenerateWebhookEndpointID() string {
	uniqueString := uniuri.NewLen(WebhookEndpointIDLength)
	return WebhookEndpointIDPrefix + uniqueString
}

// ApiKeyParams contains parameters for creating or updating an API key
type ApiKeyParams struct {
	// Name is the name of the API key
	Name *string `json:"name"`
	// Description is a description of the API key's purpose
	Description *string `json:"description"`
	// KeyType is the type of API key ("public" or "secret")
	KeyType *string `json:"key_type"`
}

// ApiKey represents an API key for authentication
type ApiKey struct {
	// ID is the unique identifier for the API key
	ID string `json:"id"`
	// Object is the type identifier, always "api_key"
	Object string `json:"object"`
	// Key is the actual API key string
	Key string `json:"key"`
	// KeyType is the type of key ("public" or "secret")
	KeyType string `json:"key_type"`
	// Name is the name of the API key
	Name string `json:"name"`
	// Description is a description of the API key's purpose
	Description string `json:"description"`
	// Created is the Unix timestamp when the key was created
	Created int64 `json:"created"`
	// LastAccessedAt is the Unix timestamp when the key was last used
	LastAccessedAt int64 `json:"last_accessed_at"`
	// Status is the current status of the API key
	Status string `json:"status"`
	// ExpiredAt is the Unix timestamp when the key expires
	ExpiredAt int64 `json:"expired_at"`
}

// GenerateApiKeyID generates a new unique API key identifier with the 'ak_' prefix
func GenerateApiKeyID() string {
	uniqueString := uniuri.NewLen(ApiKeyIDLength)
	return ApiKeyIDPrefix + uniqueString
}
