package epicgames

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	freegames "github.com/arkiant/freegames/internal"
)

const gameFreeGames = "https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=en-US&country=ES&allowCountries=ES"

// platform platform integration
type platform struct{}

// NewEpicGames create a new epicgames instance (constructor)
func NewEpicGames() freegames.Platform {
	return &platform{}
}

//Run fetch free games from epicgames store
func (u *platform) Run() (freegames.FreeGames, error) {

	games := make([]freegames.Game, 0, 4)
	rq, err := u.getFreeGames()
	if err != nil {
		return games, err
	}

	for _, v := range rq.Data.Catalog.SearchStore.Elements {
		if isFreeGame(v.Promotions) {
			availableTo, _ := time.Parse(time.RFC3339, v.Promotions.PromotinalOffers[0].PromotinalOffersItem[0].EndDate)
			game := freegames.Game{
				Name:             v.Title,
				Photo:            searchImage(v.KeyImages),
				Platform:         u.GetName(),
				URL:              fmt.Sprintf("https://www.epicgames.com/store/es-ES/product/%s", v.ProductURL),
				Slug:             v.ProductURL,
				OfferID:          v.ID,
				ProductNamespace: v.Namespace,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
				AvailableTo:      availableTo,
			}

			games = append(games, game)
		}

	}

	return games, nil

}

func searchImage(images keyImages) string {
	for _, p := range images {
		if p.Type == "Thumbnail" {
			return p.URL
		}
	}
	return ""
}

func isFreeGame(p promotions) bool {
	return len(p.PromotinalOffers) > 0 && len(p.PromotinalOffers[0].PromotinalOffersItem) > 0
}

func (u *platform) getFreeGames() (epicgamesResponse, error) {
	rq := epicgamesResponse{}

	request, err := createRequest("GET", gameFreeGames, nil)
	if err != nil {
		return rq, err
	}

	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(request)
	if err != nil {
		return rq, err
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))

	json.Unmarshal(data, &rq)
	return rq, nil
}

// GetName Get platform name
func (u *platform) GetName() string {
	return "EpicGames"
}
