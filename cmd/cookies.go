package main

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const cookieFile = "cookies.json"

// SaveCookies saves browser cookies to a file
func SaveCookies(page *rod.Page) error {
	cookies, err := page.Browser().GetCookies()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		return err
	}

	return os.WriteFile(cookieFile, data, 0644)
}

// LoadCookies loads cookies from file into browser
func LoadCookies(browser *rod.Browser) (bool, error) {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return false, nil // cookies not found
	}

	var storedCookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return false, err
	}

	// 🔁 CONVERT cookie type
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

	if err := browser.SetCookies(params); err != nil {
		return false, err
	}

	return true, nil
}
