package search

import (
	"encoding/json"
	"os"
)

func LoadProfilesFromJSON(path string) ([]SearchResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profiles []SearchResult
	err = json.Unmarshal(data, &profiles)
	return profiles, err
}
