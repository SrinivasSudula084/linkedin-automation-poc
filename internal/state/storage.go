package state

import (
	"encoding/json"
	"os"
)

func LoadJSON[T any](path string) ([]T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []T{}, nil // file may not exist yet
	}

	var items []T
	err = json.Unmarshal(data, &items)
	return items, err
}

func SaveJSON[T any](path string, data []T) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}
