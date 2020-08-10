package discord

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/arkiant/freegames/freegames"
	"github.com/bwmarrin/discordgo"
)

// Client structure
type Client struct {
	db       *freegames.Repository
	token    string
	dg       *discordgo.Session
	commands *freegames.CommandHandler
}

// NewDiscordClient is a constructor to create a new discord client
func NewDiscordClient(db *freegames.Repository, commands *freegames.CommandHandler) *Client {
	return &Client{db: db, commands: commands}
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

	// 1- AUTHENTICATION
	c.dg, err = discordgo.New("Bot " + c.token)
	if err != nil {
		return err
	}

	// 2- CONFIGURATION

	// 3- HANDLER COMMANDS
	c.dg.AddHandler(c.handlerCommands)
	c.dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// 4- BOT CONNECTIOn
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

// SendFreeGames send all free games to discord client
func (c *Client) SendFreeGames() error {
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

			c.SendFreeGamesToChannel(channel.ID)
		}
	}
	return nil
}

// SendFreeGamesToChannel this method send all games into a specific channel
func (c *Client) SendFreeGamesToChannel(channelID string) error {
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

// handlerCommands execute freegames command
func (c *Client) handlerCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	ctx := context.WithValue(context.Background(), freegames.ChannelID, m.ChannelID)

	log.Printf("Command %s received from %s", m.Content, m.ChannelID)

	err := freegames.ExecuteCommand(ctx, c, c.commands, m.Content)
	if err != nil {
		log.Printf("Some error ocurried with command: %s\n", err.Error())
	}
}
