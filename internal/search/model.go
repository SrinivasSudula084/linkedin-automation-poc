package search

// SearchCriteria defines filters used for LinkedIn profile search
// These fields control how profiles are selected and matched
type SearchCriteria struct {

	// Desired job title (e.g., "golang developer")
	JobTitle string

	// Company name filter (optional)
	Company string

	// Geographic location filter (e.g., "India")
	Location string

	// Additional keywords to match in the profile headline
	Keywords string
}

// SearchResult represents a single LinkedIn search result
// Used throughout search, connection, and messaging modules
type SearchResult struct {

	// Unique identifier for the profile
	ID string `json:"id"`

	// Full name of the LinkedIn user
	Name string `json:"name"`

	// Professional headline shown on LinkedIn
	Headline string `json:"headline"`

	// Location mentioned on the profile
	Location string `json:"location"`

	// Public LinkedIn profile URL (used as unique key)
	URL string `json:"url"`

	// Relevance score (reserved for future ranking logic)
	Score int `json:"score"`
}
