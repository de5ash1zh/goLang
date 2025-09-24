package main
package main

import (
	"fmt"
	"log"
	"time"
)

// Custom error type
type ValidationError struct {
	Field string
	Error string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Error)
}

// Function decorator (middleware) example
type HttpHandler func(string) (string, error)

func withLogging(handler HttpHandler) HttpHandler {
	return func(req string) (string, error) {
		start := time.Now()
		result, err := handler(req)
		duration := time.Since(start)
		log.Printf("Request: %s, Duration: %v, Error: %v\n", req, duration, err)
		return result, err
	}
}

func withRetry(attempts int, handler HttpHandler) HttpHandler {
	return func(req string) (string, error) {
		var lastErr error
		for i := 0; i < attempts; i++ {
			result, err := handler(req)
			if err == nil {
				return result, nil
			}
			lastErr = err
			log.Printf("Attempt %d failed: %v\n", i+1, err)
			time.Sleep(time.Millisecond * 100) // Simple backoff
		}
		return "", fmt.Errorf("all %d attempts failed. Last error: %v", attempts, lastErr)
	}
}

// Example handler
func handleRequest(req string) (string, error) {
	// Simulate random failure
	if time.Now().UnixNano()%2 == 0 {
		return "", fmt.Errorf("random error")
	}
	return fmt.Sprintf("processed: %s", req), nil
}

// Function that returns multiple functions
func mathOperations() (func(int) int, func(int) int) {
	// Closure capturing shared variable
	base := 10
	
	square := func(n int) int {
		return n * n
	}
	
	addBase := func(n int) int {
		return n + base
	}
	
	return square, addBase
}

// Generic function example
type Number interface {
	int | float64
}

func sum[T Number](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

// Builder pattern with functional options
type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

type ServerOption func(*Server)

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int) ServerOption {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func NewServer(host string, options ...ServerOption) *Server {
	server := &Server{
		host:    host,
		port:    8080,    // default
		timeout: 30 * time.Second, // default
		maxConn: 100,     // default
	}
	
	for _, option := range options {
		option(server)
	}
	
	return server
}

func main() {
	// Example 1: Function decorators
	handler := handleRequest
	loggedHandler := withLogging(handler)
	retryingLoggedHandler := withRetry(3, loggedHandler)

	fmt.Println("Testing decorated handler:")
	result, err := retryingLoggedHandler("test-request")
	fmt.Printf("Final result: %v, err: %v\n", result, err)

	// Example 2: Multiple function returns and closures
	square, addBase := mathOperations()
	fmt.Printf("\nSquare of 5: %d\n", square(5))
	fmt.Printf("5 + base: %d\n", addBase(5))

	// Example 3: Generic function
	fmt.Printf("\nSum of integers: %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("Sum of floats: %.2f\n", sum(1.1, 2.2, 3.3))

	// Example 4: Builder pattern with functional options
	server := NewServer(
		"localhost",
		WithPort(9000),
		WithTimeout(1*time.Minute),
		WithMaxConn(1000),
	)
	
	fmt.Printf("\nServer configuration:\nHost: %s\nPort: %d\nTimeout: %v\nMax Connections: %d\n",
		server.host, server.port, server.timeout, server.maxConn)

	// Example 5: Error handling
	if err := validateUser("", 15); err != nil {
		fmt.Printf("\nValidation error: %v\n", err)
	}
}

func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{
			Field: "name",
			Error: "cannot be empty",
		}
	}
	if age < 18 {
		return &ValidationError{
			Field: "age",
			Error: "must be 18 or older",
		}
	}
	return nil
}