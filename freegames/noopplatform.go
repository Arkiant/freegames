package freegames

// NoOpPlatform is a not implemented platform used by tests
type NoOpPlatform struct{}

// NewNoOpPlatform create a new epicgames instance (constructor)
func NewNoOpPlatform() *NoOpPlatform {
	return new(NoOpPlatform)
}

//Run fetch free games from epicgames store
func (u *NoOpPlatform) Run() ([]Game, error) {

	games := make([]Game, 0, 4)

	// TODO: ADD FIXED GAMES

	return games, nil

}

// IsFreeGame check if a game is currently free or not
func (u *NoOpPlatform) IsFreeGame(game Game) bool {
	return true
}

// GetName Get platform name
func (u *NoOpPlatform) GetName() string {
	return "NoOpPlatform"
}
