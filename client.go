package apex

import (
	"net/http"
	"time"

	publicv1 "go.reefassistant.com/apex/client/public/v1"
)

// Client is a local Apex API client.
type Client struct {
	publicv1 *publicv1.APIClient
}

// ClientConfig stores the configuration of the API client.
type ClientConfig struct {
	HTTPClient *http.Client
	UserAgent  string
	Debug      bool
}

// NewDefaultClient creates a new local Apex Client using defaults.
func NewDefaultClient(url string) (*Client, error) {
	return NewClient(url, ClientConfig{
		HTTPClient: &http.Client{
			// We use a fairly big default timeout since Apex controllers can take
			// some time to load certain data. Use request context for better control.
			Timeout: 60 * time.Second,
		},
	})
}

// NewClient creates a new local Apex Client using given configuration.
func NewClient(url string, cfg ClientConfig) (*Client, error) {

	// Public API configuration
	pubCfg := publicv1.NewConfiguration()
	pubCfg.Debug = cfg.Debug
	pubCfg.Servers = publicv1.ServerConfigurations{
		{
			URL: url,
		},
	}

	if cfg.UserAgent != "" {
		pubCfg.UserAgent = cfg.UserAgent
	}

	if cfg.HTTPClient != nil {
		pubCfg.HTTPClient = cfg.HTTPClient
	}

	return &Client{
		publicv1: publicv1.NewAPIClient(pubCfg),
	}, nil
}

// PublicV1 manages communication with the Apex Public API v1.
func (c *Client) PublicV1() *publicv1.APIClient {
	return c.publicv1
}
