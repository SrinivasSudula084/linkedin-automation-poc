package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// RandomScroll simulates human-like scrolling behavior
// It adds randomness, pauses, and small corrections
func RandomScroll(page *rod.Page) {

	// -------------------------------------------------
	// DETERMINE NUMBER OF SCROLL ACTIONS
	// -------------------------------------------------
	// Humans usually scroll multiple times while reading
	scrollSteps := rand.Intn(5) + 5 // 5–10 scroll actions

	for i := 0; i < scrollSteps; i++ {

		// -------------------------------------------------
		// SCROLL DOWN
		// -------------------------------------------------
		// Scroll by a random vertical distance
		scrollBy := rand.Intn(400) + 200 // 200–600 pixels

		page.MustEval(`(y) => {
			window.scrollBy({ top: y, behavior: 'smooth' });
		}`, scrollBy)

		// Pause to simulate reading time
		HumanDelay(800, 2000)

		// -------------------------------------------------
		// OCCASIONAL SCROLL UP (HUMAN CORRECTION)
		// -------------------------------------------------
		// Humans sometimes scroll back slightly
		if rand.Float64() < 0.25 {
			page.MustEval(`() => {
				window.scrollBy({ top: -150, behavior: 'smooth' });
			}`)
			HumanDelay(500, 1200)
		}
	}

	// -------------------------------------------------
	// FINAL IDLE PAUSE
	// -------------------------------------------------
	// Simulates user stopping after reading
	time.Sleep(time.Duration(rand.Intn(1500)+1000) * time.Millisecond)
}
