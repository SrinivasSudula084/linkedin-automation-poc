package search

// SearchCriteria defines how users are searched
type SearchCriteria struct {
	JobTitle string
	Company  string
	Location string
	Keywords string
}

// SearchResult represents a profile found via search
type SearchResult struct {
	Name     string
	Headline string
	URL      string
}
