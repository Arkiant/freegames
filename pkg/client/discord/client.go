package discord

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	freegames "github.com/arkiant/freegames/pkg"
	"github.com/bwmarrin/discordgo"
)

// client structure
type client struct {
	db       *freegames.Repository
	token    string
	dg       *discordgo.Session
	channel  string
	commands *freegames.CommandHandler
}

// NewDiscordClient is a constructor to create a new discord client
func NewDiscordClient(db *freegames.Repository, token string) freegames.Client {
	c := &client{db: db, token: token}
	ch := freegames.NewCommandHandler(c)
	c.commands = ch
	return c
}

// TODO: Create complete discord configuration

// GetName get bot name
func (c *client) GetName() string {
	return "Discord"
}

// Execute discord bot
func (c *client) Execute() error {
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
func (c *client) Close() {
	c.dg.Close()
}

// SendFreeGames send all free games to discord client
func (c *client) SendFreeGames() error {
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
func (c *client) SendFreeGamesToChannel(channelID string) error {
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

// JoinChannel functionality, this implementation use a channel property to save current channel to answer
func (c *client) JoinChannel(channelID string) error {
	c.channel = channelID
	return nil
}

// handlerCommands execute freegames command
func (c *client) handlerCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, args, err := freegames.ExtractCommand(m.Content)
	if err != nil {
		log.Printf("Some error ocurried while extract command: %s\n", err.Error())
		return
	}

	if c.channel == "" {
		c.channel = m.ChannelID
	}

	ctx := freegames.Context{
		Channel: c.channel,
		Args:    args,
	}

	log.Printf("Command %s received from %s", command, m.ChannelID)

	err = freegames.ExecuteCommand(ctx, c, c.commands, command, args)
	if err != nil {
		log.Printf("Some error ocurried with command: %s\n", err.Error())
	}
}

// FreegamesCommand create a new freegames command concrete to discord client
func (c *client) FreegamesCommand() freegames.Command {
	return NewFreeGamesCommand()
}

// JoinChannelCommand create a new join channel command concrete to discord client
func (c *client) JoinChannelCommand() freegames.Command {
	return NewJoinChannelCommand()
}

// ExtractChannel extracts channel number concrete to discord channels
func (c *client) ExtractChannel(channel string) string {
	re := regexp.MustCompile(`[0-9]+`)
	extractedChannel := re.FindAll([]byte(channel), -1)
	if len(extractedChannel) <= 0 {
		return ""
	}
	return string(extractedChannel[0])
}
