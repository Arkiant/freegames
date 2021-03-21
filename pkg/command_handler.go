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

// CommandHandler has responsability to handler commands
type CommandHandler struct {
	commands map[string]Command
}

// NewCommandHandler create a new command handler registering commands availables
func NewCommandHandler(client Client) *CommandHandler {
	cmd := make(map[string]Command)
	cmd["freegames"] = NewFreegamesCommand(client)
	cmd["join"] = NewJoinChannelCommand(client)
	return &CommandHandler{
		commands: cmd,
	}
}

// TODO: get prefix from configuration
const prefix = "!"

// ParseCommand parse text and extract command and arguments
func ParseCommand(text string) (string, string, error) {

	if len(text) <= len(prefix) {
		return "", "", errors.New("wrong command composition")
	}

	if text[0:len(prefix)] != prefix {
		return "", "", errors.New("wrong prefix")
	}

	c := text[len(prefix):]
	s := strings.Fields(c)
	commandName := strings.ToLower(s[0]) // Command name is at first

	if len(s) > 1 {
		args := s[1] // Arguments is after the first element
		return commandName, args, nil
	}

	return commandName, "", nil
}

// ExecuteCommand with context and name
func ExecuteCommand(handler *CommandHandler, name string, arg string) error {
	command, ok := handler.commands[name]
	if !ok {
		return fmt.Errorf("command not found: %s", name)
	}
	err := command.Execute(arg)
	if err != nil {
		return err
	}

	return nil
}
