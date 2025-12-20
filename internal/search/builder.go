package search

import "net/url"

// BuildSearchURL constructs a LinkedIn people search URL
// based on the provided search criteria
func BuildSearchURL(c SearchCriteria) string {

	// Base LinkedIn search endpoint for people
	base := "https://www.linkedin.com/search/results/people/"

	// URL parameters container
	params := url.Values{}

	// -------------------------------------------------
	// Add search filters dynamically
	// -------------------------------------------------
	// Only non-empty criteria are included in the URL
	if c.JobTitle != "" {
		params.Add("title", c.JobTitle)
	}
	if c.Company != "" {
		params.Add("company", c.Company)
	}
	if c.Location != "" {
		params.Add("location", c.Location)
	}
	if c.Keywords != "" {
		params.Add("keywords", c.Keywords)
	}

	// Encode parameters and return full search URL
	return base + "?" + params.Encode()
}
