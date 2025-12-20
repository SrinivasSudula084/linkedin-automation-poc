package main

import (
	"log"

	// Rod is used to control the browser (Chromium)
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	// Internal application modules
	"linkedin-automation-poc/internal/auth"
	"linkedin-automation-poc/internal/config"
	"linkedin-automation-poc/internal/stealth"
	
	"linkedin-automation-poc/internal/search"
	"linkedin-automation-poc/internal/connection"
	"linkedin-automation-poc/internal/messaging"
)

/*
main() is the starting point of the application.
It orchestrates the full automation flow:
1. Browser setup
2. Authentication (cookies or login)
3. Search & targeting
4. Connection requests
5. Accepted connections simulation
6. Messaging
*/
func main() {
	log.Println("Starting LinkedIn Automation POC (DEMO MODE)")

	// ===============================
	// BROWSER SETUP (POC / DEMO)
	// ===============================
	// Launch a visible (non-headless) browser for demonstration
	l := launcher.New().
		Headless(false).   // false => browser UI is visible
		Leakless(false)   // disabled for simplicity in POC

	url := l.MustLaunch()

	// Connect Rod to the launched browser instance
	browser := rod.New().ControlURL(url).MustConnect()

	// Open a new page (tab)
	page := browser.MustPage()

	// Apply stealth techniques to reduce bot detection
	stealth.ApplyFingerprintMask(page)

	// Load environment variables (.env)
	cfg := config.Load()

	// ===============================
	// COOKIES: LOAD EXISTING SESSION
	// ===============================
	// Attempt to reuse a previous login session using cookies
	loaded, err := LoadCookies(browser)
	if err != nil {
		log.Println("[COOKIES] Failed to load cookies:", err)
	}

	// Navigate to LinkedIn homepage
	page.MustNavigate("https://www.linkedin.com")
	page.MustWaitLoad()

	// If cookies were loaded, verify if login is still valid
	if loaded {
		log.Println("[COOKIES] Cookies loaded, checking session...")

		// Presence of LinkedIn navigation bar confirms logged-in state
		if page.MustHas("nav.global-nav") {
			log.Println("[AUTH] Logged in via cookies")
			goto AFTER_LOGIN // Skip login step
		}
	}

	// ===============================
	// LOGIN (DEMONSTRATION ONLY)
	// ===============================
	// Perform login if cookies are missing or expired
	if err := auth.Login(page, cfg.LinkedInEmail, cfg.LinkedInPassword); err != nil {
		log.Println("[AUTH] Login handled (demo):", err)
	}

	// ===============================
	// SAVE COOKIES AFTER LOGIN
	// ===============================
	// Persist cookies to avoid logging in again next run
	if err := SaveCookies(page); err != nil {
		log.Println("[COOKIES] Failed to save cookies:", err)
	} else {
		log.Println("[COOKIES] Cookies saved successfully")
	}

AFTER_LOGIN:

	// ===============================
	// STEP 1: SEARCH & TARGETING
	// ===============================
	// Define search criteria (used for filtering profiles)
	criteria := search.SearchCriteria{
		JobTitle: "golang",
		Location: "india",
		Keywords: "backend",
	}

	log.Println("[SEARCH] Running search with criteria:", criteria)

	// Load demo profiles from JSON (POC — no real scraping)
	allProfiles, err := search.LoadProfilesFromJSON("demo_profiles.json")
	if err != nil {
		log.Fatal("[SEARCH] Failed to load profiles:", err)
	}

	// Track already-seen profiles to avoid duplicates
	seenSearch := make(map[string]bool)

	// -------- PAGINATION SIMULATION --------
	pageSize := 3     // Number of profiles per "page"
	pageNumber := 1   // Current page index

	results := []search.SearchResult{}

	// Simulate pagination logic
	for i := 0; i < len(allProfiles); i += pageSize {
		end := i + pageSize
		if end > len(allProfiles) {
			end = len(allProfiles)
		}

		log.Printf("[SEARCH][PAGE %d] Processing profiles %d to %d\n",
			pageNumber, i+1, end)

		// Current page profiles
		pageProfiles := allProfiles[i:end]

		// Apply filtering based on criteria
		pageResults := search.SearchProfiles(pageProfiles, criteria, seenSearch)

		// Append results to global result list
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
	sentToday := 0                    // Daily request counter
	sentTracker := make(map[string]bool) // Prevent duplicate sends

	for _, r := range results {
		// Convert search result to connection profile format
		profile := connection.Profile{
			ID:       r.ID,
			Name:     r.Name,
			Headline: r.Headline,
			URL:      r.URL,
		}

		// Send connection request with daily limit enforcement
		if err := connection.SendConnection(profile, &sentToday, sentTracker); err != nil {
			log.Println("[CONNECT] Stopped:", err)
			break
		}
	}

	// ===============================
	// STEP 3: ACCEPTED CONNECTIONS
	// ===============================
	// Simulate that some sent requests are accepted (POC behavior)
	if err := connection.SimulateAcceptedConnections(); err != nil {
		log.Println("[STATE] Accept simulation error:", err)
	}

	// Load accepted connections from persistent state
	connected, err := connection.LoadConnectedProfiles()
	if err != nil {
		log.Fatal("[STATE] Failed to load connected profiles:", err)
	}

	log.Printf("[STATE] %d profiles accepted connection\n", len(connected))

	// ===============================
	// STEP 4: MESSAGING SYSTEM
	// ===============================
	// Message template with placeholder for personalization
	template := "Hi {{name}}, thanks for accepting my connection. Let's stay in touch!"

	history := []messaging.MessageRecord{}

	for _, p := range connected {
		// Render personalized message
		msg := messaging.RenderTemplate(template, p.Name)

		record := messaging.MessageRecord{
			ProfileURL: p.URL,
			Name:       p.Name,
			Message:    msg,
		}

		// Send follow-up message
		if err := messaging.SendMessage(&record); err != nil {
			log.Println("[MESSAGE] Failed:", err)
			continue
		}

		history = append(history, record)
		log.Println("[MESSAGE] Sent follow-up to:", p.Name)
	}

	// Persist full message history
	if err := messaging.SaveMessageHistory("message_history.json", history); err != nil {
		log.Println("[MESSAGE] Failed to save history:", err)
	}

	// ===============================
	// CLEANUP
	// ===============================
	// Close browser resources gracefully
	page.MustClose()
	browser.MustClose()

	log.Println("Browser closed. Application finished.")
}
