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

	platforms := make([]Platform, 0)
	clients := make([]Client, 0)

	return &Freegames{
		db:        db,
		platforms: platforms,
		clients:   clients,
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
		// Get all free games
		fg := getAllFreeGames(f.platforms, *f.db)
		log.Printf("Found %v new free games", len(fg))

		// Delete all old free games
		deleteAllOldFreeGames(f.platforms, fg, *f.db)

		// Send new games to clients connected
		sendNewGamesToClientsConnected(fg, f)
	}

	// Execute functionality at runtime first time
	do()
	for range ticker.C {
		// Execute functionality ticker time
		do()
	}
}

// sendNewGamesToClientsConnected send new games to all clients connected
func sendNewGamesToClientsConnected(fg []Game, f *Freegames) {
	if len(fg) > 0 {
		sendMessageToClientsConnected(f)
	}
}

// sendMessageToClientsConnected send message to client connected
func sendMessageToClientsConnected(f *Freegames) {
	for _, v := range f.clients {
		log.Printf("Sending Message to: %s\n", v.GetName())
		err := v.SendFreeGames()
		if err != nil {
			log.Printf("Some error ocurried: %s", err.Error())
		}
	}

}
