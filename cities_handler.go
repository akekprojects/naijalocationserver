package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (apiConfig *Config) getCitiesHandler(w http.ResponseWriter, r *http.Request) {
	dbcities, err := apiConfig.DB.GetCities(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting cities: %v", err))
		return
	}
	cities := databaseCitiesToCities(dbcities)
	respondWithJson(w, 200, cities)
}
func (apiConfig *Config) getStateCitiesHandler(w http.ResponseWriter, r *http.Request) {
	stateName := chi.URLParam(r, "state")
	state, err := apiConfig.DB.GetState(r.Context(), capitalize(stateName))
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting state: %v", err))
		return
	}
	dbcities, err := apiConfig.DB.GetStateCities(r.Context(), state.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting cities: %v", err))
		return
	}
	cities := databaseCitiesToCities(dbcities)
	respondWithJson(w, 200, cities)
}
