package foodFin

type restaurantsResponse struct {
	RestaurantList []restaurant `json:"restaurantList"`
}

type restaurant struct {
	Name           string   `json:"name"`
	RestaurantType string   `json:"type"`
	Price          string   `json:"price"`
	Distance       float64  `json:"distance"`
	Location       location `json:"location"`
}

type location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func searchRestaurants(currentLocation location, radius float64) restaurantsResponse {

	return restaurantsResponse{
		RestaurantList: []restaurant{
			restaurant{
				Name:           "SOUL Food & Drinks",
				RestaurantType: "Cafe",
				Price:          "Low",
				Distance:       150.00,
				Location: location{
					Lat:  13.8073120,
					Long: 100.568980,
				},
			},
			restaurant{
				Name:           "ลาบอุดร",
				RestaurantType: "อาหารอีสาน",
				Price:          "Low",
				Distance:       80.00,
				Location: location{
					Lat:  13.8073284,
					Long: 100.568153,
				},
			},
			restaurant{
				Name:           "ทิศเหนือ",
				RestaurantType: "อาหารเหนือ",
				Price:          "Low",
				Distance:       50.00,
				Location: location{
					Lat:  13.8073120,
					Long: 100.568980,
				},
			},
		},
	}
}
