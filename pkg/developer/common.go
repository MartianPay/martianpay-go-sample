// common.go contains common data structures and utilities shared across the SDK.
package developer

// Pagination contains pagination parameters for list requests.
// It is used by all list endpoints to paginate results consistently across the API.
//
// Example usage:
//
//	req := &ProductListRequest{
//	    Pagination: Pagination{
//	        Page:     0,     // First page (zero-indexed)
//	        PageSize: 20,    // 20 items per page
//	    },
//	}
type Pagination struct {
	// Page is the page number to retrieve (zero-indexed, default is 0)
	Page int32 `json:"page" form:"page" binding:"min=0"`
	// PageSize is the number of items per page (default is 10, max is 50)
	PageSize int32 `json:"page_size" form:"page_size" binding:"omitempty,min=1,max=50"`
}

// SetDefault sets default values for pagination parameters.
// It ensures that Page is non-negative (defaults to 0) and PageSize is set to 10 if not specified.
// This method is typically called before making API requests to ensure valid pagination values.
func (p *Pagination) SetDefault() {
	if p.Page < 0 {
		p.Page = 0
	}

	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
