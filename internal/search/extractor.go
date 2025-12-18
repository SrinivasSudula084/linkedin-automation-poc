package search

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/go-rod/rod"
)

func ExtractResults(
	page *rod.Page,
	criteria SearchCriteria,
	seen map[string]bool,
) ([]SearchResult, error) {

	log.Println("[SEARCH] Loading demo profiles")
	data, err := os.ReadFile("demo_profiles.json")
	if err != nil {
		return nil, err
	}

	all := []SearchResult{}
	if err := json.Unmarshal(data, &all); err != nil {
		return nil, err
	}

	results := []SearchResult{}

	for _, p := range all {

		// ---------- FILTERING ----------
		if criteria.JobTitle != "" &&
			!strings.Contains(strings.ToLower(p.Headline),
				strings.ToLower(criteria.JobTitle)) {
			continue
		}

		if criteria.Location != "" &&
			!strings.EqualFold(p.Location, criteria.Location) {
			continue
		}

		if criteria.Keywords != "" &&
			!strings.Contains(strings.ToLower(p.Headline),
				strings.ToLower(criteria.Keywords)) {
			continue
		}

		// ---------- DEDUPLICATION ----------
		if seen[p.URL] {
			log.Println("[SEARCH] Duplicate skipped:", p.URL)
			continue
		}

		seen[p.URL] = true
		results = append(results, p)

		log.Printf("[SEARCH] Found: %s | %s | %s\n",
			p.Name, p.Headline, p.Location)
	}

	log.Printf("[SEARCH] Total unique profiles: %d\n", len(results))
	return results, nil
}
