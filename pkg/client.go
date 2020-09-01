package freegames

// Client abstraction
type Client interface {
	Execute() error
	GetName() string
	Close()
	SendFreeGames() error
	SendFreeGamesToChannel(string) error
}

// ClientCommands available to implement
type ClientCommands interface {
	FreegamesCommand() Command
	JoinChannelCommand() Command
}
