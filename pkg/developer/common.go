package developer

// Pagination contains pagination parameters for list requests
type Pagination struct {
	// Page is the page number to retrieve (zero-indexed, default is 0)
	Page int32 `json:"page" form:"page" binding:"min=0"`
	// PageSize is the number of items per page (default is 10, max is 50)
	PageSize int32 `json:"page_size" form:"page_size" binding:"omitempty,min=1,max=50"`
}

func (p *Pagination) SetDefault() {
	if p.Page < 0 {
		p.Page = 0
	}

	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
