// Package martianpay provides SDK methods for managing invoices.
// Invoices are detailed billing documents sent to customers requesting payment for goods or services.
package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListInvoices retrieves a paginated list of merchant invoices.
// Can be filtered by status, customer, due date, and other criteria.
//
// Parameters:
//   - params: Query parameters including pagination and filters (status, customer, date range, etc.)
//
// Returns:
//   - *developer.ListInvoicesResponse: List of invoices with pagination metadata
//   - error: nil on success, error on failure
func (c *Client) ListInvoices(params *developer.ListMerchantInvoicesRequest) (*developer.ListInvoicesResponse, error) {
	var resp developer.ListInvoicesResponse
	err := c.sendRequestWithQuery("GET", "/v1/invoices", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoice retrieves detailed information about a specific invoice.
// Includes line items, customer details, payment status, and due date.
//
// Parameters:
//   - invoiceID: The unique identifier of the invoice
//
// Returns:
//   - *developer.InvoiceDetails: Complete invoice details with line items
//   - error: nil on success, error on failure (e.g., invoice not found)
func (c *Client) GetInvoice(invoiceID string) (*developer.InvoiceDetails, error) {
	path := fmt.Sprintf("/v1/invoices/%s", invoiceID)
	var resp developer.InvoiceDetails
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoicePaymentIntent retrieves the payment intent associated with an invoice.
// The payment intent is used to collect payment for the invoice.
//
// Parameters:
//   - invoiceID: The unique identifier of the invoice
//
// Returns:
//   - *developer.PaymentIntent: The payment intent associated with the invoice
//   - error: nil on success, error on failure (e.g., invoice has no payment intent)
func (c *Client) GetInvoicePaymentIntent(invoiceID string) (*developer.PaymentIntent, error) {
	path := fmt.Sprintf("/v1/invoices/%s/payment_intent", invoiceID)
	var resp developer.PaymentIntent
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoicePDF downloads the invoice as a PDF document.
// The PDF can be saved to disk or sent to the customer.
//
// Parameters:
//   - invoiceID: The unique identifier of the invoice
//
// Returns:
//   - []byte: The PDF content as a byte array
//   - error: nil on success, error on failure or if not implemented
//
// Note: This functionality is currently not implemented.
func (c *Client) GetInvoicePDF(invoiceID string) ([]byte, error) {
	_ = fmt.Sprintf("/v1/invoices/%s/pdf", invoiceID)
	// TODO: implement PDF download
	return nil, fmt.Errorf("not implemented")
}

// SendInvoice sends the invoice to the customer via email.
// The customer will receive an email with the invoice details and a payment link.
//
// Parameters:
//   - invoiceID: The unique identifier of the invoice to send
//
// Returns:
//   - *developer.InvoiceDetails: The updated invoice details with sent status
//   - error: nil on success, error on failure (e.g., customer email not configured)
func (c *Client) SendInvoice(invoiceID string) (*developer.InvoiceDetails, error) {
	path := fmt.Sprintf("/v1/invoices/%s/send", invoiceID)
	var resp developer.InvoiceDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VoidInvoice voids an invoice, marking it as cancelled.
// Voided invoices cannot be paid and will not be included in revenue reports.
//
// Parameters:
//   - invoiceID: The unique identifier of the invoice to void
//
// Returns:
//   - *developer.InvoiceDetails: The updated invoice details with voided status
//   - error: nil on success, error on failure (e.g., invoice already paid)
func (c *Client) VoidInvoice(invoiceID string) (*developer.InvoiceDetails, error) {
	path := fmt.Sprintf("/v1/invoices/%s/void", invoiceID)
	var resp developer.InvoiceDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
