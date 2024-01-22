package main

import (
	"strings"
)

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// func saveCities(api *Config) {
// 	lgaFileData, err := os.ReadFile("ng.json")
// 	if err != nil {
// 		fmt.Println("error reading file ", err)
// 		return
// 	}
// 	cityM := []struct {
// 		Name       string `json:"city"`
// 		State      string `json:"admin_name"`
// 		Population string `json:"population"`
// 	}{}
// 	err = json.Unmarshal(lgaFileData, &cityM)
// 	if err != nil {
// 		fmt.Println("error unmarshalling file ", err)
// 		return
// 	}
// 	for _, city := range cityM {
// 		if city.State == "Federal Capital Territory" {
// 			city.State = "FCT"
// 		}

// 		state, err := api.DB.GetState(context.Background(), capitalize(city.State))

// 		if err != nil {
// 			fmt.Println("error getting state object ", err)
// 			return
// 		}
// 		_, err = api.DB.AddCity(context.Background(), database.AddCityParams{
// 			StateID:    state.ID,
// 			Name:       city.Name,
// 			Population: city.Population,
// 		})
// 		if err != nil {
// 			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
// 				continue
// 			}
// 			fmt.Println("error saving city ", err)
// 			return
// 		}
// 	}

// }

// // func saveLga(api *Config) {
// // 	data, err := os.ReadFile("lgas.json")
// // 	if err != nil {
// // 		fmt.Println("error reading file ", err)
// // 		return
// // 	}
// // 	lgaM := make(map[string][]string)
// // 	err = json.Unmarshal(data, &lgaM)
// // 	if err != nil {
// // 		fmt.Println("error unmarshalling data ", err)
// // 		return
// // 	}
// // 	counts := 0
// // 	countL := 0
// // 	leng := 0
// // 	for s := range lgaM {
// // 		for _ = range lgaM[s] {
// // 			leng++
// // 		}

// // 	}
// // 	for st := range lgaM {

// // 		st = strings.TrimSpace(st)
// // 		log.Println(st)
// // 		dbstate, err := api.DB.GetState(context.Background(), st)
// // 		if err != nil {
// // 			fmt.Println("error getting state ", err)
// // 			return
// // 		}

// // 		state := databaseStateToState(dbstate)

// // 		lgas := lgaM[st]
// // 		counts++

// // 		for _, lga := range lgas {
// // 			countL++
// // 			_, err = api.DB.AddLga(context.Background(), database.AddLgaParams{
// // 				Name:    lga,
// // 				StateID: state.ID,
// // 			})
// // 			if err != nil {
// // 				log.Println("error adding lga ", err)
// // 				return
// // 			}
// // 		}

// // 	}
// // 	log.Println("total length ", leng)
// // 	log.Println("total state ", counts)
// // 	log.Println("total lga ", countL)

// // }
