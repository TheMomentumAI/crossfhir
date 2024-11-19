package oauth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

func GetAuthToken(config *oauth2.Config, client *http.Client, grantType string) (*oauth2.Token, error) {
	if grantType == "client_credentials" {
		return getAuthTokenClientCredentials(config, client)
	} else if grantType == "authorization_code" {
		return nil, fmt.Errorf("authorization_code grant type not supported")
	} else {
		return nil, fmt.Errorf("unsupported grant type: %s", grantType)
	}
}

func getAuthTokenClientCredentials(config *oauth2.Config, client *http.Client) (*oauth2.Token, error) {
	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("scope", config.Scopes[0])
	formEncoded := form.Encode()

	req, err := http.NewRequest("POST", config.Endpoint.TokenURL, strings.NewReader(formEncoded))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}

	basicAuth := base64.StdEncoding.EncodeToString([]byte(config.ClientID + ":" + config.ClientSecret))
	req.Header.Set("Authorization", "Basic "+basicAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Log request for future verbose mode
	// log.Printf("Token request URL: %s", config.Endpoint.TokenURL)
	// log.Printf("Token request scope: %s", form.Get("scope"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute token request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read token response: %w", err)
	}

	var tokenRes struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}

	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	// Create oauth2.Token
	token := &oauth2.Token{
		AccessToken: tokenRes.AccessToken,
		TokenType:   tokenRes.TokenType,
		Expiry:      time.Now().Add(time.Duration(tokenRes.ExpiresIn) * time.Second),
	}

	return token, nil
}

// func getAuthTokenAuthCode(config *oauth2.Config, client *http.Client) (*oauth2.Token, error) {
// 	// Generate random state
// 	state := generateRandomState()

// 	// Generate authorization URL
// 	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOnline)

// 	// Print instructions for user
// 	fmt.Printf("Please visit this URL to authorize the application:\n%v\n", authURL)

// 	// Start local server to receive callback
// 	callbackChan := make(chan string)
// 	server := &http.Server{Addr: ":8080"}

// 	// Handle callback
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Verify state
// 		if r.URL.Query().Get("state") != state {
// 			http.Error(w, "Invalid state", http.StatusBadRequest)
// 			callbackChan <- ""
// 			return
// 		}

// 		// Get authorization code
// 		code := r.URL.Query().Get("code")
// 		if code == "" {
// 			http.Error(w, "No code received", http.StatusBadRequest)
// 			callbackChan <- ""
// 			return
// 		}

// 		// Display success message
// 		fmt.Fprintf(w, "Authorization successful! You can close this window.")

// 		// Send code to channel
// 		callbackChan <- code

// 		// Shutdown server after delay
// 		go func() {
// 			time.Sleep(1 * time.Second)
// 			server.Shutdown(context.Background())
// 		}()
// 	})

// 	// Start server in goroutine
// 	go func() {
// 		if err := server.ListenAndServe(); err != http.ErrServerClosed {
// 			log.Printf("HTTP server error: %v", err)
// 		}
// 	}()

// 	// Wait for callback
// 	code := <-callbackChan
// 	if code == "" {
// 		return nil, fmt.Errorf("failed to get authorization code")
// 	}

// 	// Exchange code for token
// 	ctx := context.Background()
// 	token, err := config.Exchange(ctx, code)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
// 	}

// 	return token, nil
// }

// func generateRandomState() string {
// 	b := make([]byte, 16)
// 	rand.Read(b)
// 	return base64.URLEncoding.EncodeToString(b)
// }

// func extractJobId(locationUrl string) (string, error) {
// 	if locationUrl == "" {
// 		return "", fmt.Errorf("empty location URL")
// 	}

// 	parts := strings.Split(locationUrl, "/")
// 	if len(parts) < 2 {
// 		return "", fmt.Errorf("invalid location URL format: %s", locationUrl)
// 	}

// 	jobId := parts[len(parts)-1]
// 	if jobId == "" {
// 		return "", fmt.Errorf("empty job ID in location URL: %s", locationUrl)
// 	}

// 	return jobId, nil
// }
