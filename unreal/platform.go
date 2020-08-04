package unreal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arkiant/freegames/freegames"
)

const graphqlURL = "https://www.epicgames.com/store/backend/graphql-proxy"

type unrealGames struct{}

// NewUnrealGames create a new unrealGames instance (constructor)
func NewUnrealGames() *unrealGames {
	return new(unrealGames)
}

//Run fetch free games from unreal store
func (u *unrealGames) Run() ([]freegames.Game, error) {

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

	rq := unrealRequest{}
	json.Unmarshal(data, &rq)

	for _, v := range rq.Data.Catalog.SearchStore.Elements {

		var photo string

		for _, p := range v.KeyImages {
			if p.Type == "Thumbnail" {
				photo = p.URL
			}
		}

		game := freegames.Game{
			Name:      v.Title,
			Photo:     photo,
			Platform:  "unreal",
			URL:       fmt.Sprintf("https://www.epicgames.com/store/es-ES/product/%s", v.ProductURL),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		games = append(games, game)
	}

	return games, nil

}
