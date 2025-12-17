package stealth

import (
	"math/rand"
	"time"
)

// HumanDelay sleeps for a random duration between minMs and maxMs
func HumanDelay(minMs, maxMs int) {
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}

	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
