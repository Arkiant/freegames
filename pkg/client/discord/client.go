package discord

import (
	"errors"
	"fmt"
	"log"

	freegames "github.com/arkiant/freegames/pkg"
	"github.com/bwmarrin/discordgo"
)

// client structure
type client struct {
	db    freegames.Repository
	token string
	dg    *discordgo.Session
	ch    *freegames.CommandHandler
}

// NewDiscordClient is a constructor to create a new discord client
func NewDiscordClient(db freegames.Repository, token string) freegames.Client {
	c := &client{db: db, token: token}
	ch := freegames.NewCommandHandler(c)
	c.ch = ch
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
		return errors.New("token need to be configured")
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

			c.sendFreeGamesToChannel(channel.ID)
		}
	}
	return nil
}

// sendFreeGamesToChannel this method send all games into a specific channel
func (c *client) sendFreeGamesToChannel(channelID string) error {
	games, err := c.db.GetGames()
	if err != nil {
		return err
	}

	for _, v := range games {
		c.dg.ChannelMessageSend(channelID, v.URL)
	}

	return nil
}

// handlerCommands execute freegames command
func (c *client) handlerCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, args, err := freegames.ParseCommand(m.Content)
	if err != nil {
		log.Printf("Some error ocurried while extract command: %s\n", err.Error())
		return
	}

	log.Printf("Command %s received from %s", command, m.ChannelID)

	err = freegames.ExecuteCommand(c.ch, command, args)
	if err != nil {
		log.Printf("Some error ocurried with command: %s\n", err.Error())
	}
}

// extractChannel extracts channel number concrete to discord channels
// func (c *client) extractChannel(channel string) string {
// 	re := regexp.MustCompile(`[0-9]+`)
// 	extractedChannel := re.FindAll([]byte(channel), -1)
// 	if len(extractedChannel) <= 0 {
// 		return ""
// 	}
// 	return string(extractedChannel[0])
// }
