package epicgames

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arkiant/freegames/freegames"
)

const graphqlURL = "https://www.epicgames.com/store/backend/graphql-proxy"
const gameInfoURL = "https://store-content.ak.epicgames.com/api/es-ES/content/products/%s"

type epicGames struct{}

// NewEpicGames create a new epicgames instance (constructor)
func NewEpicGames() *epicGames {
	return new(epicGames)
}

//Run fetch free games from epicgames store
func (u *epicGames) Run() ([]freegames.Game, error) {

	games := make([]freegames.Game, 0, 4)

	jsonData := createQueryFreeGames()
	request, err := createRequest(jsonData)
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

// IsFree check if a game is currently free or not
func (u *epicGames) IsFreeGame(game freegames.Game) bool {

	jsonData := createQueryGame(game.OfferID, game.ProductNamespace)
	request, err := createRequest(jsonData)
	if err != nil {
		return true
	}

	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		return true
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	rq := epicgamesGameResponse{}
	json.Unmarshal(data, &rq)

	return u.IsFree(rq.Data.Catalog.CatalogOffer.Price)
}

// IsFree It's business logic exclusive to the epic games platform that checks whether a game is free
func (u *epicGames) IsFree(price price) bool {
	return price.TotalPrice.OriginalPrice == price.TotalPrice.Discount || price.TotalPrice.OriginalPrice == 0
}

// GetName Get platform name
func (u *epicGames) GetName() string {
	return "EpicGames"
}
