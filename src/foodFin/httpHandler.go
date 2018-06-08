package foodFin

import (
	"encoding/json"
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"
)

type httpHandler struct {
	mongoConnector       *mgo.Session
	restaurantCollection *mgo.Collection
}

func NewHandler(mongoConnector *mgo.Session, restaurantCollection *mgo.Collection) httpHandler {
	return httpHandler{
		mongoConnector:       mongoConnector,
		restaurantCollection: restaurantCollection,
	}
}

func (h httpHandler) RestaurantHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	lat, _ := strconv.ParseFloat(queryString.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryString.Get("long"), 64)
	radius, _ := strconv.ParseFloat(queryString.Get("radius"), 64)
	currentLocation := location{lat, long}

	responses := searchRestaurants(h.restaurantCollection, currentLocation, radius)
	err := json.NewEncoder(w).Encode(responses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
