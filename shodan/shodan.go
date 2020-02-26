package shodan

// BaseURL of the Shodan API
const BaseURL = "https://api.shodan.io"

// Client which holds the client key
type Client struct {
	apiKey string
}

// New func returns a pointer to the Client struct's apiKey field
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
