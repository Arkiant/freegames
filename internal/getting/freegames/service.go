package freegames

import (
	"context"

	freegames "github.com/arkiant/freegames/internal"
)

type Service struct {
	freegamesRepository freegames.GameRepository
	platforms           []freegames.Platform
}

func NewService(freegamesRepository freegames.GameRepository, platforms ...freegames.Platform) Service {
	return Service{freegamesRepository: freegamesRepository, platforms: platforms}
}

func (f Service) GetFreeGames(ctx context.Context) (interface{}, error) {

	var freeGames freegames.FreeGames

	for _, platform := range f.platforms {

		cachedGames, err := f.getCachedGames(platform)
		if err != nil {
			return nil, err
		}

		if !cachedGames.IsEmpty() {
			freeGames = append(freeGames, cachedGames...)
			continue
		}

		games, err := platform.Run()
		if err != nil {
			return nil, err
		}

		f.saveNewGames(&games)
		freeGames = append(freeGames, games...)

	}

	return freeGames, nil
}

func (f Service) getCachedGames(platform freegames.Platform) (freegames.FreeGames, error) {
	return f.freegamesRepository.GetGames(platform)
}

func (f Service) saveNewGames(games *freegames.FreeGames) {
	for _, g := range *games {
		if !f.freegamesRepository.Exists(g) {
			f.freegamesRepository.Store(g)
		}
	}
}
