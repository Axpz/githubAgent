package auth

import "os"

type Config struct {
	ClientID     string // OAuth client ID
	ClientSecret string // OAuth client secret
	RedirectURI  string // OAuth callback URL
	Issuer       string // OAuth issuer API URL
}

var GithubAuthConfig = Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	RedirectURI:  "http://localhost:8080/api/auth/callback/github",
	Issuer:       "https://api.github.com",
}

var GoogleAuthConfig = Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	RedirectURI:  "http://localhost:8080/api/auth/callback/google",
	Issuer:       "https://accounts.google.com",
}
