package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// MoveMouseHuman simulates human-like mouse movement using Bezier curves
// via browser mousemove events (version-safe)
func MoveMouseHuman(page *rod.Page, startX, startY, endX, endY float64) {

	// Random control point
	controlX := (startX+endX)/2 + rand.Float64()*120 - 60
	controlY := (startY+endY)/2 + rand.Float64()*120 - 60

	steps := rand.Intn(30) + 30 // 30–60 steps

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)

		// Quadratic Bezier curve
		x := math.Pow(1-t, 2)*startX +
			2*(1-t)*t*controlX +
			math.Pow(t, 2)*endX

		y := math.Pow(1-t, 2)*startY +
			2*(1-t)*t*controlY +
			math.Pow(t, 2)*endY

		// Dispatch mousemove event in browser
		page.MustEval(`(x, y) => {
			document.dispatchEvent(
				new MouseEvent('mousemove', {
					clientX: x,
					clientY: y,
					bubbles: true
				})
			);
		}`, x, y)

		time.Sleep(time.Duration(rand.Intn(18)+6) * time.Millisecond)
	}
}
