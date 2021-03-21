package freegames

// Command abstraction
type Command interface {
	Execute(arg string) error
}

// FreegamesCommand get all new freegames
type FreegamesCommand struct {
	receiver Client
}

func (fc *FreegamesCommand) Execute(arg string) error {
	return fc.receiver.FreegamesCommand(arg)
}

func NewFreegamesCommand(client Client) Command {
	return &FreegamesCommand{receiver: client}
}

// JoinChannelCommand used to join into a new channel
type JoinChannelCommand struct {
	receiver Client
}

func (fc *JoinChannelCommand) Execute(arg string) error {
	return fc.receiver.JoinChannelCommand(arg)
}

func NewJoinChannelCommand(client Client) Command {
	return &JoinChannelCommand{receiver: client}
}
