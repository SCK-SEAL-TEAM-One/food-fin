package foodFin

import (
	"math"

	haversine "github.com/paultag/go-haversine"
)

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

func CalculateDistance(startLat, startLong, endLat, endLong float64) float64 {
	origin := haversine.Point{Lat: startLat, Lon: startLong}
	destination := haversine.Point{Lat: endLat, Lon: endLong}

	return math.Round(float64(origin.MetresTo(destination))*100) / 100
}

func searchRestaurants(currentLocation location, radius float64) restaurantsResponse {
	restaurants := []restaurant{
		restaurant{
			Name:           "SOUL Food & Drinks",
			RestaurantType: "Cafe",
			Price:          "Low",
			Location: location{
				Lat:  13.8123479,
				Long: 100.5647322,
			},
		},
		restaurant{
			Name:           "ลาบอุดร",
			RestaurantType: "อาหารอีสาน",
			Price:          "Low",
			Location: location{
				Lat:  13.8122417,
				Long: 100.5636951,
			},
		},
		restaurant{
			Name:           "ทิศเหนือ",
			RestaurantType: "อาหารเหนือ",
			Price:          "Low",
			Location: location{
				Lat:  13.8120852,
				Long: 100.5630476,
			},
		},
	}
	for i, restaurant := range restaurants {
		restaurants[i].Distance = CalculateDistance(currentLocation.Lat, currentLocation.Long, restaurant.Location.Lat, restaurant.Location.Long)
	}
	return restaurantsResponse{
		RestaurantList: restaurants,
	}
}
