package freegames

import "time"

// Game struct represent a free game
type Game struct {
	Name             string    `json:"name"`
	Photo            string    `json:"photo"`
	URL              string    `json:"url"`
	Slug             string    `json:"slug"`
	Platform         string    `json:"platform"`
	OfferID          string    `json:"offerID"`
	ProductNamespace string    `json:"productNamespace"`
	CreatedAt        time.Time `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt        time.Time `bson:"updated_at" json:"updated_at,omitempty"`
	AvailableTo      time.Time `bson:"available_to" json:"available_to"`
}

type FreeGames []Game

func (f FreeGames) IsEmpty() bool {
	return len(f) <= 0
}

//go:generate mockery --case=snake --outpkg=mockrepository --output=platform/storage/mockrepository --name=GameRepository
type GameRepository interface {
	GetGames(platform Platform) (FreeGames, error)
	Exists(game Game) bool
	Store(game Game) error
	Delete(game Game) error
}
