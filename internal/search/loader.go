package search

import (
	"encoding/json"
	"os"
)

// LoadProfilesFromJSON loads search result profiles from a local JSON file
// This is used for demo / POC purposes instead of live scraping
func LoadProfilesFromJSON(path string) ([]SearchResult, error) {

	// Read the entire JSON file from disk
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Deserialize JSON into SearchResult structures
	var profiles []SearchResult
	err = json.Unmarshal(data, &profiles)

	return profiles, err
}
