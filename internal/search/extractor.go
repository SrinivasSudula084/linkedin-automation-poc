package search

import "github.com/go-rod/rod"

func ExtractResults(page *rod.Page, seen map[string]bool) ([]SearchResult, error) {
	results := []SearchResult{}

	links, err := page.Elements(`a[href*="/in/"]`)
	if err != nil {
		return results, err
	}

	for _, link := range links {
		href, _ := link.Attribute("href")
		if href == nil || seen[*href] {
			continue
		}

		seen[*href] = true
		results = append(results, SearchResult{
			URL: *href,
		})
	}

	return results, nil
}
