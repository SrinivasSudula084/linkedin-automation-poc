package stealth

import (
	"math/rand"
	"time"
)

// IsWithinBusinessHours checks whether the current time
// falls inside a normal working window
func IsWithinBusinessHours(startHour, endHour int) bool {

	// Get current local time
	now := time.Now()
	hour := now.Hour()

	// Allow activity only between startHour and endHour
	return hour >= startHour && hour < endHour
}

// WaitUntilBusinessHours pauses execution until business hours begin
// This prevents automation from running at suspicious hours (e.g., midnight)
func WaitUntilBusinessHours(startHour int) {

	for {
		// Check if current time is within allowed working hours
		if IsWithinBusinessHours(startHour, startHour+8) {
			return
		}

		// -------------------------------------------------
		// RANDOMIZED WAIT
		// -------------------------------------------------
		// Sleep for a random interval (30–60 minutes)
		// before checking again to avoid predictable behavior
		time.Sleep(time.Duration(rand.Intn(30)+30) * time.Minute)
	}
}

// RandomBreak simulates natural human breaks during activity
// Helps avoid continuous, bot-like execution patterns
func RandomBreak() {

	// 20% probability of taking a break
	if rand.Float64() < 0.2 {

		// Random break duration between 3–7 minutes
		breakTime := time.Duration(rand.Intn(5)+3) * time.Minute
		time.Sleep(breakTime)
	}
}
