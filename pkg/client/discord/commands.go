package discord

import (
	"context"
	"errors"
	"log"

	freegames "github.com/arkiant/freegames/pkg"
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
	log.Printf("Executing command freegames from discord\n")
	if channelID, ok := ctx.Value(freegames.ChannelID).(string); ok {
		err := c.SendFreeGamesToChannel(channelID)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("can't convert to string channelID")
}

// TODO: Test
// TODO: Channel
