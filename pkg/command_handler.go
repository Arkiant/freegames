package freegames

import (
	"context"
	"errors"
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
	Execute(ctx context.Context, c Client) error
}

// CommandHandler has responsability to handler commands
type CommandHandler struct {
	commands map[string]Command
}

// NewCommandHandler create a new empty command handler to be used
func NewCommandHandler() *CommandHandler {
	cmd := make(map[string]Command)
	return &CommandHandler{
		commands: cmd,
	}
}

// Get a client command registered
func (handler CommandHandler) Get(name string) (Command, error) {
	if v, ok := handler.commands[name]; ok {
		return v, nil
	}

	return nil, ErrCommandNotFound
}

// Register a client command to be used
func (handler CommandHandler) Register(name string, cc Command) error {
	if _, ok := handler.commands[name]; !ok {
		handler.commands[name] = cc
		return nil
	}

	return ErrCommandExists
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
func ExecuteCommand(ctx context.Context, c Client, handler *CommandHandler, name string, params []string) error {
	v, err := handler.Get(name)
	if err != nil {
		return err
	}

	// TODO: params usage

	err = v.Execute(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
