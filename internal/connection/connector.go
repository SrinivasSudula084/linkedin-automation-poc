package connection

import (
	"errors"
	"log"
)

// DailyLimit defines the maximum number of connection requests
// allowed per execution (safety + anti-spam control)
const DailyLimit = 2

// SendConnection handles sending a LinkedIn connection request
// It enforces limits, avoids duplicates, and persists state safely
func SendConnection(profile Profile, sentToday *int, sentTracker map[string]bool) error {

	// -------------------------------------------------
	// GLOBAL DAILY LIMIT CHECK
	// -------------------------------------------------
	// Prevents excessive connection requests in a single run
	if *sentToday >= DailyLimit {
		return errors.New("daily connection limit reached")
	}

	// -------------------------------------------------
	// 1. RUNTIME DUPLICATE PROTECTION
	// -------------------------------------------------
	// Prevent sending multiple requests to the same profile
	// during the current execution
	if sentTracker[profile.URL] {
		log.Println("[CONNECT] Skipping (already sent this run):", profile.Name)
		return nil
	}

	// -------------------------------------------------
	// 2. PERSISTENT DUPLICATE PROTECTION
	// -------------------------------------------------
	// Load previously sent connection requests from storage
	existing, _ := loadState(SentFile)

	for _, p := range existing {
		if p.URL == profile.URL {
			// Profile already contacted in a previous run
			log.Println("[CONNECT] Skipping (already sent earlier):", profile.Name)
			sentTracker[profile.URL] = true
			return nil
		}
	}

	// -------------------------------------------------
	// 3. SEND CONNECTION REQUEST (POC SIMULATION)
	// -------------------------------------------------
	// In real automation this would click the Connect button
	log.Println("[CONNECT] Navigating to profile:", profile.URL)
	log.Println("[CONNECT] Connect button found")
	log.Println("[CONNECT] Sending note: Hi", profile.Name+", I'd like to connect.")

	// Update runtime counters
	*sentToday++
	sentTracker[profile.URL] = true

	// -------------------------------------------------
	// 4. PERSIST STATE (SAVE ONCE)
	// -------------------------------------------------
	// Save the sent profile so it is not contacted again
	existing = append(existing, profile)
	_ = saveState(SentFile, existing)

	log.Println("[CONNECT] Request sent. Total today:", *sentToday)
	return nil
}
