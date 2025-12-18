package connection

import "linkedin-automation-poc/internal/search"

// FromSearchResult maps a search result to a connection profile
func FromSearchResult(r search.SearchResult) Profile {
	return Profile{
		ID:       r.ID,
		Name:     r.Name,
		Headline: r.Headline,
		URL:      r.URL,
	}
}
