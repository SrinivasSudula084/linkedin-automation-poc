package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// HumanType types text into an input element
// It mimics real human typing behavior with per-character delays
// Suitable for sensitive inputs like email, password, and OTP fields
func HumanType(el *rod.Element, text string) {

	// -------------------------------------------------
	// ENSURE INPUT FOCUS
	// -------------------------------------------------
	// Click the element to bring it into focus
	el.MustClick()

	// Clear any existing value
	// Important for controlled inputs (e.g., React)
	el.MustEval(`() => { this.value = "" }`)

	// -------------------------------------------------
	// TYPE CHARACTER BY CHARACTER
	// -------------------------------------------------
	for _, char := range text {

		// Type one character at a time
		el.MustInput(string(char))

		// Add a slow, stable delay between keystrokes
		time.Sleep(time.Duration(rand.Intn(120)+120) * time.Millisecond)
	}

	// Small pause after typing finishes
	HumanDelay(300, 600)
}

// randomChar returns a random lowercase alphabet character
// (utility for future typing simulations or noise injection)
func randomChar() rune {

	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	return letters[rand.Intn(len(letters))]
}
