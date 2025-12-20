package connection

import (
	"encoding/json"
	"os"
)

// LoadDemoProfiles loads profile data from a local JSON file
// This is used only for demo / POC purposes (no real scraping)
func LoadDemoProfiles(path string) ([]Profile, error) {

	// Read the entire JSON file from disk
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse JSON data into a slice of Profile structs
	var profiles []Profile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, err
	}

	// Return the list of demo profiles
	return profiles, nil
}
