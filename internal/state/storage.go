package state

import (
	"encoding/json"
	"os"
)

// LoadJSON is a generic helper function to load JSON data from a file
// It can be used with any struct type (SentRequest, ConnectedProfile, etc.)
func LoadJSON[T any](path string) ([]T, error) {

	// Read the JSON file from disk
	data, err := os.ReadFile(path)
	if err != nil {
		// File may not exist on first run
		return []T{}, nil
	}

	// Deserialize JSON into a slice of the given type
	var items []T
	err = json.Unmarshal(data, &items)

	return items, err
}

// SaveJSON is a generic helper function to save data as JSON
// Uses formatted output for readability and debugging
func SaveJSON[T any](path string, data []T) error {

	// Convert data into indented JSON format
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write JSON file with safe permissions
	return os.WriteFile(path, bytes, 0644)
}
