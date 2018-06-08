package foodFin

type restaurantsResponse struct {
	RestaurantList []restaurant `json:"restaurantList"`
}

type restaurant struct {
	name           string   `json:"name"`
	restaurantType string   `json:"type"`
	price          string   `json:"price"`
	distance       float64  `json:"distance"`
	location       location `json:"location"`
}

type location struct {
	lat  float64 `json:"lat"`
	long float64 `json:"long"`
}

func searchRestaurants(currentLocation location, radius float64) restaurantsResponse {

	return restaurantsResponse{
		RestaurantList: []restaurant{
			restaurant{
				name:           "SOUL Food & drintcs",
				restaurantType: "Cafe",
				price:          "Low",
				distance:       150.00,
				location: location{
					lat:  13.8073120,
					long: 100.568980,
				},
			},
			restaurant{
				name:           "ลาบอุดร",
				restaurantType: "อาหารอีสาน",
				price:          "Low",
				distance:       80.00,
				location: location{
					lat:  13.8073284,
					long: 100.568153,
				},
			},
			restaurant{
				name:           "ทิศเหนือ",
				restaurantType: "อาหารเหนือ",
				price:          "Low",
				distance:       50.00,
				location: location{
					lat:  13.8073120,
					long: 100.568980,
				},
			},
		},
	}
}
