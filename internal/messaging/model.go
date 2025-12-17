package messaging

// MessageRecord tracks message state
type MessageRecord struct {
	ProfileURL string `json:"profile_url"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Sent       bool   `json:"sent"`
}
