package connection

import (
	"encoding/json"
	"os"
)

// File paths used for persistent state storage
// These JSON files allow the application to remember actions
// across multiple runs
const (
	SentFile      = "sent_requests.json"
	ConnectedFile = "connected_profiles.json"
)

// loadState reads profile data from a JSON file
// If the file does not exist, it safely returns an empty list
func loadState(path string) ([]Profile, error) {

	// Check if the file exists
	// If not, treat it as a first-run scenario
	if _, err := os.Stat(path); err != nil {
		return []Profile{}, nil
	}

	// Read file contents
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Deserialize JSON into Profile slice
	var profiles []Profile
	_ = json.Unmarshal(data, &profiles)

	return profiles, nil
}

// saveState writes profile data to a JSON file
// Used to persist sent requests and connected profiles
func saveState(path string, profiles []Profile) error {

	// Convert profiles into formatted JSON for readability
	data, _ := json.MarshalIndent(profiles, "", "  ")

	// Save JSON to disk with safe file permissions
	return os.WriteFile(path, data, 0644)
}
