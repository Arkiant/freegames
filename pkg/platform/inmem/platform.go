package inmem

import freegames "github.com/arkiant/freegames/pkg"

// Platform is a not implemented platform used by tests
type Platform struct{}

// NewPlatform create a new epicgames instance (constructor)
func NewPlatform() *Platform {
	return new(Platform)
}

//Run fetch free games from epicgames store
func (u *Platform) Run() ([]freegames.Game, error) {

	games := make([]freegames.Game, 0, 4)

	// TODO: ADD FIXED GAMES

	return games, nil

}

// IsFreeGame check if a game is currently free or not
func (u *Platform) IsFreeGame(game freegames.Game) bool {
	return true
}

// GetName Get platform name
func (u *Platform) GetName() string {
	return "InmemPlatform"
}
