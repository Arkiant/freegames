package viper

import (
	"errors"

	freegames "github.com/arkiant/freegames/pkg"
	"github.com/spf13/viper"
)

type viperConfiguration struct {
	freegames.Configuration
}

// NewViperConfiguration create a new config implementation using viper library
func NewViperConfiguration(file string) (freegames.Config, error) {
	viper.SetConfigFile(file)

	configuration := freegames.Configuration{}
	err := viper.UnmarshalExact(configuration)
	if err != nil {
		return nil, err
	}

	return &viperConfiguration{Configuration: configuration}, nil
}

func (vc *viperConfiguration) ClientConfig() (map[string]freegames.ClientConfig, error) {
	if len(vc.Configuration.Clients) <= 0 {
		return nil, errors.New("No clients specified")
	}
	return vc.Configuration.Clients, nil
}
