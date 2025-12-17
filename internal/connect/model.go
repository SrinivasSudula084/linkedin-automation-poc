package connect

import "time"

type ConnectionRequest struct {
	ProfileURL string
	Note       string
	SentAt     time.Time
}
