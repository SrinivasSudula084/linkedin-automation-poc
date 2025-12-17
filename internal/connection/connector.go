package connection

import (
	"errors"
	"log"
)

const DailyLimit = 2

func SendConnection(profile Profile, sentToday *int) error {
	if *sentToday >= DailyLimit {
		return errors.New("daily connection limit reached")
	}

	log.Println("[CONNECT] Navigating to profile:", profile.URL)
	log.Println("[CONNECT] Connect button found")
	log.Println("[CONNECT] Sending note: Hi", profile.Name, ", I'd like to connect.")

	*sentToday++

	// Save to sent_requests.json
	sent, _ := loadState("sent_requests.json")
	sent = append(sent, profile)
	_ = saveState("sent_requests.json", sent)

	log.Println("[CONNECT] Request sent. Total today:", *sentToday)
	return nil
}
