package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// HTTPClient defines an interface for the HTTP client
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}







// FetchUser fetches user data using the provided HTTP client
func FetchUser(client HTTPClient, url string) (*User, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching user data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("error decoding user data: %w", err)
	}

	return &user, nil
}
