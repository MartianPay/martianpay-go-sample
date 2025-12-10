package developer

import "github.com/dchest/uniuri"

const (
	MerchantNotificationConfigIDLength = 24
	MerchantNotificationConfigIDPrefix = "mnc_"

	MerchantNotificationLogIDLength = 24
	MerchantNotificationLogIDPrefix = "mnl_"
)

// GenerateMerchantNotificationConfigID generates a unique ID for merchant notification config
func GenerateMerchantNotificationConfigID() string {
	uniqueString := uniuri.NewLen(MerchantNotificationConfigIDLength)
	return MerchantNotificationConfigIDPrefix + uniqueString
}

// GenerateMerchantNotificationLogID generates a unique ID for merchant notification log
func GenerateMerchantNotificationLogID() string {
	uniqueString := uniuri.NewLen(MerchantNotificationLogIDLength)
	return MerchantNotificationLogIDPrefix + uniqueString
}
