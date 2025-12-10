package developer

// ApprovalInstance represents an approval workflow instance
type ApprovalInstance struct {
	// ID is the unique identifier for the approval instance
	ID string `json:"id"`
	// Object is the object type, typically "approval_instance"
	Object string `json:"object"`
	// MerchantId is the ID of the merchant who owns this approval instance
	MerchantId string `json:"merchant_id"`
	// Status is the current status (pending, approved, rejected)
	Status string `json:"status"`
	// ResourceID is the ID of the resource being approved
	ResourceID string `json:"resource_id"`
	// ResourceType is the type of resource (payroll, refund, etc.)
	ResourceType string `json:"resource_type"`
	// CreatedAt is the Unix timestamp when the approval instance was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the approval instance was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Records is the history of approval actions
	Records []*ApprovalRecord `json:"records"`
	// CurrentStep is the current step in the approval flow
	CurrentStep *ApprovalStep `json:"current_step"`
}

// ApprovalRecord represents a single approval action record
type ApprovalRecord struct {
	// ID is the unique identifier for the approval record
	ID string `json:"id"`
	// Action is the action taken (approve, reject, start)
	Action string `json:"action"`
	// Comment is the optional comment provided by the approver
	Comment string `json:"comment"`
	// ApproverID is the ID of the user who performed this action
	ApproverID string `json:"approver_id"`
	// ApproverName is the name of the user who performed this action
	ApproverName string `json:"approver_name"`
	// ApproverRole is the role of the approver
	ApproverRole string `json:"approver_role"`
	// CreatedAt is the Unix timestamp when this record was created
	CreatedAt int64 `json:"created_at"`
	// Namespace is the namespace (merchant or platform)
	Namespace string `json:"namespace"`
}

// ApprovalFlow represents an approval workflow configuration
type ApprovalFlow struct {
	// ID is the unique identifier for the approval flow
	ID string `json:"id"`
	// Resource is the resource type this flow applies to
	Resource string `json:"resource"`
	// Name is the name of the approval flow
	Name string `json:"name"`
	// Description is the description of the approval flow
	Description string `json:"description"`
	// Version is the version number of the flow
	Version int `json:"version"`
	// Status is the status of the flow (active or inactive)
	Status string `json:"status"`
	// CreatedAt is the Unix timestamp when the flow was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when the flow was last updated
	UpdatedAt int64 `json:"updated_at"`
}

// ApprovalStep represents a single step in an approval workflow
type ApprovalStep struct {
	// ID is the unique identifier for the approval step
	ID string `json:"id"`
	// FlowID is the ID of the approval flow this step belongs to
	FlowID string `json:"flow_id"`
	// StepOrder is the order of this step in the flow
	StepOrder int `json:"step_order"`
	// Namespace is the namespace (merchant or platform)
	Namespace string `json:"namespace"`
	// Roles is the list of roles allowed to approve this step
	Roles []string `json:"roles"`
	// NextOnApprove is the next step ID when approved (0 means end)
	NextOnApprove int `json:"next_on_approve"`
	// NextOnReject is the next step ID when rejected (0 means end)
	NextOnReject int `json:"next_on_reject"`
	// CreatedAt is the Unix timestamp when this step was created
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is the Unix timestamp when this step was last updated
	UpdatedAt int64 `json:"updated_at"`
}

// ================================
// Request Types
// ================================

// PayoutApproveRequest represents a request to approve a payout
type PayoutApproveRequest struct {
	// Comment explaining the approval decision (max 1024 characters)
	Comment string `json:"comment" binding:"max=1024"`
}

// PayoutRejectRequest represents a request to reject a payout
type PayoutRejectRequest struct {
	// Comment explaining the rejection reason (max 1024 characters)
	Comment string `json:"comment" binding:"max=1024"`
}

// ApprovalGetRequest represents a request to get approval instance details
type ApprovalGetRequest struct {
	// Filter by resource ID
	ResourceID *string `json:"resource_id" form:"resource_id"`
	// Filter by resource type
	ResourceType *string `json:"resource_type" form:"resource_type"`
	// Filter by approval instance ID
	InstanceID *string `json:"instance_id" form:"instance_id"`
}

// ApprovalListRequest represents a request to list approval instances with filters
type ApprovalListRequest struct {
	// Page number, starting from 0
	Page int32 `json:"page" binding:"min=0" form:"page"`
	// Items per page (max 50)
	PageSize int32 `json:"page_size" binding:"required,min=1,max=50" form:"page_size"`
	// Filter by resource ID
	ResourceID *string `json:"resource_id" form:"resource_id"`
	// Filter by resource type
	ResourceType *string `json:"resource_type" form:"resource_type"`
	// Filter by approval status
	Status *string `json:"status" form:"status"`
}

// ApprovalCreateRequest represents a request to create a new approval instance
type ApprovalCreateRequest struct {
	// Resource ID that requires approval
	ResourceID string `json:"resource_id" binding:"required"`
	// Type of the resource requiring approval
	ResourceType string `json:"resource_type" binding:"required"`
}

// ================================
// Response Types
// ================================

// ApprovalListResp represents the response containing a paginated list of approval instances
type ApprovalListResp struct {
	// List of approval instances
	Approvals []*ApprovalInstance `json:"approvals"`
	// Total number of records matching the filters
	Total int64 `json:"total"`
	// Current page number
	Page int32 `json:"page"`
	// Items per page
	PageSize int32 `json:"page_size"`
}
