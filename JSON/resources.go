package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, nil
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&resources); err != nil {
		return nil, fmt.Errorf("error decoding: %w", err)
	}

	return resources, nil
}

func logResources(resources []map[string]any) {
	var formattedString []string

	for _, resource := range resources {
		for k, v := range resource {
			formatted := fmt.Sprintf("Key: %s - Value: %v", k, v)
			formattedString = append(formattedString, formatted)
		}
	}
	sort.Strings(formattedString)
	for _, str := range formattedString {
		fmt.Println(str)
	}
}
