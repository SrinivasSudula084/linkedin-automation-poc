package messaging

import (
	"encoding/json"
	"os"
)

// SaveMessageHistory persists the full message history to disk
// This provides an audit trail of all messages sent by the system
func SaveMessageHistory(path string, records []MessageRecord) error {

	// Convert message records into readable, formatted JSON
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	// Write message history to file with safe permissions
	return os.WriteFile(path, data, 0644)
}
