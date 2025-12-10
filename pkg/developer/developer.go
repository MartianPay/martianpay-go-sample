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
	MartianPayApiVersion = "2025-01-22"

	WebhookEndpointIDLength  = 24
	WebhookEndpointIDPrefix  = "wh_"
	WebhookEndpointPrefixKey = "whsec_"
	WebhookEndpointObject    = "webhook_endpoint"

	ApiKeyIDLength = 24
	ApiKeyIDPrefix = "ak_"
	ApiKeyObject   = "api_key"

	ApiKeyPrefixPublicTest = "pk_test_"
	ApiKeyPrefixSecretTest = "sk_test_"
	ApiKeyPrefixPublicLive = "pk_live_"
	ApiKeyPrefixSecretLive = "sk_live_"
	ApiKeyRandomPartLength = 96
)

const (
	DeveloperKeyTypePublic string = "public"
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

func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%d", t.Unix())))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

// GenerateApiKeyPair generates a pair of public and secret API keys
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
	DeveloperStatusActive   = "active"
	DeveloperStatusInactive = "inactive"
	DeveloperStatusDeleted  = "deleted"
	DeveloperStatusDeleting = "deleting"
	DeveloperStatusUnknown  = "unknown"
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

func GenerateApiKeyID() string {
	uniqueString := uniuri.NewLen(ApiKeyIDLength)
	return ApiKeyIDPrefix + uniqueString
}
