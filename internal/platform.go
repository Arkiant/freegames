package freegames

//Platform interface to implement
type Platform interface {
	Run() (FreeGames, error)
	GetName() string
}
