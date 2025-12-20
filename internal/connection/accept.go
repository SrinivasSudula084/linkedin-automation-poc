package connection

import "log"

// SimulateAcceptedConnections simulates LinkedIn accepting connection requests
// This is a POC-only function (no real LinkedIn API calls)
// It moves profiles from "sent" state to "connected" state safely
func SimulateAcceptedConnections() error {
	// Load profiles to whom connection requests were sent
	sent, _ := loadState(SentFile)

	// Load profiles that are already connected
	existing, _ := loadState(ConnectedFile)

	// Use a map to ensure uniqueness based on profile URL
	unique := map[string]Profile{}

	// -----------------------------------------
	// Load existing connected profiles safely
	// -----------------------------------------
	for _, p := range existing {
		unique[p.URL] = p
	}

	// -----------------------------------------
	// Simulate acceptance of new connections
	// -----------------------------------------
	for _, p := range sent {
		// Add only profiles that are not already connected
		if _, ok := unique[p.URL]; !ok {
			log.Println("[STATE] Profile accepted:", p.Name)
			unique[p.URL] = p
		}
	}

	// -----------------------------------------
	// Rebuild a clean, duplicate-free list
	// -----------------------------------------
	clean := []Profile{}
	for _, p := range unique {
		clean = append(clean, p)
	}

	// Save updated connected profiles state
	return saveState(ConnectedFile, clean)
}

// LoadConnectedProfiles returns all profiles that are marked as connected
func LoadConnectedProfiles() ([]Profile, error) {
	return loadState(ConnectedFile)
}

// SaveSentProfiles persists the list of sent connection requests
// This helps enforce daily limits and avoid duplicates
func SaveSentProfiles(profiles []Profile) error {
	return saveState(SentFile, profiles)
}
