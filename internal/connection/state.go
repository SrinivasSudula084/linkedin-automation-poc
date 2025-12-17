package connection

import (
	"encoding/json"
	"os"
	
)


// SaveConnectedProfiles saves connected profiles to JSON
func SaveConnectedProfiles(path string, profiles []Profile) error {
	data, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// LoadConnectedProfiles loads connected profiles from JSON
func LoadConnectedProfiles(path string) ([]Profile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}
