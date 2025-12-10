package developer

const (
	MemberObject   = "member"
	MemberIDLength = 24
	MemberIDPrefix = "member_"
)

const (
	InviteIDLength = 36
	InviteIDPrefix = "invite_"
)

const (
	MemberStatusActive   = "active"
	MemberStatusInactive = "inactive"
	MemberStatusInviting = "inviting"
)

// Member represents a member of a merchant account
type Member struct {
	// ID is the unique identifier for the member
	ID string `json:"id"`
	// Object is the type identifier, always "member"
	Object string `json:"object"`
	// Email is the member's email address
	Email string `json:"email"`
	// Owner indicates whether this member is the account owner
	Owner bool `json:"owner"`
	// Roles is the list of roles assigned to this member
	Roles []Role `json:"roles"`
	// Status is the current status of the member (active, inactive, inviting)
	Status string `json:"status"`
	// InviteId is the ID of the invitation if the member is in inviting status
	InviteId *string `json:"invite_id"`
}

// Invite represents an invitation to join a merchant account
type Invite struct {
	// ID is the unique identifier for the invitation
	ID string `json:"id"`
	// MerchantName is the name of the merchant the user is invited to
	MerchantName string `json:"merchant_name"`
	// Inviter is the email of the person who sent the invitation
	Inviter *string `json:"inviter"`
	// Invitee is the email address of the person being invited
	Invitee string `json:"invitee"`
}
