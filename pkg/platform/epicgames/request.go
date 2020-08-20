package epicgames

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func createRequest(method string, url string, data io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	if data != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0")
	request.Header.Set("Accept-Language", "es-ES,es;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("TE", "Trailers")

	return request, nil

}

func createQueryFreeGames() []byte {
	return []byte(fmt.Sprintf(`{"query":"query searchStoreQuery($allowCountries: String, $category: String, $count: Int, $country: String!, $keywords: String, $locale: String, $namespace: String, $itemNs: String, $sortBy: String, $sortDir: String, $start: Int, $tag: String, $releaseDate: String, $withPrice: Boolean = false, $withPromotions: Boolean = false) {  Catalog {    searchStore(allowCountries: $allowCountries, category: $category, count: $count, country: $country, keywords: $keywords, locale: $locale, namespace: $namespace, itemNs: $itemNs, sortBy: $sortBy, sortDir: $sortDir, releaseDate: $releaseDate, start: $start, tag: $tag) {      elements {        title        id        namespace        description        effectiveDate        keyImages {          type          url        }        seller {          id          name        }        productSlug        urlSlug        url        items {          id          namespace        }        customAttributes {          key          value        }        categories {          path        }        price(country: $country) @include(if: $withPrice) {          totalPrice {            discountPrice            originalPrice            voucherDiscount            discount            currencyCode            currencyInfo {              decimals            }            fmtPrice(locale: $locale) {              originalPrice              discountPrice              intermediatePrice            }          }          lineOffers {            appliedRules {              id              endDate              discountSetting {                discountType              }            }          }        }        promotions(category: $category) @include(if: $withPromotions) {          promotionalOffers {            promotionalOffers {              startDate              endDate              discountSetting {                discountType                discountPercentage              }            }          }          upcomingPromotionalOffers {            promotionalOffers {              startDate              endDate              discountSetting {                discountType                discountPercentage              }            }          }        }      }      paging {        count        total      }    }  }}","variables":{"category":"freegames","count":30,"country":"ES","keywords":"","locale":"es-ES","sortBy":"releaseDate","sortDir":"DESC","allowCountries":"ES","start":0,"tag":"","releaseDate":"[,%sT15:00:00.000Z]","withPrice":true},"operationName":"searchStoreQuery"}`, time.Now().Format("2006-01-02")))
}

func createQueryGame(offerID string, productNamespace string) []byte {
	return []byte(fmt.Sprintf(`{"query":"query catalogQuery($productNamespace: String!, $offerId: String!, $locale: String, $country: String!, $includeSubItems: Boolean!) {\n  Catalog {\n    catalogOffer(namespace: $productNamespace, id: $offerId, locale: $locale) {\n      title\n      id\n      namespace\n      description\n      effectiveDate\n      expiryDate\n      isCodeRedemptionOnly\n      keyImages {\n        type\n        url\n      }\n      seller {\n        id\n        name\n      }\n      productSlug\n      urlSlug\n      url\n      items {\n        id\n        namespace\n      }\n      customAttributes {\n        key\n        value\n      }\n      categories {\n        path\n      }\n      price(country: $country) {\n        totalPrice {\n          discountPrice\n          originalPrice\n          voucherDiscount\n          discount\n          currencyCode\n          currencyInfo {\n            decimals\n          }\n          fmtPrice(locale: $locale) {\n            originalPrice\n            discountPrice\n            intermediatePrice\n          }\n        }\n        lineOffers {\n          appliedRules {\n            id\n            endDate\n            discountSetting {\n              discountType\n            }\n          }\n        }\n      }\n    }\n    offerSubItems(namespace: $productNamespace, id: $offerId) @include(if: $includeSubItems) {\n      namespace\n      id\n      releaseInfo {\n        appId\n        platform\n      }\n    }\n  }\n}\n","variables":{"productNamespace":"%s","offerId":"%s","locale":"es-ES","country":"ES","lineOffers":[{"offerId":"%s","quantity":1}],"calculateTax":false,"includeSubItems":true}`, productNamespace, offerID, offerID))
}
