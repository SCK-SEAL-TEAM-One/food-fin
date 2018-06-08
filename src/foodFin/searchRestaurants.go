package foodFin

import (
	"math"

	haversine "github.com/paultag/go-haversine"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type restaurantsResponse struct {
	RestaurantList []restaurant `json:"restaurantList"`
}

type restaurant struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name           string        `json:"name" bson:"name"`
	RestaurantType string        `json:"type" bson:"type"`
	Price          string        `json:"price" bson:"price"`
	Distance       float64       `json:"distance" bson:"distance"`
	Location       location      `json:"location" bson:"location"`
}

type location struct {
	Lat  float64 `json:"lat" bson:"lat"`
	Long float64 `json:"long" bson:"long"`
}

func CalculateDistance(startLat, startLong, endLat, endLong float64) float64 {
	origin := haversine.Point{Lat: startLat, Lon: startLong}
	destination := haversine.Point{Lat: endLat, Lon: endLong}

	return math.Round(float64(origin.MetresTo(destination))*100) / 100
}

func searchRestaurants(RestaurantsCollection *mgo.Collection, currentLocation location, radius float64) restaurantsResponse {
	restaurants := []restaurant{}
	RestaurantsCollection.Find(bson.M{}).All(&restaurants)
	for i, restaurant := range restaurants {
		restaurants[i].Distance = CalculateDistance(currentLocation.Lat, currentLocation.Long, restaurant.Location.Lat, restaurant.Location.Long)
	}
	return restaurantsResponse{
		RestaurantList: restaurants,
	}
}
