package main

import (
	"log"
	"net/http"

	"github.com/feirm/azure/internal/controller"
)

func main() {
	log.Println("Azure microservice...")

	// Routes
	http.HandleFunc("/v1/coins", controller.ListAllCoins)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
