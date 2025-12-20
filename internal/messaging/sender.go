package messaging

import (
	"encoding/json"
	"log"
	"os"
)

// SendMessage sends a follow-up message ONLY ONCE per profile
// It prevents duplicate messages across multiple runs
func SendMessage(record *MessageRecord) error {

	// Load message state to check if a message was already sent
	sent, _ := LoadMessageState("message_state.json")

	// -------------------------------------------------
	// DUPLICATE MESSAGE PREVENTION
	// -------------------------------------------------
	if sent[record.ProfileURL] {
		log.Println("[MESSAGE] Already sent to:", record.Name)
		return nil
	}

	// -------------------------------------------------
	// MESSAGE SENDING (POC SIMULATION)
	// -------------------------------------------------
	// In real automation, this would open LinkedIn chat UI
	log.Println("[MESSAGE] Opening chat with:", record.ProfileURL)
	log.Println("[MESSAGE] Sending message:", record.Message)

	// Mark message as sent for this profile
	sent[record.ProfileURL] = true
	_ = SaveMessageState("message_state.json", sent)

	log.Println("[MESSAGE] Message sent to:", record.Name)
	return nil
}

// =====================================================
// MESSAGE STATE PERSISTENCE
// =====================================================

// LoadMessageState loads message-sent status from a JSON file
// Returns an empty map if the file does not exist (first run)
func LoadMessageState(path string) (map[string]bool, error) {

	// Attempt to read message state file
	data, err := os.ReadFile(path)
	if err != nil {
		return map[string]bool{}, nil
	}

	// Deserialize JSON into a map of profileURL → sent status
	m := map[string]bool{}
	_ = json.Unmarshal(data, &m)

	return m, nil
}

// SaveMessageState persists message-sent status to disk
// Ensures messages are not sent again in future runs
func SaveMessageState(path string, m map[string]bool) error {

	// Convert map into formatted JSON for readability
	data, _ := json.MarshalIndent(m, "", "  ")

	// Write JSON file with safe permissions
	return os.WriteFile(path, data, 0644)
}
