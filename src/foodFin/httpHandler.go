package foodFin

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func RestaurantHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	lat, _ := strconv.ParseFloat(queryString.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryString.Get("long"), 64)
	radius, _ := strconv.ParseFloat(queryString.Get("radius"), 64)
	currentLocation := location{lat, long}
	err := json.NewEncoder(w).Encode(searchRestaurants(currentLocation, radius))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
