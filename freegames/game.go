package freegames

import "time"

// Game struct represent a free game
type Game struct {
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	URL       string    `json:"url"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

// GetAllFreeGames get all free games from all injected platforms
func GetAllFreeGames(pool []Platform, db Repository) []Game {
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
