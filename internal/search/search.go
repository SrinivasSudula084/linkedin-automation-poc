package search

import (
	"log"
	"strings"
)

func SearchProfiles(
	all []SearchResult,
	criteria SearchCriteria,
	seen map[string]bool,
) []SearchResult {

	results := []SearchResult{}

	for _, p := range all {

		// -------- Duplicate detection --------
		if seen[p.URL] {
			log.Println("[SEARCH] Duplicate skipped:", p.URL)
			continue
		}

		score := 0

		// -------- Partial matching --------
		if strings.Contains(
			strings.ToLower(p.Headline),
			strings.ToLower(criteria.JobTitle),
		) {
			score += 3
		}

		if strings.Contains(
			strings.ToLower(p.Headline),
			strings.ToLower(criteria.Keywords),
		) {
			score += 2
		}

		if strings.Contains(
			strings.ToLower(p.Location),
			strings.ToLower(criteria.Location),
		) {
			score += 1
		}

		if score == 0 {
			continue
		}

		p.Score = score
		seen[p.URL] = true
		results = append(results, p)

		log.Printf(
			"[SEARCH] Match: %s | %s | %s | score=%d",
			p.Name,
			p.Headline,
			p.Location,
			score,
		)
	}

	log.Printf("[SEARCH] %d matching profiles found", len(results))
	return results
}
