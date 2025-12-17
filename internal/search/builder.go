package search

import "net/url"

func BuildSearchURL(c SearchCriteria) string {
	base := "https://www.linkedin.com/search/results/people/"
	params := url.Values{}

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

	return base + "?" + params.Encode()
}
