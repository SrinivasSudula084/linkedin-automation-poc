package state

type SentRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ConnectedProfile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
