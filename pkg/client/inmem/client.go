package inmem

// Client is a not implemented client used by tests
type Client struct{}

// NewClient constructor
func NewClient() *Client {
	return &Client{}
}

// Execute functionality
func (cc *Client) Execute() error {
	return nil
}

// GetName functionality
func (cc *Client) GetName() string {
	return "InmemClient"
}

// Close functionality
func (cc *Client) Close() {}

// SendMessage functionality
func (cc *Client) SendMessage() error {
	return nil
}

// SendFreeGames functionality
func (cc *Client) SendFreeGames() error {
	return nil
}

// SendFreeGamesToChannel functionality
func (cc *Client) SendFreeGamesToChannel(channel string) error {
	return nil
}
