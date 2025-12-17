package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// RandomScroll simulates human-like scrolling behavior
func RandomScroll(page *rod.Page) {

	scrollSteps := rand.Intn(5) + 5 // 5–10 scroll actions

	for i := 0; i < scrollSteps; i++ {

		// Scroll down by a random amount
		scrollBy := rand.Intn(400) + 200 // 200–600px

		page.MustEval(`(y) => {
			window.scrollBy({ top: y, behavior: 'smooth' });
		}`, scrollBy)

		// Pause like reading
		HumanDelay(800, 2000)

		// Occasionally scroll back up (human correction)
		if rand.Float64() < 0.25 {
			page.MustEval(`() => {
				window.scrollBy({ top: -150, behavior: 'smooth' });
			}`)
			HumanDelay(500, 1200)
		}
	}

	// Final pause
	time.Sleep(time.Duration(rand.Intn(1500)+1000) * time.Millisecond)
}
