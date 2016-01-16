package slacker

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// DefaultAPIURL is the default URL + Path for slack API requests
const DefaultAPIURL = "https://slack.com/api"

// APIClient contains simple logic for starting the RTM Messaging API for Slack
type APIClient struct {
	client   *http.Client
	SlackURL string
}

// NewAPIClient returns a new APIClient with a token set.
func NewAPIClient(token string, url string) *APIClient {
	if url == "" {
		url = DefaultAPIURL
	}

	tkn := &oauth2.Token{AccessToken: token}
	source := oauth2.StaticTokenSource(tkn)
	client := oauth2.NewClient(oauth2.NoContext, source)

	return &APIClient{
		client:   client,
		SlackURL: url,
	}
}

func (c *APIClient) slackMethod(method string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.SlackURL, method), nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}
