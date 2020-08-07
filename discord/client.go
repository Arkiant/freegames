package discord

import (
	"errors"
	"fmt"

	"github.com/arkiant/freegames/freegames"
	"github.com/bwmarrin/discordgo"
)

// Client structure
type Client struct {
	db    *freegames.Repository
	token string
	dg    *discordgo.Session
}

// NewDiscordClient is a constructor to create a new discord client
func NewDiscordClient(db *freegames.Repository) *Client {
	return &Client{db: db}
}

// Configure discord bot
func (d *Client) Configure(token string) *Client {
	d.token = token
	return d
}

// GetName get bot name
func (d *Client) GetName() string {
	return "Discord"
}

// Execute discord bot
func (d *Client) Execute() error {
	if d.token == "" {
		return errors.New("Token need to be configured")
	}

	var err error

	d.dg, err = discordgo.New("Bot " + d.token)
	if err != nil {
		return err
	}

	d.dg.AddHandler(messageCreate)
	d.dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = d.dg.Open()
	if err != nil {
		return err
	}

	return nil
}

// Close function to close bot
func (d *Client) Close() {
	d.dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	// TODO: USE COMMAND PATTERN
	if m.Content == "!freegames" {
		fmt.Println("Getting freegames!")
	}
}
