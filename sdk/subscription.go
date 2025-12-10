package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

func (c *Client) ListSubscriptions(params *developer.ListMerchantSubscriptionsRequest) (*developer.ListSubscriptionsResponse, error) {
	var resp developer.ListSubscriptionsResponse
	err := c.sendRequestWithQuery("GET", "/v1/subscriptions", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSubscription(subscriptionID string) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CancelSubscription(subscriptionID string, params *developer.CancelMerchantSubscriptionRequest) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/cancel", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PauseSubscription(subscriptionID string, params *developer.PauseMerchantSubscriptionRequest) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/pause", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ResumeSubscription(subscriptionID string) (*developer.SubscriptionDetails, error) {
	path := fmt.Sprintf("/v1/subscriptions/%s/resume", subscriptionID)
	var resp developer.SubscriptionDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
