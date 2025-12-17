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

	log.Println("Navigating to profile:", profile.URL)
	log.Println("Connect button found")
	log.Println("Sending personalized note to:", profile.Name)

	*sentToday = *sentToday + 1
	log.Println("Connection request sent. Total today:", *sentToday)

	return nil
}
