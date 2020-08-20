package epicgames

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	freegames "github.com/arkiant/freegames/pkg"
)

const graphqlURL = "https://www.epicgames.com/store/backend/graphql-proxy"
const gameInfoURL = "https://store-content.ak.epicgames.com/api/es-ES/content/products/%s"
const gameFreeGames = "https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=es-ES&country=ES&allowCountries=ES"

// EpicGames platform integration
type EpicGames struct{}

// NewEpicGames create a new epicgames instance (constructor)
func NewEpicGames() *EpicGames {
	return new(EpicGames)
}

//Run fetch free games from epicgames store
func (u *EpicGames) Run() ([]freegames.Game, error) {

	games := make([]freegames.Game, 0, 4)

	jsonData := createQueryFreeGames()
	request, err := createRequest("POST", graphqlURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return games, err
	}

	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(request)
	if err != nil {
		return games, err
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	rq := epicgamesResponse{}
	json.Unmarshal(data, &rq)

	for _, v := range rq.Data.Catalog.SearchStore.Elements {
		if u.IsFree(v.Price) {

			var photo string

			for _, p := range v.KeyImages {
				if p.Type == "Thumbnail" {
					photo = p.URL
				}
			}

			game := freegames.Game{
				Name:             v.Title,
				Photo:            photo,
				Platform:         u.GetName(),
				URL:              fmt.Sprintf("https://www.epicgames.com/store/es-ES/product/%s", v.ProductURL),
				Slug:             v.ProductURL,
				OfferID:          v.ID,
				ProductNamespace: v.Namespace,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}

			games = append(games, game)
		}
	}

	return games, nil

}

// IsFreeGame check if a game is currently free or not
func (u *EpicGames) IsFreeGame(game freegames.Game) bool {

	response, err := http.Get(gameFreeGames)
	if err != nil {
		log.Printf("Error ocurried in IsFreeGame method in platform %s: %s\n", u.GetName(), err.Error())
		return true
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	rq := epicgamesResponse{}
	json.Unmarshal(data, &rq)

	for _, newGame := range rq.Data.Catalog.SearchStore.Elements {
		if game.Name == newGame.Title {
			log.Printf("No delete %s game, it's free now.\n", game.Name)
			return true
		}
	}

	log.Printf("Deleted %s game, no longer free.\n", game.Name)
	return false
}

// IsFree It's business logic exclusive to the epic games platform that checks whether a game is free
func (u *EpicGames) IsFree(price price) bool {
	return price.TotalPrice.OriginalPrice == price.TotalPrice.Discount || price.TotalPrice.OriginalPrice == 0
}

// GetName Get platform name
func (u *EpicGames) GetName() string {
	return "EpicGames"
}
