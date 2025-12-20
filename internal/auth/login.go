package auth

import (
	"errors"
	"log"
	"time"

	// Rod is used for browser page interaction
	"github.com/go-rod/rod"

	// Stealth utilities simulate human behavior
	"linkedin-automation-poc/internal/stealth"
)

// Login handles the LinkedIn login flow safely
// It simulates real human behavior and detects common failure scenarios
func Login(page *rod.Page, email, password string) error {
	log.Println("Navigating to LinkedIn login page")

	// Open LinkedIn login page
	page.MustNavigate("https://www.linkedin.com/login")

	// Add a random delay to mimic human reading/thinking time
	stealth.HumanDelay(2000, 4000)

	// -------------------------------------------------
	// EMAIL INPUT HANDLING
	// -------------------------------------------------
	// Wait up to 20 seconds for the email input field
	emailInput, err := page.Timeout(20 * time.Second).
		Element(`input[name="session_key"]`)
	if err != nil {
		// If email input is not found, check for security challenges
		if page.MustHas(`input[name="pin"]`) || page.MustHas(`iframe`) {
			// LinkedIn checkpoint or captcha detected
			return errors.New("login blocked: security checkpoint or captcha detected")
		}
		// Unexpected page structure
		return errors.New("login page did not load as expected")
	}

	// Type email slowly to simulate human typing
	stealth.HumanType(emailInput, email)

	// -------------------------------------------------
	// PASSWORD INPUT HANDLING
	// -------------------------------------------------
	passwordInput, err := page.Element(`input[name="session_password"]`)
	if err != nil {
		return errors.New("password input not found")
	}

	// Type password using human-like keystrokes
	stealth.HumanType(passwordInput, password)

	// Short delay before clicking login
	stealth.HumanDelay(800, 1500)

	// -------------------------------------------------
	// LOGIN BUTTON CLICK
	// -------------------------------------------------
	loginBtn, err := page.Element(`button[type="submit"]`)
	if err != nil {
		return errors.New("login button not found")
	}

	// Click the login button
	loginBtn.MustClick()

	// Wait for page navigation / response
	stealth.HumanDelay(4000, 6000)

	// -------------------------------------------------
	// LOGIN RESULT VALIDATION
	// -------------------------------------------------
	// If login page is still visible, credentials may be invalid
	if page.MustHas(`input[name="session_key"]`) {
		return errors.New("login failed: invalid credentials or redirected back to login")
	}

	// Detect LinkedIn checkpoint interruptions
	if page.MustHas(`checkpoint`) {
		return errors.New("login interrupted by LinkedIn checkpoint")
	}

	log.Println("Login attempt completed")
	return nil
}
