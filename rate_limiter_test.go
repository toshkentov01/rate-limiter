// rate_limiter_test.go
package ratelimiter

import (
	"testing"
	"time"
)

// TestRateLimiter tests the rate limiter functionality.
func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(3, time.Second) // Allow 3 requests per second

	// Test with multiple requests within the rate limit
	for i := 0; i < 3; i++ {
		allowed := rl.AllowRequest("user1")
		if !allowed {
			t.Errorf("Expected request %d to be allowed", i+1)
		}
	}

	// Test with requests exceeding the rate limit
	for i := 0; i < 3; i++ {
		allowed := rl.AllowRequest("user1")
		if allowed {
			t.Errorf("Expected request %d to be blocked", i+4)
		}
	}

	// Wait for the rate limiter to reset
	time.Sleep(time.Second)

	// After waiting for a second, requests should be allowed again
	for i := 0; i < 3; i++ {
		allowed := rl.AllowRequest("user1")
		if !allowed {
			t.Errorf("Expected request %d to be allowed after reset", i+1)
		}
	}
}
