package inmem

import (
	"errors"

	freegames "github.com/arkiant/freegames/pkg"
)

type Configuration struct {
	freegames.Configuration
}

// NewViperConfiguration create a new config implementation using viper library
func NewInmemConfiguration() (freegames.Config, error) {
	configuration := freegames.Configuration{Clients: freegames.ClientsConfiguration{struct {
		Discord *freegames.DiscordConfiguration
	}{Discord: &freegames.DiscordConfiguration{Enable: true, Channel: "#freegames"}}}}

	return &Configuration{Configuration: configuration}, nil
}

func (vc *Configuration) ClientConfig() (freegames.ClientsConfiguration, error) {
	if len(vc.Configuration.Clients) <= 0 {
		return nil, errors.New("No clients specified")
	}
	return vc.Configuration.Clients, nil
}
