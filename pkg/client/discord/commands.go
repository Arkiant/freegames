package discord

import (
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
func (fgc *freeGamesCommand) Execute(ctx freegames.Context, c freegames.Client) error {
	log.Printf("Executing command freegames from discord\n")

	if ctx.Channel != "" {
		err := c.SendFreeGamesToChannel(ctx.Channel)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("can't convert to string channelID")
}

// TODO: Test
// TODO: Channel

type joinChannelCommand struct{}

// NewJoinChannelCommand create a new join channel command structure
func NewJoinChannelCommand() freegames.Command {
	return &joinChannelCommand{}
}

// Execute method for join channel command
func (fjc *joinChannelCommand) Execute(ctx freegames.Context, c freegames.Client) error {
	if len(ctx.Args) <= 0 {
		return errors.New("channel is mandatory")
	}

	channel := ctx.Args[0]

	extractedChannel := c.ExtractChannel(channel)
	if extractedChannel == "" {
		return errors.New("invalid channel")
	}

	err := c.JoinChannel(string(extractedChannel))
	if err != nil {
		return err
	}

	return nil
}
