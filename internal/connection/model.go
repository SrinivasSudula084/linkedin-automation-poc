package connection

// Profile represents a LinkedIn user profile
// This structure is used across connection, state, and messaging modules
type Profile struct {
	// Unique identifier for the profile
	ID string `json:"id"`

	// Full name of the LinkedIn user
	Name string `json:"name"`

	// Professional headline shown on LinkedIn
	Headline string `json:"headline"`

	// Public LinkedIn profile URL (used as a unique key)
	URL string `json:"url"`
}
