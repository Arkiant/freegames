package inmem

import freegames "github.com/arkiant/freegames/pkg"

// client is a not implemented client used by tests
type client struct{}

// NewClient constructor
func NewClient() freegames.Client {
	return &client{}
}

// Execute functionality
func (cc *client) Execute() error {
	return nil
}

// GetName functionality
func (cc *client) GetName() string {
	return "InmemClient"
}

// Close functionality
func (cc *client) Close() {}

// SendMessage functionality
func (cc *client) SendMessage() error {
	return nil
}

// SendFreeGames functionality
func (cc *client) SendFreeGames() error {
	return nil
}

// SendFreeGamesToChannel functionality
func (cc *client) SendFreeGamesToChannel(channel string) error {
	return nil
}
