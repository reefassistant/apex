package apex

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	privatev1 "go.reefassistant.com/apex/client/private/v1alpha1"
	publicv1 "go.reefassistant.com/apex/client/public/v1"
)

var (
	userAgent = "reefassistant.com/go-client"
)

// Client is a local Apex API client.
type Client struct {
	privatev1 *privatev1.APIClient
	publicv1  *publicv1.APIClient
}

// ClientConfig stores the configuration of the API client.
type ClientConfig struct {
	HTTPClient *http.Client
	UserAgent  string
	Debug      bool
}

// NewDefaultClient creates a new local Apex Client using defaults.
func NewDefaultClient(url string) (*Client, error) {
	return NewClient(url, ClientConfig{})
}

// NewClient creates a new local Apex Client using given configuration.
func NewClient(url string, cfg ClientConfig) (*Client, error) {

	// Setup our own http client if none supplied
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	// Ensure cookie jar is configured which is currently required
	// for authentication. This may not be necessary anymore once the
	// APIKey OAS cookie support is fixed in the generator being able
	// to use request context session passing.
	if cfg.HTTPClient.Jar == nil {
		jar, _ := cookiejar.New(&cookiejar.Options{})
		cfg.HTTPClient.Jar = jar
	}

	// Set default user agent if not set
	if cfg.UserAgent == "" {
		cfg.UserAgent = userAgent
	}

	// Private API configuration
	priCfg := privatev1.NewConfiguration()
	priCfg.Debug = cfg.Debug
	priCfg.UserAgent = cfg.UserAgent
	priCfg.HTTPClient = cfg.HTTPClient
	priCfg.Servers = privatev1.ServerConfigurations{
		{
			URL: url,
		},
	}

	// Public API configuration
	pubCfg := publicv1.NewConfiguration()
	pubCfg.Debug = cfg.Debug
	pubCfg.UserAgent = cfg.UserAgent
	pubCfg.HTTPClient = cfg.HTTPClient
	pubCfg.Servers = publicv1.ServerConfigurations{
		{
			URL: url,
		},
	}

	return &Client{
		privatev1: privatev1.NewAPIClient(priCfg),
		publicv1:  publicv1.NewAPIClient(pubCfg),
	}, nil
}

// PrivateV1 manages communication with the Apex Private API v1.
func (c *Client) PrivateV1() *privatev1.APIClient {
	return c.privatev1
}

// PublicV1 manages communication with the Apex Public API v1.
func (c *Client) PublicV1() *publicv1.APIClient {
	return c.publicv1
}
