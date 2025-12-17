package stealth

import "github.com/go-rod/rod"

// ApplyFingerprintMask applies browser fingerprint masking
func ApplyFingerprintMask(page *rod.Page) {
	// Remove webdriver flag
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});
	}`)

	// Fake plugins length
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'plugins', {
			get: () => [1, 2, 3, 4, 5]
		});
	}`)

	// Fake languages
	page.MustEval(`() => {
		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});
	}`)
}
