package stealth

import (
	"math/rand"
	"time"
)

// IsWithinBusinessHours checks if current time is within allowed activity window
func IsWithinBusinessHours(startHour, endHour int) bool {
	now := time.Now()
	hour := now.Hour()

	return hour >= startHour && hour < endHour
}

// WaitUntilBusinessHours pauses execution until business hours start
func WaitUntilBusinessHours(startHour int) {
	for {
		if IsWithinBusinessHours(startHour, startHour+8) {
			return
		}

		// Sleep for a random interval before checking again
		time.Sleep(time.Duration(rand.Intn(30)+30) * time.Minute)
	}
}

// RandomBreak simulates human breaks during activity
func RandomBreak() {
	// 20% chance to take a break
	if rand.Float64() < 0.2 {
		breakTime := time.Duration(rand.Intn(5)+3) * time.Minute
		time.Sleep(breakTime)
	}
}
