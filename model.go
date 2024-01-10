package main

import (
	"github.com/google/uuid"
	"github.com/muhammadolammi/naijalocationserver/internal/database"
)

type State struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Capital string    `json:"capital"`
}

func databaseStateToState(dbState database.State) State {
	return State{
		ID:      dbState.ID,
		Name:    dbState.Name,
		Capital: dbState.Capital,
	}
}
func databaseStatesToStates(dbStates []database.State) []State {
	states := []State{}
	for _, st := range dbStates {
		states = append(states, databaseStateToState(st))
	}
	return states
}

// Lga
type Lga struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	StateId uuid.UUID `json:"state_id"`
}

func databaseLgaToLga(dbLga database.Lga) Lga {
	return Lga{
		ID:      dbLga.ID,
		Name:    dbLga.Name,
		StateId: dbLga.StateID,
	}
}
func databaseLgasToLgas(dbLgas []database.Lga) []Lga {
	lgas := []Lga{}
	for _, dbLga := range dbLgas {
		lgas = append(lgas, databaseLgaToLga(dbLga))
	}
	return lgas
}

// City

type City struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Population string    `json:"population"`
}

func databaseCityToCity(dbCity database.City) City {
	return City{
		ID:         dbCity.ID,
		Name:       dbCity.Name,
		Population: dbCity.Population,
	}
}
func databaseCitiesToCities(dbCities []database.City) []City {
	cities := []City{}
	for _, dbCity := range dbCities {
		cities = append(cities, databaseCityToCity(dbCity))
	}
	return cities
}
