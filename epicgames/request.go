package epicgames

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type epicgamesRequest struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []struct {
					Title string `json:"title"`
					Price []struct {
						TotalPrice []struct {
							OriginalPrice string `json:"originalPrice"`
							DiscountPrice string `json:"discountPrice"`
							Discount      string `json:"discount"`
						} `json:"totalPrice"`
					} `json:"price"`
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

func createRequest(data []byte) (*http.Request, error) {
	request, err := http.NewRequest("POST", graphqlURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
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

	return request, nil
}

func createQueryFreeGames() []byte {
	return []byte(fmt.Sprintf(`{"query":"query searchStoreQuery($allowCountries: String, $category: String, $count: Int, $country: String!, $keywords: String, $locale: String, $namespace: String, $itemNs: String, $sortBy: String, $sortDir: String, $start: Int, $tag: String, $releaseDate: String, $withPrice: Boolean = false, $withPromotions: Boolean = false) {  Catalog {    searchStore(allowCountries: $allowCountries, category: $category, count: $count, country: $country, keywords: $keywords, locale: $locale, namespace: $namespace, itemNs: $itemNs, sortBy: $sortBy, sortDir: $sortDir, releaseDate: $releaseDate, start: $start, tag: $tag) {      elements {        title        id        namespace        description        effectiveDate        keyImages {          type          url        }        seller {          id          name        }        productSlug        urlSlug        url        items {          id          namespace        }        customAttributes {          key          value        }        categories {          path        }        price(country: $country) @include(if: $withPrice) {          totalPrice {            discountPrice            originalPrice            voucherDiscount            discount            currencyCode            currencyInfo {              decimals            }            fmtPrice(locale: $locale) {              originalPrice              discountPrice              intermediatePrice            }          }          lineOffers {            appliedRules {              id              endDate              discountSetting {                discountType              }            }          }        }        promotions(category: $category) @include(if: $withPromotions) {          promotionalOffers {            promotionalOffers {              startDate              endDate              discountSetting {                discountType                discountPercentage              }            }          }          upcomingPromotionalOffers {            promotionalOffers {              startDate              endDate              discountSetting {                discountType                discountPercentage              }            }          }        }      }      paging {        count        total      }    }  }}","variables":{"category":"freegames","count":30,"country":"ES","keywords":"","locale":"es-ES","sortBy":"releaseDate","sortDir":"DESC","allowCountries":"ES","start":0,"tag":"","releaseDate":"[,%sT15:00:00.000Z]","withPrice":true},"operationName":"searchStoreQuery"}`, time.Now().Format("2006-01-02")))
}
