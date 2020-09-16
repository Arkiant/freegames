package freegames

type DiscordConfiguration struct {
	Enable  bool
	Channel string
}

type ClientsConfiguration []struct {
	Discord *DiscordConfiguration
}

// Configuration config structure
type Configuration struct {
	Clients  ClientsConfiguration
	Platform []string
}

// Config abstraction to decouple and create an adapter to use any configuration library
type Config interface {
	ClientConfig() (ClientsConfiguration, error)
}
