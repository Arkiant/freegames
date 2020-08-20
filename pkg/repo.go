package freegames

// Repository find and store free games into a database
type Repository interface {
	Exists(game Game) bool
	GetGames() ([]Game, error)
	Store(game Game) error
	Delete(game Game) error
}
