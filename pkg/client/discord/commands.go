package discord

import (
	"errors"
	"fmt"
)

// DISCORD COMMANDS IMPLEMENTATION

// FreegamesCommand create a new freegames command concrete to discord client
func (c *client) FreegamesCommand(arg string) error {
	err := c.sendFreeGamesToChannel(c.channel)
	if err != nil {
		return err
	}

	return nil
}

// JoinChannelCommand create a new join channel command concrete to discord client
func (c *client) JoinChannelCommand(arg string) error {
	if len(arg) <= 0 {
		return errors.New("channel is mandatory")
	}

	channel := c.extractChannel(arg)

	fmt.Printf("Hi! I'm joining to channel %s\n", channel)

	return nil
}

// TODO: Test
// TODO: Channel
