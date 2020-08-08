package discord

import (
	"errors"
	"fmt"
	"regexp"

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

// TODO: Create complete discord configuration

// Configure discord bot
func (c *Client) Configure(token string) *Client {
	c.token = token
	return c
}

// GetName get bot name
func (c *Client) GetName() string {
	return "Discord"
}

// Execute discord bot
func (c *Client) Execute() error {
	if c.token == "" {
		return errors.New("Token need to be configured")
	}

	var err error

	c.dg, err = discordgo.New("Bot " + c.token)
	if err != nil {
		return err
	}

	c.dg.AddHandler(c.freeGamesCommand)
	c.dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = c.dg.Open()
	if err != nil {
		return err
	}

	return nil
}

// Close function to close bot
func (c *Client) Close() {
	c.dg.Close()
}

// SendMessage send message to discord client
func (c *Client) SendMessage() error {
	fmt.Println("Sending message...")
	for _, guild := range c.dg.State.Guilds {
		channels, _ := c.dg.GuildChannels(guild.ID)

		fmt.Printf("Connected to guild: %s\n", guild.Name)

		for _, channel := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if channel.Type != discordgo.ChannelTypeGuildText {
				continue
			}

			fmt.Printf("Connected to channel: %s\n", channel.Name)

			c.sendMessageToChannel(channel.ID)
		}
	}
	return nil
}

// ExecuteCommand execute a specific command
func (c *Client) ExecuteCommand(command string, clientCommands freegames.ClientCommand) error {

	reg, err := regexp.Compile(`^(?<command>!\w+) (?<text>\S+)|^(?<command>!\w+)`)
	if err != nil {
		return err
	}

	panic("not implemented")
}

// sendMessageToChannel
func (c *Client) sendMessageToChannel(channelID string) error {
	database := *c.db
	games, err := database.GetGames()
	if err != nil {
		return err
	}

	for _, v := range games {
		c.dg.ChannelMessageSend(channelID, v.URL)
	}

	return nil
}

// freeGamesCommand execute freegames command
func (c *Client) freeGamesCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// TODO: USE COMMAND PATTERN
	if m.Content == "!freegames" {
		err := c.sendMessageToChannel(m.ChannelID)
		if err != nil {
			fmt.Printf("Some error ocurried with command: %s", err.Error())
		}
	}
}
