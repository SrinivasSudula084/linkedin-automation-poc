package auth

import (
	"errors"
	"log"
	"time"

	"github.com/go-rod/rod"

	"linkedin-automation-poc/internal/stealth"
)

func Login(page *rod.Page, email, password string) error {
	log.Println("Navigating to LinkedIn login page")

	page.MustNavigate("https://www.linkedin.com/login")
	stealth.HumanDelay(2000, 4000)

	// Wait for email input safely
	emailInput, err := page.Timeout(20 * time.Second).
		Element(`input[name="session_key"]`)
	if err != nil {
		// Check for checkpoint / security page
		if page.MustHas(`input[name="pin"]`) || page.MustHas(`iframe`) {
			return errors.New("login blocked: security checkpoint or captcha detected")
		}
		return errors.New("login page did not load as expected")
	}

	stealth.HumanType(emailInput, email)

	passwordInput, err := page.
		Element(`input[name="session_password"]`)
	if err != nil {
		return errors.New("password input not found")
	}

	stealth.HumanType(passwordInput, password)
	stealth.HumanDelay(800, 1500)

	loginBtn, err := page.Element(`button[type="submit"]`)
	if err != nil {
		return errors.New("login button not found")
	}

	loginBtn.MustClick()
	stealth.HumanDelay(4000, 6000)

	// Detect login failure or checkpoint
	if page.MustHas(`input[name="session_key"]`) {
		return errors.New("login failed: invalid credentials or redirected back to login")
	}

	if page.MustHas(`checkpoint`) {
		return errors.New("login interrupted by LinkedIn checkpoint")
	}

	log.Println("Login attempt completed")
	return nil
}
