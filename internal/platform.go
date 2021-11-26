//go:generate mockery --case=snake --outpkg=mockplatform --output=platform/platform/mockplatform --name=Platform

package freegames

//Platform interface to implement
type Platform interface {
	Run() (FreeGames, error)
	GetName() string
}
