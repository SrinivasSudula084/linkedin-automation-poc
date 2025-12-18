package main

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/config"
	"linkedin-automation-poc/internal/stealth"
	
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/connection"
	"linkedin-automation-poc/internal/messaging"
)

func main() {
	log.Println("Starting LinkedIn Automation POC (DEMO MODE)")

	// -------------------------------
	// Browser (POC only)
	// -------------------------------
	l := launcher.New().Headless(false).Leakless(false)
	url := l.MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage()

	stealth.ApplyFingerprintMask(page)
	cfg := config.Load()

	// -------------------------------
	// COOKIES: LOAD SESSION
	// -------------------------------
	loaded, err := LoadCookies(browser)
	if err != nil {
		log.Println("[COOKIES] Failed to load cookies:", err)
	}

	page.MustNavigate("https://www.linkedin.com")
	page.MustWaitLoad()

	if loaded {
		log.Println("[COOKIES] Cookies loaded, checking session...")
		if page.MustHas("nav.global-nav") {
			log.Println("[AUTH] Logged in via cookies")
			goto AFTER_LOGIN
		}
	}

	// -------------------------------
	// LOGIN (DEMONSTRATIVE ONLY)
	// -------------------------------
	if err := auth.Login(page, cfg.LinkedInEmail, cfg.LinkedInPassword); err != nil {
		log.Println("[AUTH] Login handled (demo):", err)
	}

	// -------------------------------
	// SAVE COOKIES AFTER LOGIN
	// -------------------------------
	if err := SaveCookies(page); err != nil {
		log.Println("[COOKIES] Failed to save cookies:", err)
	} else {
		log.Println("[COOKIES] Cookies saved successfully")
	}

AFTER_LOGIN:

	// ===============================
	// STEP 1: SEARCH & TARGETING
	// ===============================
	criteria := search.SearchCriteria{
		JobTitle: "golang",
		Location: "india",
		Keywords: "backend",
	}

	log.Println("[SEARCH] Running search with criteria:", criteria)

	allProfiles, err := search.LoadProfilesFromJSON("demo_profiles.json")
	if err != nil {
		log.Fatal("[SEARCH] Failed to load profiles:", err)
	}

	seenSearch := make(map[string]bool)

	// -------- PAGINATION START --------
	pageSize := 3
	pageNumber := 1

	results := []search.SearchResult{}

	for i := 0; i < len(allProfiles); i += pageSize {
		end := i + pageSize
		if end > len(allProfiles) {
			end = len(allProfiles)
		}

		log.Printf("[SEARCH][PAGE %d] Processing profiles %d to %d\n",
			pageNumber, i+1, end)

		pageProfiles := allProfiles[i:end]

		pageResults := search.SearchProfiles(pageProfiles, criteria, seenSearch)
		results = append(results, pageResults...)

		log.Printf("[SEARCH][PAGE %d] %d matches found\n",
			pageNumber, len(pageResults))

		pageNumber++
	}
	// -------- PAGINATION END --------

	log.Printf("[SEARCH] Total unique matching profiles found: %d\n", len(results))

	// ===============================
	// STEP 2: CONNECTION REQUESTS
	// ===============================
	sentToday := 0
	sentTracker := make(map[string]bool)

	for _, r := range results {
		profile := connection.Profile{
			ID:       r.ID,
			Name:     r.Name,
			Headline: r.Headline,
			URL:      r.URL,
		}

		if err := connection.SendConnection(profile, &sentToday, sentTracker); err != nil {
			log.Println("[CONNECT] Stopped:", err)
			break
		}
	}

	// ===============================
	// STEP 3: ACCEPTED CONNECTIONS
	// ===============================
	if err := connection.SimulateAcceptedConnections(); err != nil {
		log.Println("[STATE] Accept simulation error:", err)
	}

	connected, err := connection.LoadConnectedProfiles()
	if err != nil {
		log.Fatal("[STATE] Failed to load connected profiles:", err)
	}

	log.Printf("[STATE] %d profiles accepted connection\n", len(connected))

	// ===============================
	// STEP 4: MESSAGING SYSTEM
	// ===============================
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

	if err := messaging.SaveMessageHistory("message_history.json", history); err != nil {
		log.Println("[MESSAGE] Failed to save history:", err)
	}

	// -------------------------------
	// Cleanup
	// -------------------------------
	page.MustClose()
	browser.MustClose()

	log.Println("Browser closed. Application finished.")
}
