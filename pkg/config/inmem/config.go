package inmem

import (
	"errors"

	freegames "github.com/arkiant/freegames/pkg"
	"github.com/arkiant/freegames/pkg/client/discord"
)

// Configuration structure
type Configuration struct {
	freegames.Configuration
}

// NewInmemConfiguration create a new config implementation using viper library
func NewInmemConfiguration() (freegames.Config, error) {

	clientConfig := make(map[string]freegames.ClientConfig)
	clientConfig["discord"] = discord.NewDiscordConfiguration("test", "!", "#general")

	configuration := freegames.Configuration{Clients: clientConfig}

	return &Configuration{Configuration: configuration}, nil
}

// ClientConfig implementation
func (vc *Configuration) ClientConfig() (map[string]freegames.ClientConfig, error) {
	if len(vc.Configuration.Clients) <= 0 {
		return nil, errors.New("No clients specified")
	}
	return vc.Configuration.Clients, nil
}
