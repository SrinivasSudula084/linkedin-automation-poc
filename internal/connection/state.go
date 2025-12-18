package connection

import (
	"encoding/json"
	"os"
)

const (
	SentFile      = "sent_requests.json"
	ConnectedFile = "connected_profiles.json"
)

func loadState(path string) ([]Profile, error) {
	if _, err := os.Stat(path); err != nil {
		return []Profile{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profiles []Profile
	_ = json.Unmarshal(data, &profiles)
	return profiles, nil
}

func saveState(path string, profiles []Profile) error {
	data, _ := json.MarshalIndent(profiles, "", "  ")
	return os.WriteFile(path, data, 0644)
}
