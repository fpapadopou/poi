package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "github.com/fpapadopou/poi/config"
	"github.com/gorilla/mux"
)

// POI represents a point-of-interest
type POI struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var pois []POI

// Main application function
func main() {
	// Load app configuration
	appConfig := config.GetApplicationConfig()

	// Create a router and register all routes
	router := mux.NewRouter()
	registerRoutes(router)

	// Serve POIs to the people
	if err := http.ListenAndServe(":"+appConfig.Port, router); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	fmt.Println("Waiting for connections..")
}

// Register available routes to router
func registerRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bytes.NewBufferString("Hiya!").Bytes())
	}).Methods(http.MethodGet)

	router.HandleFunc("/pois", GetPOIs).Methods(http.MethodGet)
	router.HandleFunc("/pois/{id:[0-9]+}", GetPOIByID).Methods(http.MethodGet)
	router.HandleFunc("/pois", CreatePOI).Methods(http.MethodPost)
	router.HandleFunc("/pois/{id:[0-9]+}", UpdatePOI).Methods(http.MethodPut)
	router.HandleFunc("/pois/{id:[0-9]+}", DeletePOI).Methods(http.MethodDelete)
}

// CreatePOI creates a new point-of-interest with the provided data
func CreatePOI(w http.ResponseWriter, r *http.Request) {
	// TODO: Add validator for create/update
	fmt.Println("POI created")
	w.WriteHeader(http.StatusCreated)
}

// GetPOIs returns all available POIs from database
func GetPOIs(w http.ResponseWriter, r *http.Request) {
	pois = nil
	pois = append(pois, POI{ID: "12", Name: "First POI"})
	pois = append(pois, POI{ID: "24", Name: "Second POI"})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pois)
}

// UpdatePOI creates a new point-of-interest with the provided data
func UpdatePOI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POI updated")
	w.WriteHeader(http.StatusOK)
}

// GetPOIByID returns the specified POI
func GetPOIByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(POI{ID: vars["id"], Name: "POI by its ID"})
}

// DeletePOI deletes an existing POI from the database
func DeletePOI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POI deleted")
	w.WriteHeader(http.StatusOK)
}
