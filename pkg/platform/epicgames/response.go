package epicgames

type epicgamesResponse struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []struct {
					ID        string `json:"id"`
					Namespace string `json:"namespace"`
					Title     string `json:"title"`
					Price     price  `json:"price"`
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

// type epicgamesGameResponse struct {
// 	Data struct {
// 		Catalog struct {
// 			CatalogOffer struct {
// 				Price price `json:"price"`
// 			} `json:"catalogOffer"`
// 		} `json:"Catalog"`
// 	} `json:"data"`
// }

type price struct {
	TotalPrice struct {
		OriginalPrice float64 `json:"originalPrice"`
		DiscountPrice float64 `json:"discountPrice"`
		Discount      float64 `json:"discount"`
	} `json:"totalPrice"`
}
