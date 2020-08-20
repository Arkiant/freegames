package freegames

// Client abstraction
type Client interface {
	Execute() error
	GetName() string
	Close()
	SendFreeGames() error
	SendFreeGamesToChannel(string) error
}
