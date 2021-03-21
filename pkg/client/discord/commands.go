package discord

import "errors"

// DISCORD COMMANDS IMPLEMENTATION

// FreegamesCommand create a new freegames command concrete to discord client
func (c *client) FreegamesCommand(arg string) error {
	err := c.sendFreeGamesToChannel(arg)
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

	panic("not implemented")
}

// TODO: Test
// TODO: Channel
