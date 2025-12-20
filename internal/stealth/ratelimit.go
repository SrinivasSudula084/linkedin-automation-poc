package stealth

import (
	"errors"
	"time"
)

// RateLimiter controls how frequently actions are performed
// This helps prevent aggressive automation behavior
type RateLimiter struct {

	// Maximum number of actions allowed per day
	DailyLimit int

	// Number of actions already used
	Used int

	// Timestamp of the last performed action
	LastAction time.Time

	// Minimum delay required between two actions
	Cooldown time.Duration
}

// NewRateLimiter creates and initializes a rate limiter
// dailyLimit → max allowed actions per day
// cooldownSeconds → delay between actions
func NewRateLimiter(dailyLimit int, cooldownSeconds int) *RateLimiter {
	return &RateLimiter{
		DailyLimit: dailyLimit,
		Used:       0,
		Cooldown:   time.Duration(cooldownSeconds) * time.Second,
	}
}

// Check verifies whether a new action can be performed
// It enforces both daily limits and cooldown timing
func (r *RateLimiter) Check() error {

	// -------------------------------------------------
	// DAILY LIMIT ENFORCEMENT
	// -------------------------------------------------
	if r.Used >= r.DailyLimit {
		return errors.New("daily limit reached")
	}

	// -------------------------------------------------
	// COOLDOWN ENFORCEMENT
	// -------------------------------------------------
	// Ensure enough time has passed since last action
	if !r.LastAction.IsZero() {
		elapsed := time.Since(r.LastAction)
		if elapsed < r.Cooldown {
			time.Sleep(r.Cooldown - elapsed)
		}
	}

	// Record the action
	r.Used++
	r.LastAction = time.Now()

	return nil
}
