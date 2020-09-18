package discord

import freegames "github.com/arkiant/freegames/pkg"

type discordConfiguration struct {
	token   string
	channel string
	prefix  string
}

// NewDiscordConfiguration contructor
func NewDiscordConfiguration(token, prefix, channel string) freegames.ClientConfig {
	return &discordConfiguration{token: token, prefix: prefix, channel: channel}
}

func (dc *discordConfiguration) GetChannel() string {
	return dc.channel
}

func (dc *discordConfiguration) GetToken() string {
	return dc.token
}

func (dc *discordConfiguration) GetPrefix() string {
	return dc.prefix
}
