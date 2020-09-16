package service

import (
	"testing"

	freegames "github.com/arkiant/freegames/pkg"
	inmemClient "github.com/arkiant/freegames/pkg/client/inmem"
	inmemConfiguration "github.com/arkiant/freegames/pkg/config/inmem"
	inmemRepository "github.com/arkiant/freegames/pkg/storage/inmem"
	"github.com/stretchr/testify/assert"
)

func TestAddClient(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []freegames.Client
		Output int
	}{
		{
			Name:   "Add a single client",
			Input:  []freegames.Client{inmemClient.NewClient()},
			Output: 1,
		},
		{
			Name:   "Add three clients",
			Input:  []freegames.Client{inmemClient.NewClient(), inmemClient.NewClient(), inmemClient.NewClient()},
			Output: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {

			db, err := inmemRepository.NewRepository()
			assert.NoError(t, err)

			conf, err := inmemConfiguration.NewInmemConfiguration()
			fg := NewFreeGames(db, conf)
			for _, v := range tc.Input {
				fg.AddClient(v)
			}

			assert.Len(t, fg.clients, tc.Output)
		})
	}
}
