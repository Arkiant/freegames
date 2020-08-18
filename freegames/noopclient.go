package freegames

// NoOpClient is a not implemented client used by tests
type NoOpClient struct{}

// NewNoOpClient constructor
func NewNoOpClient() *NoOpClient {
	return &NoOpClient{}
}

// Execute functionality
func (cc *NoOpClient) Execute() error {
	return nil
}

// GetName functionality
func (cc *NoOpClient) GetName() string {
	return "NoOpClient"
}

// Close functionality
func (cc *NoOpClient) Close() {}

// SendMessage functionality
func (cc *NoOpClient) SendMessage() error {
	return nil
}

// SendFreeGames functionality
func (cc *NoOpClient) SendFreeGames() error {
	return nil
}

// SendFreeGamesToChannel functionality
func (cc *NoOpClient) SendFreeGamesToChannel(channel string) error {
	return nil
}
