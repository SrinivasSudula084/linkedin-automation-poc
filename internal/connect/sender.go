package connect

import (
	"fmt"
	"log"
	"time"

	"linkedin-automation-poc/internal/search"
)

func SendConnection(profile search.Profile, limiter *RateLimiter) error {
	// Rate limit check
	if err := limiter.Allow(); err != nil {
		return err
	}

	// Generate personalized note (POC)
	note := fmt.Sprintf(
		"Hi %s, I came across your profile and would like to connect.",
		profile.Name,
	)

	// Enforce LinkedIn note character limit (~300 chars)
	if len(note) > 300 {
		note = note[:300]
	}

	// DEMO MODE — simulate behavior
	log.Println("Navigating to profile:", profile.URL)
	log.Println("Connect button detected")
	log.Println("Sending connection request with note:")
	log.Println(note)

	// Simulated delay
	time.Sleep(1 * time.Second)

	log.Println("Connection request sent (demo mode)")
	return nil
}
