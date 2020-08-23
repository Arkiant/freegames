package discord

import (
	"context"
	"errors"
	"log"

	freegames "github.com/arkiant/freegames/pkg"
)

// DISCORD COMMANDS IMPLEMENTATION

// freeGamesCommand implementation
type freeGamesCommand struct{}

// NewFreeGamesCommand constructor
func NewFreeGamesCommand() freegames.Command {
	return &freeGamesCommand{}
}

// Execute method
func (fgc *freeGamesCommand) Execute(ctx context.Context, c freegames.Client) error {
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
