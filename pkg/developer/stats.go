package developer

// ================================
// Request Types
// ================================

// StatsBalanceFlowListReq represents a request to list merchant balance flows with pagination
type StatsBalanceFlowListReq struct {
	// Page number, starting from 0
	Page int32 `json:"page" binding:"min=0"`
	// Items per page
	PageSize int32 `json:"page_size" binding:"required,min=1"`
	// Currency to filter balance flows
	Currency string `json:"currency" binding:"required"`
}

// StatsChartReq represents a request to retrieve statistics chart data
type StatsChartReq struct {
	// Category of the stats chart
	Category string `json:"category"`
	// Title of the stats chart
	Title string `json:"title"`
	// Time unit, can be hour, day, week, month
	Unit string `json:"unit"`
	// Start timestamp for current period
	CurrentStartTime int64 `json:"current_start_time"`
	// End timestamp for current period
	CurrentEndTime int64 `json:"current_end_time"`
	// Start timestamp for previous period
	PreviousStartTime int64 `json:"previous_start_time"`
	// End timestamp for previous period
	PreviousEndTime int64 `json:"previous_end_time"`
}

// ================================
// Response Types
// ================================

// ValueList represents a time-series data point in statistics
type ValueList struct {
	// Timestamp for the data point
	Time int64 `json:"time"`
	// Value at this time point
	Value string `json:"value"`
	// Percentage change compared to previous period
	Percent int64 `json:"percent"`
}

// StatsChartItem represents statistics chart data with time-series values
type StatsChartItem struct {
	// Category of the stats
	Category string `json:"category"`
	// Title of the chart
	Title string `json:"title"`
	// Time unit (hour, day, week, month)
	Unit string `json:"unit"`
	// Unit of the value (e.g., USD, count)
	ValueUnit string `json:"value_unit"`
	// Total value for current period
	TotalCurrentValue string `json:"total_current_value"`
	// Total value for previous period
	TotalPreviousValue string `json:"total_previous_value"`
	// Overall percentage change
	TotalPercent int64 `json:"total_percent"`
	// List of values for current period
	CurrentValue []*ValueList `json:"current_value"`
	// List of values for previous period
	PreviousValue []*ValueList `json:"previous_value"`
}

// StatsBalanceFlowListResp represents the response containing balance flow list with pagination
type StatsBalanceFlowListResp struct {
	// List of balance flow entries
	BalanceFlows []*MerchantBalanceFlow `json:"balance_flows"`
	// Current page number
	Page int32 `json:"page"`
	// Number of items per page
	PageSize int32 `json:"page_size"`
	// Total number of balance flows
	Total int32 `json:"total"`
}

// StatsChartResp represents the response containing statistics chart data
type StatsChartResp struct {
	StatsChartItem
}
