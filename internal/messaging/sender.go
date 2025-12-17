package messaging

import "log"

// SendMessage simulates sending a LinkedIn message (DEMO MODE)
func SendMessage(record *MessageRecord) error {
	log.Println("[MESSAGE] Opening chat with:", record.ProfileURL)
	log.Println("[MESSAGE] Sending message:", record.Message)
	log.Println("[MESSAGE] Message sent successfully")
	return nil
}
