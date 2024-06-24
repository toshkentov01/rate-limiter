# rate-limiter

## Rate Limiter in Go

## This repository contains an implementation of a rate limiter using the token bucket algorithm in Go.

## Overview:
    A rate limiter is used to control the rate of traffic sent or received by a network interface controller. It can be essential in maintaining the performance and stability of applications, especially in a distributed environment or when dealing with third-party services.

## Token Bucket Algorithm:
    The token bucket algorithm is a widely used rate-limiting algorithm. It works as follows:
        1. A bucket is used to hold tokens. Tokens represent the capacity to process requests.
        2. Tokens are added to the bucket at a fixed rate (up to the maximum bucket size).
        3. When a request is made, a token is removed from the bucket. If there are no tokens available, the request is denied or delayed until a token becomes available.
        4. The bucket has a maximum capacity, meaning it cannot hold more than a fixed number of tokens.
        5. This approach allows for a flexible and efficient rate-limiting mechanism.


## Requirements
    - Go 1.21 or higher

## Instructions

1. Clone the repository
    ```
    git clone https://github.com/toshkentov01/rate-limiter.git
    cd ratelimiter
    ```

2. Run the application
    ```
    go run main.go
    ```

3. Run tests
    ```
    go test -v
    ```

## Code Example
    ```go
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
```