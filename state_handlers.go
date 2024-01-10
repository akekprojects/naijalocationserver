package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (apiConfig *Config) getStatesHandler(w http.ResponseWriter, r *http.Request) {
	dbstates, err := apiConfig.DB.GetAllStates(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting states from db %v", err))
		return
	}
	states := databaseStatesToStates(dbstates)
	respondWithJson(w, 200, states)

}
func (apiConfig *Config) getStatesAHandler(w http.ResponseWriter, r *http.Request) {

	stateA := chi.URLParam(r, "state")

	if stateA == "" {
		log.Println("empty state")
		respondWithError(w, 200, "empty data for the input state, make sure the state spelling is correct, or input the first letter to list all states with that letter")
		return

	}
	stateA = capitalize(stateA)

	dbstates, err := apiConfig.DB.GetStatesA(r.Context(), sql.NullString{String: stateA, Valid: true})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting states from db %v", err))
		return
	}
	states := databaseStatesToStates(dbstates)
	respondWithJson(w, 200, states)
}

func (apiConfig *Config) getStateHandler(w http.ResponseWriter, r *http.Request) {
	stateN := chi.URLParam(r, "state")
	stateN = capitalize(stateN)
	dbstate, err := apiConfig.DB.GetState(r.Context(), stateN)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			respondWithError(w, 200, `no state with input name, go to endpoint /states to get a list of all states 
			or /states/{state} to get a list of states that start with the input string`)
			return
		}
		respondWithError(w, 500, fmt.Sprintf("error getting state from db %v", err))
		return
	}
	state := databaseStateToState(dbstate)
	respondWithJson(w, 200, state)

}
