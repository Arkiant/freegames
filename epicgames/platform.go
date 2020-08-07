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

	rq := epicgamesRequest{}
	json.Unmarshal(data, &rq)

	for _, v := range rq.Data.Catalog.SearchStore.Elements {

		// Check is free
		if v.Price.TotalPrice.OriginalPrice == v.Price.TotalPrice.Discount || v.Price.TotalPrice.OriginalPrice == 0 {

			var photo string

			for _, p := range v.KeyImages {
				if p.Type == "Thumbnail" {
					photo = p.URL
				}
			}

			game := freegames.Game{
				Name:      v.Title,
				Photo:     photo,
				Platform:  u.GetName(),
				URL:       fmt.Sprintf("https://www.epicgames.com/store/es-ES/product/%s", v.ProductURL),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			games = append(games, game)
		}
	}

	return games, nil

}

// IsFree check if a game is currently free or not
func (u *epicGames) IsFree(game freegames.Game) bool {
	//TODO: Create request to check if currently is free game
	return true
}

func (u *epicGames) GetName() string {
	return "EpicGames"
}
