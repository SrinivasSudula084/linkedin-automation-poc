package messaging

import (
	"encoding/json"
	"log"
	"os"
)

// SendMessage sends follow-up ONLY ONCE per profile
func SendMessage(record *MessageRecord) error {
	sent, _ := LoadMessageState("message_state.json")

	if sent[record.ProfileURL] {
		log.Println("[MESSAGE] Already sent to:", record.Name)
		return nil
	}

	log.Println("[MESSAGE] Opening chat with:", record.ProfileURL)
	log.Println("[MESSAGE] Sending message:", record.Message)

	sent[record.ProfileURL] = true
	_ = SaveMessageState("message_state.json", sent)

	log.Println("[MESSAGE] Message sent to:", record.Name)
	return nil
}

// ---- Message state ----

func LoadMessageState(path string) (map[string]bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return map[string]bool{}, nil
	}
	m := map[string]bool{}
	_ = json.Unmarshal(data, &m)
	return m, nil
}

func SaveMessageState(path string, m map[string]bool) error {
	data, _ := json.MarshalIndent(m, "", "  ")
	return os.WriteFile(path, data, 0644)
}
