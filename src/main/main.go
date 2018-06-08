package main

import (
	"connector"
	"foodFin"
	"log"
	"net/http"
)

func main() {
	Host := []string{
		"127.0.0.1:27017",
	}
	const (
		Database   = "foodfin"
		Collection = "restaurants"
	)
	mongoConnector := connector.Connect(Host)
	RestaurantsCollection := mongoConnector.DB(Database).C(Collection)
	defer mongoConnector.Close()

	httpHandler := foodFin.NewHandler(mongoConnector, RestaurantsCollection)

	http.HandleFunc("/restaurants", httpHandler.RestaurantHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
