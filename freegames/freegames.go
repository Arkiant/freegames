package freegames

import (
	"log"
	"time"
)

const (
	// OnceADay interval
	OnceADay time.Duration = time.Second * 24
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
	// Once a day check for free games
	ticker := time.NewTicker(OnceADay)
	defer ticker.Stop()

	fg := getAllFreeGames(f.pool, *f.db)
	log.Printf("Found %v new free games", len(fg))

	for range ticker.C {
		fg = getAllFreeGames(f.pool, *f.db)
		log.Printf("Found %v new free games", len(fg))
	}
}

// GetAllFreeGames get all free games from all injected platforms
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
