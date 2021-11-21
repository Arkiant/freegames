package getting

import (
	"context"
	"log"

	freegames "github.com/arkiant/freegames/internal"
	"github.com/arkiant/freegames/kit/event"
)

type FreegamesService struct {
	freegamesRepository freegames.GameRepository
	platforms           []freegames.Platform
	eventBus            event.Bus
}

func NewFreegamesService(freegamesRepository freegames.GameRepository, platforms []freegames.Platform, eventBus event.Bus) FreegamesService {
	return FreegamesService{freegamesRepository: freegamesRepository, platforms: platforms, eventBus: eventBus}
}

func (f FreegamesService) GetFreeGames(ctx context.Context) (interface{}, error) {

	var freeGames freegames.FreeGames

	for _, v := range f.platforms {

		gamesCached, err := f.getCachedGames(v)
		if err != nil {
			log.Fatal(err)
		}

		if !gamesCached.IsEmpty() {
			freeGames = append(freeGames, gamesCached...)
			continue
		}

		games, err := v.Run()
		if err != nil {
			log.Fatal(err)
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
