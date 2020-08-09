package freegames

import (
	"errors"
)

// Command abstraction
type Command interface {
	Execute() error
}

// CommandHandler has responsability to handler commands
type CommandHandler struct {
	commands map[string]Command
}

// NewCommandHandler create a new empty command handler to be used
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

// Get a client command registered
func (handler CommandHandler) Get(name string) (Command, error) {
	if v, ok := handler.commands[name]; ok {
		return v, nil
	}

	return nil, errors.New("command not found")
}

// Register a client command to be used
func (handler CommandHandler) Register(name string, cc Command) error {
	if _, ok := handler.commands[name]; !ok {
		handler.commands[name] = cc
	}

	return errors.New("command exists")
}

// ExecuteCommand with context and name
func ExecuteCommand(c Client, handler CommandHandler, name string) error {
	v, err := handler.Get(name)
	if err != nil {
		return err
	}

	err = v.Execute()
	if err != nil {
		return err
	}

	return nil
}
