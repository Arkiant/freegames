package inmem

import freegames "github.com/arkiant/freegames/internal"

// repository is a in memory repository used by tests
type repository struct {
	database map[string]freegames.FreeGames
}

// NewRepository create a new test repository
func NewRepository() (freegames.GameRepository, error) {
	repo := repository{
		database: make(map[string]freegames.FreeGames, 10),
	}

	return &repo, nil

}

// GetGames get all current free games
func (r *repository) GetGames(platform freegames.Platform) (freegames.FreeGames, error) {
	return r.database[platform.GetName()], nil
}

// Exists check if a game exists in database
func (r *repository) Exists(game freegames.Game) bool {
	for _, v := range r.database[game.Platform] {
		if v.Name == game.Name {
			return true
		}
	}

	return false
}

// Store a free game into the database
func (r *repository) Store(game freegames.Game) error {
	r.database[game.Platform] = append(r.database[game.Platform], game)
	return nil
}

// Delete a old free game from the database
func (r *repository) Delete(game freegames.Game) error {
	platform := game.Platform
	for index, v := range r.database[platform] {
		if v.Name == game.Name {
			r.database[platform] = append(r.database[platform][:index], r.database[platform][index+1:]...)
		}
	}

	return nil
}
