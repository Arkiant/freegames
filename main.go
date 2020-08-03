package main

import (
	"fmt"
	"time"

	"github.com/arkiant/freegames/freegames"
	"github.com/arkiant/freegames/unreal"
)

func main() {

	const OnceADay = time.Hour * 24

	pool := []freegames.Platform{
		unreal.NewUnrealGames(),
	}

	// Once a day check for free games
	ticker := time.NewTicker(OnceADay)
	defer ticker.Stop()

	getAllFreeGames(pool)

	for range ticker.C {
		getAllFreeGames(pool)
	}

}

func getAllFreeGames(pool []freegames.Platform) {
	freeGames := make([]freegames.Game, 0)
	for _, v := range pool {
		games, err := v.Run()
		if err != nil {
			panic(err)
		}
		if len(games) > 0 {
			for _, g := range games {
				freeGames = append(freeGames, g)
			}
		}
	}

	fmt.Printf("Freegames: %v", freeGames)
}
