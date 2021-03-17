package freegames

// Client abstraction
type Client interface {
	Execute() error
	GetName() string
	Close()
	SendFreeGames() error
	SendFreeGamesToChannel(string) error
	JoinChannel(string) error
	ExtractChannel(string) string
	ClientCommands
}

// ClientCommands available to implement
type ClientCommands interface {
	FreegamesCommand() error
	JoinChannelCommand() error
}
