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
				name:           "",
				restaurantType: "",
				price:          "",
				distance:       0.00,
				location: location{
					lat:  0.00,
					long: 0.00,
				},
			},
			restaurant{
				name:           "",
				restaurantType: "",
				price:          "",
				distance:       0.00,
				location: location{
					lat:  0.00,
					long: 0.00,
				},
			},
			restaurant{
				name:           "",
				restaurantType: "",
				price:          "",
				distance:       0.00,
				location: location{
					lat:  0.00,
					long: 0.00,
				},
			},
		},
	}
}
