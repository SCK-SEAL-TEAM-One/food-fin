package foodFin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func RestaurantHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	lat, _ := strconv.ParseFloat(queryString.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryString.Get("long"), 64)
	radius, _ := strconv.ParseFloat(queryString.Get("radius"), 64)
	currentLocation := location{lat, long}

	jsonData, _ := json.Marshal(searchRestaurants(currentLocation, radius))
	fmt.Println(string(jsonData))

	err := json.NewEncoder(w).Encode(searchRestaurants(currentLocation, radius))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
