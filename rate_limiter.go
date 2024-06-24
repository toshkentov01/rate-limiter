package ratelimiter

import (
	"sync"
	"time"
)

// RateLimiter defines a rate limiter struct
type RateLimiter struct {
	requests map[string][]time.Time // Map to store requests for each identifier
	mu       sync.Mutex             // Mutex for thread-safe access to requests map
	maxReq   int                    // Maximum requests allowed within interval
	interval time.Duration          // Time interval for rate limiting
}

// NewRateLimiter creates a new RateLimiter instance
func NewRateLimiter(maxReq int, interval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		maxReq:   maxReq,
		interval: interval,
	}
	go rl.clearExpiredRequests()
	return rl
}

// AllowRequest checks if a request is allowed based on rate limiting rules
func (rl *RateLimiter) AllowRequest(identifier string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if _, ok := rl.requests[identifier]; !ok {
		rl.requests[identifier] = make([]time.Time, 0)
	}

	// Remove expired requests
	rl.cleanupRequests(identifier)

	// Check if number of requests within interval is within limits
	if len(rl.requests[identifier]) < rl.maxReq {
		rl.requests[identifier] = append(rl.requests[identifier], time.Now())
		return true
	}

	return false
}

// cleanupRequests removes expired requests from the store
func (rl *RateLimiter) cleanupRequests(identifier string) {
	cutoff := time.Now().Add(-rl.interval)
	requests := rl.requests[identifier]
	var newRequests []time.Time
	for _, req := range requests {
		if req.After(cutoff) {
			newRequests = append(newRequests, req)
		}
	}
	rl.requests[identifier] = newRequests
}

// clearExpiredRequests periodically clears expired requests from the store
func (rl *RateLimiter) clearExpiredRequests() {
	ticker := time.NewTicker(rl.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		rl.mu.Lock()
		for identifier := range rl.requests {
			rl.cleanupRequests(identifier)
		}
		rl.mu.Unlock()
	}
}
