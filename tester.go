package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func getFromEndpoint() interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs

	}
	resp, err := httpClient.Get(" http://naijalocationserver.com/api/states")
	if err != nil {
		log.Printf("error getting resp from endpoint: %v", err)
		return "error getting resp from endpoint"
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var bodyData interface{}
	//this is a bad practice, you should create a struct for body data to be marshalled into that correlate with the json data from endpoint. To manage type better.
	err = decoder.Decode(&bodyData)
	if err != nil {
		log.Printf("error creating decoder from resp.Body(): %v", err)
		return "error creating decoder from resp.Body()"
	}

	return bodyData

}
