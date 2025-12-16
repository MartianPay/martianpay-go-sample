// order.go contains types for managing e-commerce orders.
// It provides structures for order management, order items, shipping information,
// and order status tracking for payment link transactions.
package developer

// ================================
// Constants and Enums
// ================================

// OrderStatus represents the status of an order
type OrderStatus string

const (
	// OrderStatusPending indicates the order is pending payment
	OrderStatusPending OrderStatus = "Pending"
	// OrderStatusPaid indicates the order has been paid
	OrderStatusPaid OrderStatus = "Paid"
	// OrderStatusShipping indicates the order is being shipped
	OrderStatusShipping OrderStatus = "Shipping"
	// OrderStatusRefunded indicates the order has been refunded
	OrderStatusRefunded OrderStatus = "Refunded"
	// OrderStatusCanceled indicates the order has been canceled
	OrderStatusCanceled OrderStatus = "Canceled"
)

// ================================
// Core Types
// ================================

// OrderItem represents a single product item within an order
type OrderItem struct {
	// ProductID is the unique identifier of the product
	ProductID string `json:"product_id"`

	// ProductName is the display name of the product
	ProductName string `json:"product_name"`

	// ProductDescription is the description of the product
	ProductDescription string `json:"product_description,omitempty"`

	// VariantID is the unique identifier of the product variant (if applicable)
	VariantID string `json:"variant_id,omitempty"`

	// OptionValues contains the variant options (e.g., {"Color": "Black", "Size": "M"})
	OptionValues map[string]string `json:"option_values,omitempty"`

	// Quantity is the number of units purchased
	Quantity int `json:"quantity"`

	// UnitPrice is the price per unit in the specified currency
	UnitPrice *AssetAmount `json:"unit_price"`

	// Total is the total price for this item (quantity * unit_price)
	Total *AssetAmount `json:"total"`

	// MediaOrder contains ordered list of media asset IDs for the variant or product
	MediaOrder []string `json:"media_order,omitempty"`

	// Product contains the full product details (same as payment link response)
	Product *Product `json:"product,omitempty"`

	// Variant contains the full variant details (same as payment link response)
	Variant *ProductVariant `json:"variant,omitempty"`
}

// OrderCustomer contains customer information for the order
type OrderCustomer struct {
	// ID is the customer ID
	ID string `json:"id,omitempty"`

	// Name is the customer name
	Name string `json:"name,omitempty"`

	// Email is the customer email address
	Email string `json:"email"`

	// Phone is the customer phone number
	Phone *string `json:"phone,omitempty"`
}

// OrderShippingAddress contains the shipping destination
type OrderShippingAddress struct {
	// Country is the two-letter ISO country code
	Country string `json:"country"`

	// State is the state/province code
	State *string `json:"state,omitempty"`

	// City is the city name
	City string `json:"city"`

	// PostalCode is the postal/ZIP code
	PostalCode string `json:"postal_code"`

	// Line1 is the first line of the street address
	Line1 string `json:"line1"`

	// Line2 is the second line of the street address (optional)
	Line2 *string `json:"line2,omitempty"`
}

// OrderTaxRegion contains tax jurisdiction information
type OrderTaxRegion struct {
	// Country is the two-letter ISO country code
	Country string `json:"country"`

	// State is the state/province code
	State *string `json:"state,omitempty"`
}

// OrderListItem represents an order in list view
type OrderListItem struct {
	// OrderNumber is the human-readable order identifier (e.g., "ORD-2025-01047")
	OrderNumber string `json:"order_number"`

	// ExternalID is the external system identifier for the order
	ExternalID string `json:"external_id"`

	// Customer contains the customer information associated with the order
	Customer *OrderCustomer `json:"customer"`

	// ShippingAddress contains the shipping destination (if available)
	ShippingAddress *OrderShippingAddress `json:"shipping_address,omitempty"`

	// TaxRegion contains the tax jurisdiction information (if available)
	TaxRegion *OrderTaxRegion `json:"tax_region,omitempty"`

	// Items contains the list of all items in the order
	Items []OrderItem `json:"items"`

	// TotalAmount is the total amount for the order including all items
	TotalAmount *AssetAmount `json:"total_amount"`

	// Status represents the current status of the order
	Status OrderStatus `json:"status"`

	// PaymentMethod indicates the payment method used ("Crypto", "Paypal", "Stripe")
	PaymentMethod string `json:"payment_method"`

	// CreatedAt is the Unix timestamp when the order was created
	CreatedAt int64 `json:"created_at"`

	// UpdatedAt is the Unix timestamp when the order was last updated
	UpdatedAt int64 `json:"updated_at"`
}

// OrderDetail represents full order details
type OrderDetail struct {
	// OrderNumber is the human-readable order identifier (e.g., "ORD-2025-01047")
	OrderNumber string `json:"order_number"`

	// ExternalID is the external system identifier for the order
	ExternalID string `json:"external_id"`

	// Customer contains the customer information associated with the order
	Customer *OrderCustomer `json:"customer"`

	// ShippingAddress contains the shipping destination (if available)
	ShippingAddress *OrderShippingAddress `json:"shipping_address,omitempty"`

	// TaxRegion contains the tax jurisdiction information (if available)
	TaxRegion *OrderTaxRegion `json:"tax_region,omitempty"`

	// Items contains the list of all items in the order
	Items []OrderItem `json:"items"`

	// TotalAmount is the total amount for the order including all items
	TotalAmount *AssetAmount `json:"total_amount"`

	// Status represents the current status of the order
	Status OrderStatus `json:"status"`

	// PaymentMethod indicates the payment method used ("Crypto", "Paypal", "Stripe")
	PaymentMethod string `json:"payment_method"`

	// CreatedAt is the Unix timestamp when the order was created
	CreatedAt int64 `json:"created_at"`

	// UpdatedAt is the Unix timestamp when the order was last updated
	UpdatedAt int64 `json:"updated_at"`
}

// ================================
// Request Types
// ================================

// OrderListRequest lists orders with filters
type OrderListRequest struct {
	Pagination

	// Search by order number, customer name, email
	Search *string `json:"search" form:"search"`

	// Filter by status
	Status *OrderStatus `json:"status" form:"status"`
}

// OrderDetailRequest gets order details
type OrderDetailRequest struct {
	// OrderNumber is the order identifier
	OrderNumber string `json:"order_number"`
}

// ================================
// Response Types
// ================================

// OrderListResponse represents order list response
type OrderListResponse struct {
	// Orders is the list of orders
	Orders []OrderListItem `json:"orders"`

	// Total is the total number of orders matching the filters
	Total int32 `json:"total"`

	// Page is the current page number
	Page int32 `json:"page"`

	// PageSize is the number of items per page
	PageSize int32 `json:"page_size"`
}

// OrderDetailResponse represents order detail response
type OrderDetailResponse struct {
	// Order contains the complete order details
	Order OrderDetail `json:"order"`
}
