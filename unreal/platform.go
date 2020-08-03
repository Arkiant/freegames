package unreal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arkiant/freegames/freegames"
)

const graphqlURL = "https://www.epicgames.com/store/backend/graphql-proxy"

type unrealRequest struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []struct {
					Title     string `json:"title"`
					Price     string `json:"price"`
					KeyImages []struct {
						URL  string `json:"url"`
						Type string `json:"type"`
					} `json:"keyImages"`
					ProductURL string `json:"productSlug"`
				} `json:"elements"`
			} `json:"searchStore"`
		} `json:"Catalog"`
	} `json:"data"`
}

type unrealGames struct{}

// NewUnrealGames create a new unrealGames instance (constructor)
func NewUnrealGames() *unrealGames {
	return new(unrealGames)
}

//Run fetch free games from unreal store
func (u *unrealGames) Run() ([]freegames.Game, error) {

	games := make([]freegames.Game, 0, 4)

	jsonData := []byte(`{"query":"query searchStoreQuery($allowCountries: String, $category: String, $count: Int, $country: String!, $keywords: String, $locale: String, $namespace: String, $itemNs: String, $sortBy: String, $sortDir: String, $start: Int, $tag: String, $releaseDate: String, $withPrice: Boolean = false, $withPromotions: Boolean = false) {  Catalog {    searchStore(allowCountries: $allowCountries, category: $category, count: $count, country: $country, keywords: $keywords, locale: $locale, namespace: $namespace, itemNs: $itemNs, sortBy: $sortBy, sortDir: $sortDir, releaseDate: $releaseDate, start: $start, tag: $tag) {      elements {        title        id        namespace        description        effectiveDate        keyImages {          type          url        }        seller {          id          name        }        productSlug        urlSlug        url        items {          id          namespace        }        customAttributes {          key          value        }        categories {          path        }        price(country: $country) @include(if: $withPrice) {          totalPrice {            discountPrice            originalPrice            voucherDiscount            discount            currencyCode            currencyInfo {              decimals            }            fmtPrice(locale: $locale) {              originalPrice              discountPrice              intermediatePrice            }          }          lineOffers {            appliedRules {              id              endDate              discountSetting {                discountType              }            }          }        }        promotions(category: $category) @include(if: $withPromotions) {          promotionalOffers {            promotionalOffers {              startDate              endDate              discountSetting {                discountType                discountPercentage              }            }          }          upcomingPromotionalOffers {            promotionalOffers {              startDate              endDate              discountSetting {                discountType                discountPercentage              }            }          }        }      }      paging {        count        total      }    }  }}","variables":{"category":"freegames","count":30,"country":"ES","keywords":"","locale":"es-ES","sortBy":"releaseDate","sortDir":"DESC","allowCountries":"ES","start":0,"tag":"","releaseDate":"[,2020-08-03T10:27:59.147Z]","withPrice":true},"operationName":"searchStoreQuery"}`)

	request, err := http.NewRequest("POST", graphqlURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return games, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0")
	request.Header.Set("Accept-Language", "es-ES,es;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Host", "www.epicgames.com")
	request.Header.Set("Origin", "https://www.epicgames.com")
	request.Header.Set("Referer", "https://www.epicgames.com/store/es-ES/browse?sortBy=releaseDate&sortDir=DESC&pageSize=30")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("TE", "Trailers")

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
