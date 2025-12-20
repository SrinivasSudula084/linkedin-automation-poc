package state

// SentRequest represents a connection request
// that has already been sent to a LinkedIn profile
type SentRequest struct {

	// Unique identifier of the profile
	ID string `json:"id"`

	// Name of the LinkedIn user
	Name string `json:"name"`

	// Public LinkedIn profile URL (used as unique key)
	URL string `json:"url"`
}

// ConnectedProfile represents a profile
// that has accepted the connection request
type ConnectedProfile struct {

	// Unique identifier of the profile
	ID string `json:"id"`

	// Name of the LinkedIn user
	Name string `json:"name"`

	// Public LinkedIn profile URL
	URL string `json:"url"`
}
