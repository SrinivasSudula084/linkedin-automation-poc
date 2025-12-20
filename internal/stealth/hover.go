package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HoverElement simulates a realistic mouse hover over a DOM element
// It calculates a random point inside the element’s bounding box
func HoverElement(el *rod.Element) {

	// -------------------------------------------------
	// FETCH ELEMENT POSITION
	// -------------------------------------------------
	// Get element bounding box using JavaScript
	box := el.MustEval(`() => {
		const r = this.getBoundingClientRect();
		return {
			x: r.x,
			y: r.y,
			width: r.width,
			height: r.height
		};
	}`).Map()

	// Choose a random point inside the element
	x := box["x"].Num() + rand.Float64()*box["width"].Num()
	y := box["y"].Num() + rand.Float64()*box["height"].Num()

	// -------------------------------------------------
	// SIMULATE MOUSE MOVE EVENT
	// -------------------------------------------------
	// Dispatch a mousemove event to imitate real user behavior
	el.Page().MustEval(`(x, y) => {
		document.dispatchEvent(
			new MouseEvent('mousemove', {
				clientX: x,
				clientY: y,
				bubbles: true
			})
		);
	}`, x, y)

	// Add a small human-like pause after hover
	HumanDelay(300, 800)
}

// IdleMouseWander simulates small, random cursor movements
// Used when the user is idle to appear more human-like
func IdleMouseWander(page *rod.Page) {

	// Randomize number of idle movements (2–4)
	wanders := rand.Intn(3) + 2

	for i := 0; i < wanders; i++ {

		// Generate random screen coordinates
		x := rand.Float64()*800 + 100
		y := rand.Float64()*500 + 100

		// Dispatch mouse movement event
		page.MustEval(`(x, y) => {
			document.dispatchEvent(
				new MouseEvent('mousemove', {
					clientX: x,
					clientY: y,
					bubbles: true
				})
			);
		}`, x, y)

		// Pause between movements to mimic natural idle behavior
		time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
	}
}
