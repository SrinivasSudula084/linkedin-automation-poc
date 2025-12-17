// internal/connection/mapper.go
package connection

import "linkedin-automation-poc/internal/search"

func FromSearchResult(r search.SearchResult) Profile {
	return Profile{
		Name:     r.Name,
		Headline: r.Headline,
		URL:      r.URL,
	}
}
