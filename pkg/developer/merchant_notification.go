// merchant_notification.go contains types and utilities for merchant notification management.
// It provides ID generation for notification configurations and logs.
package developer

import "github.com/dchest/uniuri"

const (
	// MerchantNotificationConfigIDLength is the length of the notification config ID suffix
	MerchantNotificationConfigIDLength = 24
	// MerchantNotificationConfigIDPrefix is the prefix for notification config IDs
	MerchantNotificationConfigIDPrefix = "mnc_"

	// MerchantNotificationLogIDLength is the length of the notification log ID suffix
	MerchantNotificationLogIDLength = 24
	// MerchantNotificationLogIDPrefix is the prefix for notification log IDs
	MerchantNotificationLogIDPrefix = "mnl_"
)

// GenerateMerchantNotificationConfigID generates a unique ID for merchant notification config.
// The generated ID has the format 'mnc_' followed by a 24-character random string.
// This ID is used to identify notification configuration settings for a merchant.
func GenerateMerchantNotificationConfigID() string {
	uniqueString := uniuri.NewLen(MerchantNotificationConfigIDLength)
	return MerchantNotificationConfigIDPrefix + uniqueString
}

// GenerateMerchantNotificationLogID generates a unique ID for merchant notification log.
// The generated ID has the format 'mnl_' followed by a 24-character random string.
// This ID is used to track individual notification delivery attempts and their results.
func GenerateMerchantNotificationLogID() string {
	uniqueString := uniuri.NewLen(MerchantNotificationLogIDLength)
	return MerchantNotificationLogIDPrefix + uniqueString
}
