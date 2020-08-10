package discord

import (
	"context"
	"errors"

	"github.com/arkiant/freegames/freegames"
)

// DISCORD COMMANDS IMPLEMENTATION

// FreeGamesCommand implementation
type FreeGamesCommand struct{}

// NewFreeGamesCommand constructor
func NewFreeGamesCommand() *FreeGamesCommand {
	return new(FreeGamesCommand)
}

// Execute method
func (fgc *FreeGamesCommand) Execute(ctx context.Context, c freegames.Client) error {
	if channelID, ok := ctx.Value(freegames.ChannelID).(string); ok {
		err := c.SendMessageToChannel(channelID)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("can't convert to string channelID")
}

// TODO: Test
// TODO: Channel
