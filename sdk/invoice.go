package martianpay

import (
	"fmt"

	"github.com/MartianPay/martianpay-go-sample/pkg/developer"
)

// ListInvoices lists merchant invoices
func (c *Client) ListInvoices(params *developer.ListMerchantInvoicesRequest) (*developer.ListInvoicesResponse, error) {
	var resp developer.ListInvoicesResponse
	err := c.sendRequestWithQuery("GET", "/v1/invoices", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoice gets merchant invoice
func (c *Client) GetInvoice(invoiceID string) (*developer.InvoiceDetails, error) {
	path := fmt.Sprintf("/v1/invoices/%s", invoiceID)
	var resp developer.InvoiceDetails
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoicePaymentIntent gets invoice payment intent
func (c *Client) GetInvoicePaymentIntent(invoiceID string) (*developer.PaymentIntent, error) {
	path := fmt.Sprintf("/v1/invoices/%s/payment_intent", invoiceID)
	var resp developer.PaymentIntent
	err := c.sendRequest("GET", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoicePDF downloads invoice PDF
func (c *Client) GetInvoicePDF(invoiceID string) ([]byte, error) {
	_ = fmt.Sprintf("/v1/invoices/%s/pdf", invoiceID)
	// TODO: implement PDF download
	return nil, fmt.Errorf("not implemented")
}

// SendInvoice sends invoice to customer
func (c *Client) SendInvoice(invoiceID string) (*developer.InvoiceDetails, error) {
	path := fmt.Sprintf("/v1/invoices/%s/send", invoiceID)
	var resp developer.InvoiceDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VoidInvoice voids invoice
func (c *Client) VoidInvoice(invoiceID string) (*developer.InvoiceDetails, error) {
	path := fmt.Sprintf("/v1/invoices/%s/void", invoiceID)
	var resp developer.InvoiceDetails
	err := c.sendRequest("POST", path, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
