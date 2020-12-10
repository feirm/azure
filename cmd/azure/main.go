package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/feirm/azure/internal/controller"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	asciiArt := `
░█████╗░███████╗██╗░░░██╗██████╗░███████╗
██╔══██╗╚════██║██║░░░██║██╔══██╗██╔════╝
███████║░░███╔═╝██║░░░██║██████╔╝█████╗░░
██╔══██║██╔══╝░░██║░░░██║██╔══██╗██╔══╝░░
██║░░██║███████╗╚██████╔╝██║░░██║███████╗
╚═╝░░╚═╝╚══════╝░╚═════╝░╚═╝░░╚═╝╚══════╝
	`
	fmt.Println(asciiArt)
	log.Println("Starting cryptocurrency data microservice.")

	// Load in the environment variables
	godotenv.Load()

	// Routes
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Get("/coins", controller.ListAllCoins)
	})

	// Serve the HTTP router
	bind := os.Getenv("bind")
	port := os.Getenv("port")

	log.Printf("Serving on http://%s:%s", bind, port)
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", bind, port),
		Handler: r,
	}
	s.ListenAndServe()
}
