package schema

type EmailAddress struct {
	EmailAddress *string `json:"email_address"`
}
type ClerkUser struct {
	// # "Create" & "Update" #
	ID            string          `json:"id"`
	FirstName     *string         `json:"first_name"`
	LastName      *string         `json:"last_name"`
	ImageURL      *string         `json:"image_url"`
	EmailAddresss []*EmailAddress `json:"email_addresses"`

	// # "Delete" #
	Deleted bool   `json:"deleted"` // # "false" #
	Object  string `json:"object"`  // # "user" #
}

type ClerkUserRequest struct {
	InstanceID string     `json:"instance_id"`
	Object     string     `json:"object"` // # "event" #
	Type       string     `json:"type"`   // # "event type" #
	Data       *ClerkUser `json:"data"`   // # "clerk user" #
}
