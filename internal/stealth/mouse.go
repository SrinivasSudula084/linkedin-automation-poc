package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// MoveMouseHuman simulates human-like mouse movement
// It uses a quadratic Bézier curve instead of straight lines
// to avoid robotic, bot-like behavior
func MoveMouseHuman(page *rod.Page, startX, startY, endX, endY float64) {

	// -------------------------------------------------
	// RANDOM CONTROL POINT
	// -------------------------------------------------
	// Control point adds curvature and randomness
	// Humans rarely move the mouse in a straight line
	controlX := (startX+endX)/2 + rand.Float64()*120 - 60
	controlY := (startY+endY)/2 + rand.Float64()*120 - 60

	// Randomize number of movement steps (smoothness)
	steps := rand.Intn(30) + 30 // 30–60 steps

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		// -------------------------------------------------
		// QUADRATIC BÉZIER CURVE CALCULATION
		// -------------------------------------------------
		// Generates smooth, natural mouse movement
		x := math.Pow(1-t, 2)*startX +
			2*(1-t)*t*controlX +
			math.Pow(t, 2)*endX

		y := math.Pow(1-t, 2)*startY +
			2*(1-t)*t*controlY +
			math.Pow(t, 2)*endY

		// -------------------------------------------------
		// DISPATCH MOUSE MOVE EVENT
		// -------------------------------------------------
		// Send mousemove event to the browser
		page.MustEval(`(x, y) => {
			document.dispatchEvent(
				new MouseEvent('mousemove', {
					clientX: x,
					clientY: y,
					bubbles: true
				})
			);
		}`, x, y)

		// Small randomized delay between movements
		time.Sleep(time.Duration(rand.Intn(18)+6) * time.Millisecond)
	}
}
