package freegames

//Platform interface to implement
type Platform interface {
	Run() (FreeGames, error)
	IsFreeGame(Game) bool
	GetName() string
}
