package messaging

// MessageRecord represents a single message interaction
// It is used to track what message was sent to which profile
type MessageRecord struct {

	// LinkedIn profile URL (used as a unique identifier)
	ProfileURL string `json:"profile_url"`

	// Name of the connected user
	Name string `json:"name"`

	// Actual message content that was sent
	Message string `json:"message"`

	// Indicates whether the message was successfully sent
	// Helps prevent duplicate follow-up messages
	Sent bool `json:"sent"`
}
