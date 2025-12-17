package stealth

import (
	"errors"
	"time"
)

type RateLimiter struct {
	DailyLimit int
	Used       int
	LastAction time.Time
	Cooldown   time.Duration
}

func NewRateLimiter(dailyLimit int, cooldownSeconds int) *RateLimiter {
	return &RateLimiter{
		DailyLimit: dailyLimit,
		Used:       0,
		Cooldown:   time.Duration(cooldownSeconds) * time.Second,
	}
}

// Check verifies if a new action is allowed
func (r *RateLimiter) Check() error {
	if r.Used >= r.DailyLimit {
		return errors.New("daily limit reached")
	}

	if !r.LastAction.IsZero() {
		elapsed := time.Since(r.LastAction)
		if elapsed < r.Cooldown {
			time.Sleep(r.Cooldown - elapsed)
		}
	}

	r.Used++
	r.LastAction = time.Now()
	return nil
}
