package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	state, err := randString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	nonce, err := randString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	setCallbackCookie(w, r, "state", state)
	setCallbackCookie(w, r, "nonce", nonce)

	authURL := "https://github.com/login/oauth/authorize"
	githubAuthURL := authURL + fmt.Sprintf("?client_id=%s&state=", GithubAuthConfig.ClientID) + state

	googleAuthURL := "https://accounts.google.com/o/oauth2/auth?client_id=YOUR_GOOGLE_CLIENT_ID&response_type=code&scope=openid%20email&redirect_uri=YOUR_REDIRECT_URI&state=" + state + "&nonce=" + nonce

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		<html>
			<head>
				<title>Login</title>
			</head>
			<body>
				<a href="` + githubAuthURL + `">Login with GitHub</a><br>
				<a href="` + googleAuthURL + `">Login with Google</a>
			</body>
		</html>
	`))
}
