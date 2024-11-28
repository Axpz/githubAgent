package callback

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"k8s.io/apiserver/pkg/server"

	"agent/api/auth"
)

var (
	clientID     = auth.GithubAuthConfig.ClientID     // GitHub OAuth client ID
	clientSecret = auth.GithubAuthConfig.ClientSecret // GitHub OAuth client secret
	redirectURI  = auth.GithubAuthConfig.RedirectURI  // GitHub OAuth callback URL
	githubAPI    = auth.GithubAuthConfig.Issuer       // GitHub API URL
)

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		stopCh := server.SetupSignalHandler()
		<-stopCh
		cancel()
	}()

	state, err := r.Cookie("state")
	if err != nil {
		http.Error(w, "state not found", http.StatusBadRequest)
		return
	}
	if r.URL.Query().Get("state") != state.Value {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  redirectURI,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// Exchange the code for access_token
	oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("oauth2Token: %v", oauth2Token)

	req, err := http.NewRequest("GET", githubAPI+"/user", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create request: %v", err), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+oauth2Token.AccessToken)

	// Make the request to GitHub API to fetch the user info
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Errorf("failed to send request: %v", err).Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Errorf("failed to decode response: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	// oauth2Token.AccessToken = "*REDACTED*"

	resp2 := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		GithubUser    map[string]interface{}
	}{oauth2Token, new(json.RawMessage), user}

	data, err := json.MarshalIndent(resp2, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
