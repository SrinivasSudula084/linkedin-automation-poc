package stealth

import "github.com/go-rod/rod"

// ApplyFingerprintMask modifies browser properties
// to reduce automation detection by websites like LinkedIn
func ApplyFingerprintMask(page *rod.Page) {

	// -------------------------------------------------
	// HIDE AUTOMATION FLAG
	// -------------------------------------------------
	// Removes the `navigator.webdriver` flag
	// which is commonly used to detect bots
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
	}`)

	// -------------------------------------------------
	// FAKE BROWSER PLUGINS
	// -------------------------------------------------
	// Normal browsers have installed plugins
	// Bots often expose an empty plugins list
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'plugins', {
			get: () => [1, 2, 3, 4, 5]
		});
	}`)

	// -------------------------------------------------
	// FAKE LANGUAGE SETTINGS
	// -------------------------------------------------
	// Sets realistic language preferences
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});
	}`)
}
