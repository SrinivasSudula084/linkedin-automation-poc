package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HoverElement simulates hovering over an element using DOM bounding box
func HoverElement(el *rod.Element) {

	// Get element bounding box via JS
	box := el.MustEval(`() => {
		const r = this.getBoundingClientRect();
		return {
			x: r.x,
			y: r.y,
			width: r.width,
			height: r.height
		};
	}`).Map()

	x := box["x"].Num() + rand.Float64()*box["width"].Num()
	y := box["y"].Num() + rand.Float64()*box["height"].Num()

	el.Page().MustEval(`(x, y) => {
		document.dispatchEvent(
			new MouseEvent('mousemove', {
				clientX: x,
				clientY: y,
				bubbles: true
			})
		);
	}`, x, y)

	HumanDelay(300, 800)
}

// IdleMouseWander simulates small idle cursor movements
func IdleMouseWander(page *rod.Page) {
	wanders := rand.Intn(3) + 2 // 2–4 movements

	for i := 0; i < wanders; i++ {
		x := rand.Float64()*800 + 100
		y := rand.Float64()*500 + 100

		page.MustEval(`(x, y) => {
			document.dispatchEvent(
				new MouseEvent('mousemove', {
					clientX: x,
					clientY: y,
					bubbles: true
				})
			);
		}`, x, y)

		time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
	}
}
