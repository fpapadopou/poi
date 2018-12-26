package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Poi represents a point-of-interest
type Poi struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var pois []Poi

// Main application function
func main() {
	router := mux.NewRouter()

	registerRoutes(router)
	if err := http.ListenAndServe(":"+os.Getenv("APP_HTTP_PORT"), router); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	fmt.Println("Waiting for connections..")
}

// Register available routes to router
func registerRoutes(router *mux.Router) {
	router.HandleFunc("/pois", GetPois).Methods(http.MethodGet)
	router.HandleFunc("/pois/{id}", GetPoi).Methods(http.MethodGet)
}

// GetPois returns all available POIs from database
func GetPois(w http.ResponseWriter, r *http.Request) {
	pois = append(pois, Poi{ID: "12", Name: "First Point"})
	pois = append(pois, Poi{ID: "24", Name: "Second Point"})
	json.NewEncoder(w).Encode(pois)
}

// GetPoi returns the specified poi
func GetPoi(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Fetching a poi..")
}
