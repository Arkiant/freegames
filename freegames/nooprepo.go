package freegames

// NoOpRepository is a in memory repository used by tests
type NoOpRepository struct {
	database []Game
}

// NewNoOpRepository create a new test repository
func NewNoOpRepository() (Repository, error) {
	repo := &NoOpRepository{
		database: make([]Game, 10),
	}

	return repo, nil

}

// GetGames get all current free games
func (r *NoOpRepository) GetGames() ([]Game, error) {
	return r.database, nil
}

// Exists check if a game exists in database
func (r *NoOpRepository) Exists(game Game) bool {
	for _, v := range r.database {
		if v.Name == game.Name {
			return true
		}
	}

	return false
}

// Store a free game into the database
func (r *NoOpRepository) Store(game Game) error {
	r.database = append(r.database, game)
	return nil
}

// Delete a old free game from the database
func (r *NoOpRepository) Delete(game Game) error {
	for index, v := range r.database {
		if v.Name == game.Name {
			r.database = append(r.database[:index], r.database[index+1:]...)
		}
	}

	return nil
}
