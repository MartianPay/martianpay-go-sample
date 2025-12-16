// Package main provides a comprehensive set of examples demonstrating the MartianPay Go SDK.
// This package includes interactive examples for all major API features including payment intents,
// customers, products, subscriptions, payouts, refunds, and more.
package main

// apiKey is the default API key used for MartianPay API authentication.
// This can be overridden when the program starts by entering a custom API key at the prompt.
// To obtain your API key:
// 1. Log in to the MartianPay dashboard at https://dashboard.martianpay.com
// 2. Navigate to Settings > API Keys
// 3. Create a new API key or use an existing one
const apiKey = "your_api_key_here" // Default API key (can be overridden at startup)

// currentAPIKey stores the runtime API key that will be used for all API calls.
// This variable is set during program startup and can be either the default apiKey
// or a custom key entered by the user.
var currentAPIKey string // Runtime API key (set during startup)
