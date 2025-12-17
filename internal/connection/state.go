package connection

import (
	"encoding/json"
	"os"
)

func loadState(path string) ([]Profile, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Profile{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	err = json.Unmarshal(data, &profiles)
	return profiles, err
}

func saveState(path string, profiles []Profile) error {
	data, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
