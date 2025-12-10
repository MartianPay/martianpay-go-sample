package developer

const (
	RoleObject   = "role"
	RoleIDLength = 24
	RoleIDPrefix = "role_"
)

// Permission defines access control rules for resources
type Permission struct {
	// Allow is the list of allowed resource patterns
	Allow []string `json:"allow"`
	// Deny is the list of denied resource patterns
	Deny []string `json:"deny"`
}

// Role represents a user role with associated permissions
type Role struct {
	// ID is the unique identifier for the role
	ID string `json:"id"`
	// Object is the type identifier, always "role"
	Object string `json:"object"`
	// Name is the role name (e.g., "admin", "developer", "viewer")
	Name string `json:"name"`
	// Description is a human-readable description of the role
	Description string `json:"description"`
	// Permissions defines the access control rules for this role
	Permissions *Permission `json:"permissions,omitempty"`
	// Policies is the list of policy rules associated with this role
	Policies []*Policy `json:"policies"`
}

// Policy represents a single access control policy rule
type Policy struct {
	// Role is the role name this policy applies to
	Role string
	// Resource is the resource pattern this policy controls
	Resource string
	// Action is the action type (e.g., "read", "write", "delete")
	Action string
	// Allow indicates whether the action is allowed (true) or denied (false)
	Allow bool
}
