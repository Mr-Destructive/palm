package palm

import "net/http"

type ClientConfig struct {
	authToken string

	BaseURL    string
	OrgID      string
	APIVersion string
	HTTPClient *http.Client
}

type Client struct {
	config *ClientConfig
}

func NewClient(authToken string) *Client {
	client := Client{
		config: &ClientConfig{
			authToken:  authToken,
			HTTPClient: &http.Client{},
		},
	}
    return &client
}
