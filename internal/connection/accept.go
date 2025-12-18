package connection

import "log"

func SimulateAcceptedConnections() error {
	sent, _ := loadState(SentFile)
	existing, _ := loadState(ConnectedFile)

	unique := map[string]Profile{}

	// load existing connections safely
	for _, p := range existing {
		unique[p.URL] = p
	}

	// accept new connections
	for _, p := range sent {
		if _, ok := unique[p.URL]; !ok {
			log.Println("[STATE] Profile accepted:", p.Name)
			unique[p.URL] = p
		}
	}

	// rebuild clean list
	clean := []Profile{}
	for _, p := range unique {
		clean = append(clean, p)
	}

	return saveState(ConnectedFile, clean)
}


func LoadConnectedProfiles() ([]Profile, error) {
	return loadState(ConnectedFile)
}

func SaveSentProfiles(profiles []Profile) error {
	return saveState(SentFile, profiles)
}
