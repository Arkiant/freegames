package freegames

import (
	"log"
	"time"
)

const (
	// OnceADay interval
	OnceADay time.Duration = time.Hour * 24
)

// Freegames is a struct to abstract app execution
type Freegames struct {
	db   *Repository
	pool []Platform
}

// NewFreeGames is a constructor to initialize FreeGames object
func NewFreeGames(db *Repository) *Freegames {
	return &Freegames{
		db: db,
	}
}

// AddPlatform using chain pattern we can add multiple platforms to get free games
func (f *Freegames) AddPlatform(platform Platform) *Freegames {
	f.pool = append(f.pool, platform)
	return f
}

// Run execute app logic
func (f *Freegames) Run() {

	ticker := time.NewTicker(OnceADay)
	defer ticker.Stop()

	do := func() {
		fg := getAllFreeGames(f.pool, *f.db)
		log.Printf("Found %v new free games", len(fg))

		for _, platform := range f.pool {
			og := deleteOldFreeGames(fg, platform, *f.db)
			log.Printf("Deleted %v old free games from platform: %s", len(og), platform.GetName())
		}
	}

	// Execute functionality at runtime first time
	do()
	for range ticker.C {
		// Execute functionality ticker time
		do()
	}
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
			if !platform.IsFree(game) {
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
