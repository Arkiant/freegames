package service

import (
	"testing"

	freegames "github.com/arkiant/freegames/pkg"
	inmemConfiguration "github.com/arkiant/freegames/pkg/config/inmem"
	inmemPlatform "github.com/arkiant/freegames/pkg/platform/inmem"
	inmemRepository "github.com/arkiant/freegames/pkg/storage/inmem"
	"github.com/stretchr/testify/assert"
)

func TestAddPlatform(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []freegames.Platform
		Output int
	}{
		{
			Name:   "Add a single platform",
			Input:  []freegames.Platform{inmemPlatform.NewPlatform()},
			Output: 1,
		},
		{
			Name:   "Add three platforms",
			Input:  []freegames.Platform{inmemPlatform.NewPlatform(), inmemPlatform.NewPlatform(), inmemPlatform.NewPlatform()},
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
				fg.AddPlatform(v)
			}

			assert.Len(t, fg.platforms, tc.Output)
		})
	}
}

func TestGetAllFreeGames(t *testing.T) {

	pool := []freegames.Platform{inmemPlatform.NewPlatform()}
	db, err := inmemRepository.NewRepository()
	assert.NoError(t, err)
	games := getAllFreeGames(pool, db)
	assert.Len(t, games, 2)

}
