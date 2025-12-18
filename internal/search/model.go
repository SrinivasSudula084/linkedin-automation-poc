package search

type SearchCriteria struct {
	JobTitle string
	Company  string
	Location string
	Keywords string
}

type SearchResult struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Headline string `json:"headline"`
	Location string `json:"location"`
	URL      string `json:"url"`
	Score    int    `json:"score"`
}
