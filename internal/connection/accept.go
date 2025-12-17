package connection

import "log"

func SimulateAcceptedConnections() {
	sent, _ := loadState("sent_requests.json")
	if len(sent) == 0 {
		return
	}

	connected, _ := loadState("connected_profiles.json")

	// Simulate acceptance of first sent profile
	profile := sent[0]
	connected = append(connected, profile)

	_ = saveState("connected_profiles.json", connected)
	_ = saveState("sent_requests.json", sent[1:])

	log.Println("[STATE] Profile accepted:", profile.Name)
}
