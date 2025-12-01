# MartianPay SDK Testing Guide

This directory contains the MartianPay Go SDK implementation and unit tests.

## Running Tests

Before testing, modify `const apiKey = "your_api_key_here"` in `sdk/common_test.go`

Also update the email and ID in each test case as needed.

### Payment Intent Tests

```sh
# Create and update payment intent (crypto)
go test -count=1 -v -run TestCreateAndUpdatePaymentIntent sdk/*.go

# List payment intents
go test -count=1 -v -run TestListPaymentIntents sdk/*.go

# Get specific payment intent
go test -count=1 -v -run TestGetPaymentIntent sdk/*.go

# Cancel payment intent
go test -count=1 -v -run TestCancelPaymentIntent sdk/*.go

# List payment intents with permanent deposit filter
go test -count=1 -v -run TestListPaymentIntentsWithPermanentDeposit sdk/*.go
```

### Customer Tests

```sh
# Create and update customer
go test -count=1 -v -run TestCreateAndUpdateCustomer sdk/*.go

# List customers
go test -count=1 -v -run TestListCustomers sdk/*.go

# Get specific customer
go test -count=1 -v -run TestGetCustomer sdk/*.go

# Delete customer
go test -count=1 -v -run TestDeleteCustomer sdk/*.go
```

### Payment Method Tests

```sh
# List customer's saved payment methods
go test -count=1 -v -run TestListCustomerPaymentMethods sdk/*.go
```

### Fiat/Card Payment Tests

```sh
# Fiat payment with new card
go test -count=1 -v -run TestFiatPaymentWithNewCard sdk/*.go

# Fiat payment with saved card
go test -count=1 -v -run TestFiatPaymentWithSavedCard sdk/*.go
```

### Refund Tests

```sh
# Create refund
go test -count=1 -v -run TestCreateRefund sdk/*.go

# Get specific refund
go test -count=1 -v -run TestGetRefund sdk/*.go

# List refunds
go test -count=1 -v -run TestListRefunds sdk/*.go
```

### Payroll Tests

```sh
# Create direct payroll with auto-approval
go test -count=1 -v -run TestCreateDirectPayroll sdk/*.go

# Get specific payroll
go test -count=1 -v -run TestGetPayroll sdk/*.go

# List payrolls
go test -count=1 -v -run TestListPayrolls sdk/*.go

# List payroll items
go test -count=1 -v -run TestListPayrollItems sdk/*.go
```

### Merchant Address (Wallet) Tests

```sh
# Create a merchant address
go test -count=1 -v -run TestCreateMerchantAddress sdk/*.go

# Get specific merchant address
go test -count=1 -v -run TestGetMerchantAddress sdk/*.go

# Update merchant address (set alias)
go test -count=1 -v -run TestUpdateMerchantAddress sdk/*.go

# Verify merchant address ownership
go test -count=1 -v -run TestVerifyMerchantAddress sdk/*.go

# List all merchant addresses
go test -count=1 -v -run TestListMerchantAddresses sdk/*.go

# List merchant addresses by network
go test -count=1 -v -run TestListMerchantAddressesByNetwork sdk/*.go

# Delete a merchant address
go test -count=1 -v -run TestDeleteMerchantAddress sdk/*.go

# Complete workflow: create and verify address
go test -count=1 -v -run TestCreateAndVerifyMerchantAddress sdk/*.go
```

## Notes

- All test commands should be run from the project root directory
- Tests use the API key defined in `sdk/common_test.go`
- Make sure to update test data (emails, IDs, etc.) to match your test environment
- Tests make real API calls to the MartianPay API

## Alternative: Interactive Examples

For a more user-friendly way to explore the SDK, check out the [interactive examples](../examples/README.md) which provide a menu-driven interface to run all SDK examples.
