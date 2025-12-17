package messaging

import (
	"encoding/json"
	"os"
)

// SaveMessageHistory persists messages
func SaveMessageHistory(path string, records []MessageRecord) error {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
