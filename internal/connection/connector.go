package connection

import (
	"errors"
	"log"
)

const DailyLimit = 2

func SendConnection(profile Profile, sentToday *int, sentTracker map[string]bool) error {
	if *sentToday >= DailyLimit {
		return errors.New("daily connection limit reached")
	}

	// -------------------------------
	// 1. Block duplicates (runtime)
	// -------------------------------
	if sentTracker[profile.URL] {
		log.Println("[CONNECT] Skipping (already sent this run):", profile.Name)
		return nil
	}

	// -------------------------------
	// 2. Block duplicates (persistent)
	// -------------------------------
	existing, _ := loadState(SentFile)
	for _, p := range existing {
		if p.URL == profile.URL {
			log.Println("[CONNECT] Skipping (already sent earlier):", profile.Name)
			sentTracker[profile.URL] = true
			return nil
		}
	}

	// -------------------------------
	// 3. Send connection request
	// -------------------------------
	log.Println("[CONNECT] Navigating to profile:", profile.URL)
	log.Println("[CONNECT] Connect button found")
	log.Println("[CONNECT] Sending note: Hi", profile.Name+", I'd like to connect.")

	*sentToday++
	sentTracker[profile.URL] = true

	// -------------------------------
	// 4. Save ONLY once
	// -------------------------------
	existing = append(existing, profile)
	_ = saveState(SentFile, existing)

	log.Println("[CONNECT] Request sent. Total today:", *sentToday)
	return nil
}
