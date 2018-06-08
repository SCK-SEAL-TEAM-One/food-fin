package foodFin

import (
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

func searchRestaurants(RestaurantsCollection *mgo.Collection, currentLocation location, radius float64) restaurantsResponse {
	restaurants := []restaurant{}
	RestaurantsCollection.Find(bson.M{}).All(&restaurants)
	return restaurantsResponse{
		RestaurantList: restaurants,
	}
}
