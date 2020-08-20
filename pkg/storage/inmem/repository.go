package inmem

import freegames "github.com/arkiant/freegames/pkg"

// Repository is a in memory repository used by tests
type Repository struct {
	database []freegames.Game
}

// NewRepository create a new test repository
func NewRepository() (*Repository, error) {
	repo := Repository{
		database: make([]freegames.Game, 10),
	}

	return &repo, nil

}

// GetGames get all current free games
func (r *Repository) GetGames() ([]freegames.Game, error) {
	return r.database, nil
}

// Exists check if a game exists in database
func (r *Repository) Exists(game freegames.Game) bool {
	for _, v := range r.database {
		if v.Name == game.Name {
			return true
		}
	}

	return false
}

// Store a free game into the database
func (r *Repository) Store(game freegames.Game) error {
	r.database = append(r.database, game)
	return nil
}

// Delete a old free game from the database
func (r *Repository) Delete(game freegames.Game) error {
	for index, v := range r.database {
		if v.Name == game.Name {
			r.database = append(r.database[:index], r.database[index+1:]...)
		}
	}

	return nil
}
