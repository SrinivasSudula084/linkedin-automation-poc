package search

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	// Rod page is passed to keep interface realistic,
	// even though demo data is used
	"github.com/go-rod/rod"
)

// ExtractResults filters and extracts matching profiles
// based on search criteria and avoids duplicates
// NOTE: This is demo-based (uses local JSON, no scraping)
func ExtractResults(
	page *rod.Page,
	criteria SearchCriteria,
	seen map[string]bool,
) ([]SearchResult, error) {

	// -----------------------------------------
	// LOAD DEMO PROFILE DATA
	// -----------------------------------------
	log.Println("[SEARCH] Loading demo profiles")

	data, err := os.ReadFile("demo_profiles.json")
	if err != nil {
		return nil, err
	}

	// Parse JSON into search result structures
	all := []SearchResult{}
	if err := json.Unmarshal(data, &all); err != nil {
		return nil, err
	}

	results := []SearchResult{}

	// -----------------------------------------
	// APPLY FILTERING & DEDUPLICATION
	// -----------------------------------------
	for _, p := range all {

		// ---------- JOB TITLE FILTER ----------
		if criteria.JobTitle != "" &&
			!strings.Contains(
				strings.ToLower(p.Headline),
				strings.ToLower(criteria.JobTitle),
			) {
			continue
		}

		// ---------- LOCATION FILTER ----------
		if criteria.Location != "" &&
			!strings.EqualFold(p.Location, criteria.Location) {
			continue
		}

		// ---------- KEYWORD FILTER ----------
		if criteria.Keywords != "" &&
			!strings.Contains(
				strings.ToLower(p.Headline),
				strings.ToLower(criteria.Keywords),
			) {
			continue
		}

		// ---------- DUPLICATE PREVENTION ----------
		// Ensure each profile is processed only once
		if seen[p.URL] {
			log.Println("[SEARCH] Duplicate skipped:", p.URL)
			continue
		}

		// Mark profile as seen
		seen[p.URL] = true
		results = append(results, p)

		log.Printf("[SEARCH] Found: %s | %s | %s\n",
			p.Name, p.Headline, p.Location)
	}

	log.Printf("[SEARCH] Total unique profiles: %d\n", len(results))
	return results, nil
}
