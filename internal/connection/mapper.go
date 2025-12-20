package connection

import "linkedin-automation-poc/internal/search"

// FromSearchResult converts a search result into a connection profile
// This keeps search logic and connection logic cleanly separated
func FromSearchResult(r search.SearchResult) Profile {

	// Map only the required fields needed for sending connections
	return Profile{
		ID:       r.ID,
		Name:     r.Name,
		Headline: r.Headline,
		URL:      r.URL,
	}
}
