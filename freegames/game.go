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
