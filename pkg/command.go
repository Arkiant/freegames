package freegames

// Command abstraction
type Command interface {
	Execute() error
}

// FreegamesCommand get all new freegames
type FreegamesCommand struct {
	receiver Client
}

func (fc *FreegamesCommand) Execute() error {
	return fc.receiver.FreegamesCommand()
}

func NewFreegamesCommand(client Client) Command {
	return &FreegamesCommand{receiver: client}
}

// JoinChannelCommand used to join into a new channel
type JoinChannelCommand struct {
	receiver Client
}

func (fc *JoinChannelCommand) Execute() error {
	return fc.receiver.JoinChannelCommand()
}

func NewJoinChannelCommand(client Client) Command {
	return &JoinChannelCommand{receiver: client}
}
