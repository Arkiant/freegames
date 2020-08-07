package freegames

//Platform interface to implement
type Platform interface {
	Run() ([]Game, error)
	IsFree(Game) bool
	GetName() string
}
