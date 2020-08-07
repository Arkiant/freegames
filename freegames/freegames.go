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
	db        *Repository
	platforms []Platform
	clients   []Client
}

// NewFreeGames is a constructor to initialize FreeGames object
func NewFreeGames(db *Repository) *Freegames {
	return &Freegames{
		db: db,
	}
}

// Run execute app logic
func (f *Freegames) Run() {
	// Execute clients
	executeClients(f)

	// Execute service in background
	go executeService(f)

}

// Close all connections successful
func (f *Freegames) Close() {
	closeClients(f)
}

// executeService Execute service getting all freegames from all platforms suscribed
func executeService(f *Freegames) {
	ticker := time.NewTicker(OnceADay)
	defer ticker.Stop()

	do := func() {
		fg := getAllFreeGames(f.platforms, *f.db)
		log.Printf("Found %v new free games", len(fg))

		for _, platform := range f.platforms {
			og := deleteOldFreeGames(fg, platform, *f.db)
			log.Printf("Deleted %v old free games from platform: %s", len(og), platform.GetName())
		}

		if len(fg) > 0 {
			sendGamesToClientsConnected(f)
		}
	}

	// Execute functionality at runtime first time
	do()
	for range ticker.C {
		// Execute functionality ticker time
		do()
	}
}

func sendGamesToClientsConnected(f *Freegames) {
	for _, v := range f.clients {
		log.Printf("Sending Message to: %s\n", v.GetName())
		err := v.SendMessage()
		if err != nil {
			log.Printf("Some error ocurried: %s", err.Error())
		}
	}

}
