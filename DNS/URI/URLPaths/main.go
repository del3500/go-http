package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://api.boot.dev/v1/courses_rest_api/learn-http/"

func main() {

	res, err := http.Get(url + "locations")
	if err != nil {
		log.Fatalf("error creating rƒequest: %v", err)
	}
	defer res.Body.Close()

	var locations []Location
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		log.Fatalf("error encoding response: %v", err)
	}
	logLocations(locations)
}

type Location struct {
	Discovered       bool   `json:"discovered"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	RecommendedLevel int    `json:"recommendedLevel"`
}

func logLocations(locations []Location) {
	for _, l := range locations {
		fmt.Printf("Location: %s\nLevel:%v\n", l.Name, l.RecommendedLevel)
	}
}
