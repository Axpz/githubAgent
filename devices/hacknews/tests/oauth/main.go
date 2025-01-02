package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"

	"agent/api/auth"
	"agent/api/auth/callback"
)

// 200 requests per second with a maximum burst of 100 requests
var limiter = rate.NewLimiter(20, 100)

func rateLimitMiddleware(limiter *rate.Limiter) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Try to get a token, blocking time depends on the rate limiter configuration
			if !limiter.Allow() {
				http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Set up the authentication portal URL handler
	r.HandleFunc("/api/auth", auth.AuthHandler)

	// Set up the callback URL handler
	r.HandleFunc("/api/auth/callback/github", callback.GithubCallbackHandler)
	r.HandleFunc("/api/auth/callback/google", callback.GoogleCallbackHandler)

	// Apply the rate limit middleware to the entire mux
	r.Use(rateLimitMiddleware(limiter))

	// Start the HTTP server
	port := ":8080"
	log.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
