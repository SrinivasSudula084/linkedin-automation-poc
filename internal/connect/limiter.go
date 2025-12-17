package connect

import "errors"

type RateLimiter struct {
	DailyLimit int
	SentToday  int
}

func NewRateLimiter(limit int) *RateLimiter {
	return &RateLimiter{
		DailyLimit: limit,
	}
}

func (r *RateLimiter) Allow() error {
	if r.SentToday >= r.DailyLimit {
		return errors.New("daily connection limit reached")
	}
	r.SentToday++
	return nil
}
