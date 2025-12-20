package main

import (
	"encoding/json"
	"os"

	// Rod browser control
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// cookieFile defines where browser session cookies are stored locally
// This file allows session reuse without logging in every time
const cookieFile = "cookies.json"

// SaveCookies saves the current browser session cookies to a JSON file
// This helps persist login state across multiple runs of the program
func SaveCookies(page *rod.Page) error {
	// Fetch all cookies from the active browser session
	cookies, err := page.Browser().GetCookies()
	if err != nil {
		return err
	}

	// Convert cookies struct into JSON format
	data, err := json.Marshal(cookies)
	if err != nil {
		return err
	}

	// Write cookies to disk with read/write permissions
	// 0644 → owner can write, others can read
	return os.WriteFile(cookieFile, data, 0644)
}

// LoadCookies loads previously saved cookies into the browser
// Returns:
//   - true  → cookies were found and successfully loaded
//   - false → cookies do not exist or could not be used
func LoadCookies(browser *rod.Browser) (bool, error) {
	// Attempt to read cookies file from disk
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		// Cookies file does not exist (first run scenario)
		return false, nil
	}

	// Stored cookies as received from Chrome DevTools Protocol
	var storedCookies []*proto.NetworkCookie

	// Parse JSON cookies back into Go structures
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return false, err
	}

	// ------------------------------------------------
	// Cookie Type Conversion (Important)
	// ------------------------------------------------
	// Chrome expects cookies in NetworkCookieParam format
	// so we convert from NetworkCookie → NetworkCookieParam
	params := make([]*proto.NetworkCookieParam, 0, len(storedCookies))

	for _, c := range storedCookies {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	// Inject cookies back into the browser session
	if err := browser.SetCookies(params); err != nil {
		return false, err
	}

	// Cookies successfully restored
	return true, nil
}
