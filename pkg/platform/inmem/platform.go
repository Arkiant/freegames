package inmem

import (
	"time"

	freegames "github.com/arkiant/freegames/pkg"
)

// Platform is a not implemented platform used by tests
type Platform struct{}

// NewPlatform create a new epicgames instance (constructor)
func NewPlatform() *Platform {
	return new(Platform)
}

//Run fetch free games from epicgames store
func (u *Platform) Run() ([]freegames.Game, error) {

	games := make([]freegames.Game, 0, 4)

	games = append(games, freegames.Game{
		Name:             "Test Free Game",
		ProductNamespace: "Freegame",
		CreatedAt:        time.Now(),
		OfferID:          "free-test-game",
		Photo:            "blank",
		Platform:         "InmemPlatform",
		Slug:             "test-free-game",
		URL:              "http://localhost/test-free-game",
		UpdatedAt:        time.Now(),
	})

	games = append(games, freegames.Game{
		Name:             "Test No Free Game",
		ProductNamespace: "NoFreegame",
		CreatedAt:        time.Now(),
		OfferID:          "no-free-test-game",
		Photo:            "blank",
		Platform:         "InmemPlatform",
		Slug:             "test-no-free-game",
		URL:              "http://localhost/test-no-free-game",
		UpdatedAt:        time.Now(),
	})

	return games, nil

}

// IsFreeGame check if a game is currently free or not
func (u *Platform) IsFreeGame(game freegames.Game) bool {

	if game.Name == "Test Free Game" {
		return true
	}

	return false
}

// GetName Get platform name
func (u *Platform) GetName() string {
	return "InmemPlatform"
}
