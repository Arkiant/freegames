package getting

import (
	"context"

	freegames "github.com/arkiant/freegames/internal"
)

type FreegamesService struct {
	freegamesRepository freegames.GameRepository
	platforms           []freegames.Platform
}

func NewFreegamesService(freegamesRepository freegames.GameRepository, platforms []freegames.Platform) FreegamesService {
	return FreegamesService{freegamesRepository: freegamesRepository, platforms: platforms}
}

func (f FreegamesService) GetFreeGames(ctx context.Context) (interface{}, error) {

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

func (f FreegamesService) getCachedGames(platform freegames.Platform) (freegames.FreeGames, error) {
	return f.freegamesRepository.GetGames(platform)
}

func (f FreegamesService) saveNewGames(games *freegames.FreeGames) {
	for _, g := range *games {
		if !f.freegamesRepository.Exists(g) {
			f.freegamesRepository.Store(g)
		}
	}
}
