package stealth

import (
	"math/rand"
	"time"
)

// HumanDelay pauses execution for a random duration
// This simulates natural human reaction and thinking time
func HumanDelay(minMs, maxMs int) {

	// -------------------------------------------------
	// SAFETY CHECK
	// -------------------------------------------------
	// If max is less than or equal to min, use min directly
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}

	// -------------------------------------------------
	// RANDOMIZED DELAY
	// -------------------------------------------------
	// Generate a random delay between minMs and maxMs
	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
