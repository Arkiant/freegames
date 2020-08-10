package freegames

import "log"

//Platform interface to implement
type Platform interface {
	Run() ([]Game, error)
	IsFreeGame(Game) bool
	GetName() string
}

// AddPlatform using chain pattern we can add multiple platforms to get free games
func (f *Freegames) AddPlatform(platform Platform) *Freegames {
	f.platforms = append(f.platforms, platform)
	return f
}

// getAllFreeGames from all injected platforms
func getAllFreeGames(pool []Platform, db Repository) []Game {
	freeGames := make([]Game, 0)
	for _, v := range pool {

		games, err := v.Run()
		if err != nil {
			panic(err)
		}
		if len(games) > 0 {
			for _, g := range games {
				if !db.Exists(g) {
					db.Store(g)
					freeGames = append(freeGames, g)
				}
			}
		}
	}

	return freeGames
}

// deleteAllOldFreeGames from the database
func deleteAllOldFreeGames(platforms []Platform, currentFreegames []Game, db Repository) {

	for _, platform := range platforms {
		og := deleteOldFreeGames(currentFreegames, platform, db)
		log.Printf("Deleted %v old free games from platform: %s", len(og), platform.GetName())
	}

}

// deleteAllFreeGames from the database
func deleteOldFreeGames(currentFreeGames []Game, platform Platform, db Repository) []Game {
	freeGames := make([]Game, 0)
	allGames, err := db.GetGames()
	if err != nil {
		return freeGames
	}

	for _, game := range allGames {
		deleted := false
		for _, currentGame := range currentFreeGames {
			if currentGame.Name == game.Name {
				deleted = true
			}
		}

		if !deleted {
			if !platform.IsFreeGame(game) {
				err := db.Delete(game)
				if err != nil {
					return freeGames
				}
				freeGames = append(freeGames, game)
			}
		}
	}

	return freeGames
}
