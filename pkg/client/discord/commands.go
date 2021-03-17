package discord

// DISCORD COMMANDS IMPLEMENTATION

// FreegamesCommand create a new freegames command concrete to discord client
func (c *client) FreegamesCommand() error {
	err := c.SendFreeGamesToChannel(c.channel)
	if err != nil {
		return err
	}

	return nil
}

// JoinChannelCommand create a new join channel command concrete to discord client
func (c *client) JoinChannelCommand() error {
	panic("not implemented command")
}

// TODO: Test
// TODO: Channel

// if len(ctx.Args) <= 0 {
// 	return errors.New("channel is mandatory")
// }

// channel := ctx.Args[0]

// extractedChannel := c.ExtractChannel(channel)
// if extractedChannel == "" {
// 	return errors.New("invalid channel")
// }

// err := c.JoinChannel(string(extractedChannel))
// if err != nil {
// 	return err
// }

// return nil
