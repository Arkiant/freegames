package freegames

import (
	"context"
	"errors"
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

// ExecuteCommand with context and name
func ExecuteCommand(ctx context.Context, c Client, handler *CommandHandler, name string) error {
	v, err := handler.Get(name)
	if err != nil {
		return err
	}

	err = v.Execute(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
