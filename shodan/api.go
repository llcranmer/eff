package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIInfo return information about the user's account
type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

// APIInfo returns the information about the account associated with the client passed.
func (s *Client) APIInfo() (*APIInfo, error) {
	// Use net/http to make a GET request to the Shodan API

	// example := https://api.shodan.io/api-info?key={YOUR_API_KEY}
	// Docs Link := https://developer.shodan.io/api

	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil

}
