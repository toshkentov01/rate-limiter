package main

import (
	"fmt"
	"time"

	ratelimiter "github.com/toshkentov01/rate-limiter"
)

// Example usage
func main() {
	// Example usage of RateLimiter
	rl := ratelimiter.NewRateLimiter(3, time.Second) // Allow 3 requests per second

	// Simulate requests
	for i := 0; i < 10; i++ {
		allowed := rl.AllowRequest("user1") // AllowRequest takes as an input parametr identifier
		if allowed {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request blocked")
		}
		time.Sleep(200 * time.Millisecond) // Simulate some delay between requests
	}
}
