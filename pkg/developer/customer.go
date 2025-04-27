package developer

import "github.com/dchest/uniuri"

const (
	CustomerObject   = "customer"
	CustomerIDLength = 24
	CustomerIDPrefix = "cus_"
)

type CustomerParams struct {
	Name        *string           `json:"name"`
	Email       *string           `json:"email" binding:"required,email"`
	Description *string           `json:"description"`
	Metadata    map[string]string `json:"metadata"`
	Phone       *string           `json:"phone"`
}

type Customer struct {
	ID     string `json:"id"`
	Object string `json:"object"` // "customer"

	TotalExpense int64 `json:"total_expense"`
	TotalPayment int64 `json:"total_payment"`
	TotalRefund  int64 `json:"total_refund"`

	Currency string `json:"currency,omitempty"`

	Created int64 `json:"created,omitempty"`

	Name        *string           `json:"name,omitempty"`
	Email       *string           `json:"email,omitempty"`
	Description *string           `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Phone       *string           `json:"phone,omitempty"`
}

func (c *Customer) GenerateCustomerID() string {
	return CustomerIDPrefix + uniuri.NewLen(CustomerIDLength)
}
