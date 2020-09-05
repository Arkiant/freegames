package freegames

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrCommandNotFound when command not found
	ErrCommandNotFound = errors.New("command not found")

	// ErrCommandExists when command exists
	ErrCommandExists = errors.New("command exists")
)

// CommandParams represents parameters passed by context
type CommandParams string

const (
	// ChannelID param
	ChannelID CommandParams = "channelID"
)

// Command abstraction
type Command interface {
	Execute(ctx Context, c Client) error
}

// CommandHandler has responsability to handler commands
type CommandHandler struct {
	commands map[string]Command
}

// NewCommandHandler create a new command handler registering commands availables
func NewCommandHandler(cc ClientCommands) *CommandHandler {
	cmd := make(map[string]Command)
	cmd["freegames"] = cc.FreegamesCommand()
	cmd["join"] = cc.JoinChannelCommand()
	return &CommandHandler{
		commands: cmd,
	}
}

// TODO: get prefix from configuration
const prefix = "!"

// ExtractCommand auxiliar function split string into command, args and any errors ocurried
func ExtractCommand(content string) (string, []string, error) {

	if len(content) <= len(prefix) {
		return "", nil, errors.New("wrong command composition")
	}

	if content[0:len(prefix)] != prefix {
		return "", nil, errors.New("wrong prefix")
	}

	c := content[len(prefix):]
	s := strings.Fields(c)
	commandName := strings.ToLower(s[0])
	args := s[1:]

	return commandName, args, nil
}

// ExecuteCommand with context and name
func ExecuteCommand(ctx Context, c Client, handler *CommandHandler, name string, params []string) error {
	command, ok := handler.commands[name]
	if !ok {
		return fmt.Errorf("command not found: %s", name)
	}
	err := command.Execute(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
