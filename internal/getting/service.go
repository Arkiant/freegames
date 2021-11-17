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

	var freeGames []freegames.Game

	for _, v := range f.platforms {
		games, err := v.Run()
		if err != nil {
			log.Fatal(err)
		}

		freeGames = append(freeGames, games...)

	}

	return freeGames, nil
}
