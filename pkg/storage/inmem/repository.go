package inmem

import freegames "github.com/arkiant/freegames/pkg"

// repository is a in memory repository used by tests
type repository struct {
	database []freegames.Game
}

// NewRepository create a new test repository
func NewRepository() (freegames.Repository, error) {
	repo := repository{
		database: make([]freegames.Game, 10),
	}

	return &repo, nil

}

// GetGames get all current free games
func (r *repository) GetGames() ([]freegames.Game, error) {
	return r.database, nil
}

// Exists check if a game exists in database
func (r *repository) Exists(game freegames.Game) bool {
	for _, v := range r.database {
		if v.Name == game.Name {
			return true
		}
	}

	return false
}

// Store a free game into the database
func (r *repository) Store(game freegames.Game) error {
	r.database = append(r.database, game)
	return nil
}

// Delete a old free game from the database
func (r *repository) Delete(game freegames.Game) error {
	for index, v := range r.database {
		if v.Name == game.Name {
			r.database = append(r.database[:index], r.database[index+1:]...)
		}
	}

	return nil
}
