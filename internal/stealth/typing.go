package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HumanType types text into an element with human-like behavior
// HumanTypeStable types text reliably for sensitive inputs
// (email, password, OTP fields)
func HumanType(el *rod.Element, text string) {
	// Ensure focus
	el.MustClick()

	// Clear existing value (important for React inputs)
	el.MustEval(`() => { this.value = "" }`)

	for _, char := range text {
		el.MustInput(string(char))

		// Slower, stable delay
		time.Sleep(time.Duration(rand.Intn(120)+120) * time.Millisecond)
	}

	// Small pause after typing
	HumanDelay(300, 600)
}

// randomChar returns a random lowercase character
func randomChar() rune {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	return letters[rand.Intn(len(letters))]
}
