package main

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/config"
	"linkedin-automation-poc/internal/stealth"
	"linkedin-automation-poc/internal/connection"
	"linkedin-automation-poc/internal/messaging"
)

func main() {
	log.Println("Starting LinkedIn Automation POC (DEMO MODE)")

	// ----------------------------------
	// Browser setup (POC only)
	// ----------------------------------
	l := launcher.New().
		Headless(false).
		Leakless(false)

	url := l.MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()

	page := browser.MustPage()

	// ----------------------------------
	// Apply stealth
	// ----------------------------------
	stealth.ApplyFingerprintMask(page)
	stealth.HumanDelay(1500, 3000)

	// ----------------------------------
	// Load config (env-based, demo only)
	// ----------------------------------
	cfg := config.Load()

	// ----------------------------------
	// Demonstrative login (NO dependency)
	// ----------------------------------
	if err := auth.Login(page, cfg.LinkedInEmail, cfg.LinkedInPassword); err != nil {
		log.Println("Login handled (expected in demo):", err)
	}

	// ==================================
	// DEMO MODE: CONNECTION REQUESTS
	// ==================================
	sentToday := 0
	connected := []connection.Profile{}

	profiles, err := connection.LoadDemoProfiles("demo_profiles.json")
	if err != nil {
		log.Println("Failed to load demo profiles:", err)
		return
	}

	for _, profile := range profiles {
		err := connection.SendConnection(profile, &sentToday)
		if err != nil {
			log.Println("Connection stopped:", err)
			break
		}
		connected = append(connected, profile)
	}

	// ==================================
	// DEMO MODE: MESSAGING SYSTEM
	// ONLY CONNECTED PROFILES
	// ==================================
	messageTemplate :=
		"Hi {{name}}, thanks for connecting! Looking forward to staying in touch."

	history := []messaging.MessageRecord{}

	for _, profile := range connected {
		msg := messaging.RenderTemplate(messageTemplate, profile.Name)

		record := messaging.MessageRecord{
			ProfileURL: profile.URL,
			Name:       profile.Name,
			Message:    msg,
		}

		if err := messaging.SendMessage(&record); err != nil {
			log.Println("Message failed:", err)
			continue
		}

		history = append(history, record)
	}

	if err := messaging.SaveMessageHistory("message_history.json", history); err != nil {
		log.Println("Failed to save message history:", err)
	}

	// ----------------------------------
	// Cleanup
	// ----------------------------------
	page.MustClose()
	browser.MustClose()

	log.Println("Browser closed. Application finished.")
}
