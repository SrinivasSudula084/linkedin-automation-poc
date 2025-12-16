package main

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	log.Println("Starting LinkedIn Automation POC")

	l := launcher.New().
		Headless(false).
		Leakless(false)

	url := l.MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()

	page := browser.MustPage("https://example.com")

	log.Println("Browser opened and page loaded")

	// Future automation logic will go here

	page.MustClose()
	browser.MustClose()

	log.Println("Browser closed. Application finished.")
}
