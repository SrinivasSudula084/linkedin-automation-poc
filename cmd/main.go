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
	// Browser setup (POC)
	// ----------------------------------
	l := launcher.New().Headless(false).Leakless(false)
	url := l.MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage()

	stealth.ApplyFingerprintMask(page)

	cfg := config.Load()

	// ----------------------------------
	// LOGIN (DEMONSTRATIVE)
	// ----------------------------------
	if err := auth.Login(page, cfg.LinkedInEmail, cfg.LinkedInPassword); err != nil {
		log.Println("[AUTH] Login handled (demo):", err)
	}

	// ==================================
	// STEP 1: SEARCH & TARGETING (DEMO)
	// ==================================
	log.Println("[SEARCH] Searching profiles with criteria:")
	log.Println("[SEARCH] JobTitle=Golang Developer | Location=India | Keywords=backend")

	searchResults, err := connection.LoadDemoProfiles("demo_profiles.json")
	if err != nil {
		log.Fatal("[SEARCH] Failed to load demo profiles:", err)
	}

	log.Printf("[SEARCH] %d profiles found\n", len(searchResults))
	for _, p := range searchResults {
		log.Printf("[SEARCH] Found profile: %s (%s)\n", p.Name, p.URL)
	}

	// ==================================
	// STEP 2: CONNECTION REQUESTS
	// ==================================
	sentToday := 0

	for _, p := range searchResults {
		err := connection.SendConnection(p, &sentToday)
		if err != nil {
			log.Println("[CONNECT] Stopped:", err)
			break
		}
	}

	// ==================================
	// STEP 3: ACCEPTED CONNECTIONS (STATE)
	// ==================================
	connection.SimulateAcceptedConnections()

	connected, err := connection.LoadConnectedProfiles("connected_profiles.json")
	if err != nil {
		log.Fatal("[STATE] Failed to load connected profiles:", err)
	}

	log.Printf("[STATE] %d profiles accepted connection\n", len(connected))

	// ==================================
	// STEP 4: MESSAGING SYSTEM
	// ==================================
	template := "Hi {{name}}, thanks for accepting my connection. Let's stay in touch!"

	history := []messaging.MessageRecord{}

	for _, p := range connected {
		msg := messaging.RenderTemplate(template, p.Name)

		record := messaging.MessageRecord{
			ProfileURL: p.URL,
			Name:       p.Name,
			Message:    msg,
		}

		if err := messaging.SendMessage(&record); err != nil {
			log.Println("[MESSAGE] Failed:", err)
			continue
		}

		history = append(history, record)
		log.Println("[MESSAGE] Sent follow-up to:", p.Name)
	}

	_ = messaging.SaveMessageHistory("message_history.json", history)

	// ----------------------------------
	// Cleanup
	// ----------------------------------
	page.MustClose()
	browser.MustClose()

	log.Println("Browser closed. Application finished.")
}
