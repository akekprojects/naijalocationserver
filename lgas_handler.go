package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muhammadolammi/naijalocationserver/internal/database"
)

func (apiConfig *Config) getLgasHandler(w http.ResponseWriter, r *http.Request) {
	dbstates, err := apiConfig.DB.GetAllStates(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting states: %v", err))
		return
	}
	lgaMaps := getLgasForStates(w, r, dbstates, apiConfig)

	respondWithJson(w, 200, lgaMaps)
}
func (apiConfig *Config) getLgasAHandler(w http.ResponseWriter, r *http.Request) {
	stateA := chi.URLParam(r, "state")

	if stateA == "" {
		log.Println("empty state")
		respondWithError(w, 200, "enter a state name or prefix to get a list of all lgas in that state/states")
		return

	}
	stateA = capitalize(stateA)

	dbstates, err := apiConfig.DB.GetStatesA(r.Context(), sql.NullString{String: stateA, Valid: true})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error getting states: %v", err))
		return
	}
	lgaMaps := getLgasForStates(w, r, dbstates, apiConfig)
	//for empty map of lgas
	if len(lgaMaps) == 0 {
		log.Println("empty lgas")
		respondWithError(w, 200, "empty data for the input state/prefix, make sure the state spelling is correct or the prefix exit for all 36 states of nigeria states, or input the first letter to list all states lgas with that letter")
		return
	}

	respondWithJson(w, 200, lgaMaps)
}

func getLgasForStates(w http.ResponseWriter, r *http.Request, dbstates []database.State, apiConfig *Config) map[string][]string {

	lgaMaps := make(map[string][]string)
	states := databaseStatesToStates(dbstates)
	for _, state := range states {
		dblgas, err := apiConfig.DB.GetStateLgas(r.Context(), state.ID)
		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("error getting lgas for state: %v", err))
			return map[string][]string{}
		}
		lgas := databaseLgasToLgas(dblgas)
		lgasNames := []string{}
		for _, lga := range lgas {
			lgasNames = append(lgasNames, lga.Name)
		}
		lgaMaps[state.Name] = lgasNames
	}
	return lgaMaps
}
