package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fpapadopou/poi/config"
	"github.com/fpapadopou/poi/datastore"
	"github.com/gorilla/mux"
)

// Server encapsulates router & datastore related structs
type Server struct {
	router *mux.Router
	ds     *datastore.DS
}

// Main application function
func main() {
	// Load app configuration
	appConfig := config.GetApplicationConfig()

	// Create a router and register all routes
	server := &Server{
		router: mux.NewRouter(),
		ds:     datastore.NewDS(config.GetDatabaseConfig()),
	}
	server.registerRoutes()

	// Serve POIs to the people
	if err := http.ListenAndServe(":"+appConfig.Port, server.router); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	fmt.Println("Waiting for connections..")
}

// Register available routes to router
func (s *Server) registerRoutes() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(bytes.NewBufferString("Hiya!").Bytes())
	}).Methods(http.MethodGet)

	s.router.HandleFunc("/poi", s.GetPOIs()).Methods(http.MethodGet)
	s.router.HandleFunc("/poi/{id:[0-9]+}", s.GetPOIByID).Methods(http.MethodGet)
	s.router.HandleFunc("/poi", s.CreatePOI).Methods(http.MethodPost)
	s.router.HandleFunc("/poi/{id:[0-9]+}", s.UpdatePOI).Methods(http.MethodPut)
	s.router.HandleFunc("/poi/{id:[0-9]+}", s.DeletePOI).Methods(http.MethodDelete)
}

// CreatePOI creates a new point-of-interest with the provided data
func (s *Server) CreatePOI(w http.ResponseWriter, r *http.Request) {
	// TODO: Add validator for create/update
	fmt.Println("POI created")
	w.WriteHeader(http.StatusCreated)
}

// GetPOIs returns all available POIs from database
func (s *Server) GetPOIs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pois, _ := s.ds.DatabaseProvider.GetAll()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pois)
	}
}

// GetPOIByID returns the specified POI
func (s *Server) GetPOIByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Invalid POI ID")
	}
	poi, _ := s.ds.DatabaseProvider.GetPOIByID(ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(poi)
}

// UpdatePOI creates a new point-of-interest with the provided data
func (s *Server) UpdatePOI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POI updated")
	w.WriteHeader(http.StatusOK)
}

// DeletePOI deletes an existing POI from the database
func (s *Server) DeletePOI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POI deleted")
	w.WriteHeader(http.StatusOK)
}
