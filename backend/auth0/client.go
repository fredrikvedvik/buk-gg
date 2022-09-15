package auth0

// Config is the configuration for the client
type Config struct {
	Domain    string
	Audiences []string
}

// Client is the auth client
type Client struct {
	config Config
}

// New returns a new Auth client
func New(config Config) *Client {
	return &Client{
		config,
	}
}
