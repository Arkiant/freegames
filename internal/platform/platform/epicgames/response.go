package epicgames

type epicgamesResponse struct {
	Data struct {
		Catalog struct {
			SearchStore struct {
				Elements []struct {
					ID         string     `json:"id"`
					Namespace  string     `json:"namespace"`
					Title      string     `json:"title"`
					Price      price      `json:"price"`
					KeyImages  keyImages  `json:"keyImages"`
					ProductURL string     `json:"productSlug"`
					Promotions promotions `json:"promotions"`
				} `json:"elements"`
			} `json:"searchStore"`
		} `json:"Catalog"`
	} `json:"data"`
}

type keyImages []struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type promotions struct {
	PromotinalOffers []struct {
		PromotinalOffersItem []struct {
			StartDate        string `json:"startDate"`
			EndDate          string `json:"endDate"`
			DiscountSettings struct {
				DiscountTye        string `json:"discountType"`
				DiscountPercentage int    `json:"discountPercentage"`
			} `json:"discountSetting"`
		} `json:"promotionalOffers"`
	} `json:"promotionalOffers"`
	UpcomingPromotinalOffers []struct{} `json:"upcomingPromotionalOffers"`
}

type price struct {
	TotalPrice struct {
		OriginalPrice float64 `json:"originalPrice"`
		DiscountPrice float64 `json:"discountPrice"`
		Discount      float64 `json:"discount"`
	} `json:"totalPrice"`
}
