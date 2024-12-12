package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getResources(path string) []map[string]any {
	fullURL := "https://api.boot.dev" + path

	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return nil
	}

	defer res.Body.Close()

	var resources []map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("error decoding response: ", err)
		return nil
	}
	return resources
}
