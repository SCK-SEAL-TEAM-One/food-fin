package main

import (
	"foodFin"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/restaurants", foodFin.RestaurantHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
