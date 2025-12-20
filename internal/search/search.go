package search

import (
	"log"
	"strings"
)

// SearchProfiles filters and scores profiles based on search criteria
// It performs partial matching and avoids duplicate results
func SearchProfiles(
	all []SearchResult,
	criteria SearchCriteria,
	seen map[string]bool,
) []SearchResult {

	// Holds all matching profiles
	results := []SearchResult{}

	for _, p := range all {

		// -------------------------------------------------
		// DUPLICATE DETECTION
		// -------------------------------------------------
		// Skip profiles that were already processed
		if seen[p.URL] {
			log.Println("[SEARCH] Duplicate skipped:", p.URL)
			continue
		}

		// Score represents how relevant the profile is
		score := 0

		// -------------------------------------------------
		// PARTIAL MATCHING & SCORING
		// -------------------------------------------------
		// Job title match has higher weight
		if strings.Contains(
			strings.ToLower(p.Headline),
			strings.ToLower(criteria.JobTitle),
		) {
			score += 3
		}

		// Keyword match adds medium relevance
		if strings.Contains(
			strings.ToLower(p.Headline),
			strings.ToLower(criteria.Keywords),
		) {
			score += 2
		}

		// Location match adds lower relevance
		if strings.Contains(
			strings.ToLower(p.Location),
			strings.ToLower(criteria.Location),
		) {
			score += 1
		}

		// If no criteria matched, skip the profile
		if score == 0 {
			continue
		}

		// Assign computed score to profile
		p.Score = score

		// Mark profile as processed
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
