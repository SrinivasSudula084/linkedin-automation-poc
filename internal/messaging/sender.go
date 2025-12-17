package messaging

import "log"

// SendMessage simulates sending a LinkedIn message
func SendMessage(record *MessageRecord) error {
	log.Println("Opening chat with:", record.ProfileURL)
	log.Println("Sending message:", record.Message)

	record.Sent = true
	log.Println("Message sent successfully")

	return nil
}
