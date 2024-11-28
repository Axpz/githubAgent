package main

import (
	"log"
	"net/http"

	"agent/api/auth"
	"agent/api/auth/callback"
)

func main() {

	http.HandleFunc("/api/auth", auth.AuthHandler)

	// Set up the callback URL handler
	http.HandleFunc("/api/auth/callback/github", callback.GithubCallbackHandler)

	http.HandleFunc("/api/auth/callback/google", callback.GoogleCallbackHandler)

	// Start the HTTP server
	port := ":8080"
	log.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
