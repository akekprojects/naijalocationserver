package main

import "strings"

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// func saveLga(api *Config) {
// 	data, err := os.ReadFile("lgas.json")
// 	if err != nil {
// 		fmt.Println("error reading file ", err)
// 		return
// 	}
// 	lgaM := make(map[string][]string)
// 	err = json.Unmarshal(data, &lgaM)
// 	if err != nil {
// 		fmt.Println("error unmarshalling data ", err)
// 		return
// 	}
// 	counts := 0
// 	countL := 0
// 	leng := 0
// 	for s := range lgaM {
// 		for _ = range lgaM[s] {
// 			leng++
// 		}

// 	}
// 	for st := range lgaM {

// 		st = strings.TrimSpace(st)
// 		log.Println(st)
// 		dbstate, err := api.DB.GetState(context.Background(), st)
// 		if err != nil {
// 			fmt.Println("error getting state ", err)
// 			return
// 		}

// 		state := databaseStateToState(dbstate)

// 		lgas := lgaM[st]
// 		counts++

// 		for _, lga := range lgas {
// 			countL++
// 			_, err = api.DB.AddLga(context.Background(), database.AddLgaParams{
// 				Name:    lga,
// 				StateID: state.ID,
// 			})
// 			if err != nil {
// 				log.Println("error adding lga ", err)
// 				return
// 			}
// 		}

// 	}
// 	log.Println("total length ", leng)
// 	log.Println("total state ", counts)
// 	log.Println("total lga ", countL)

// }
