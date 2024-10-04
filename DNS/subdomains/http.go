package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Name string
}

const domain = "boot.dev"
const sub = "api."

func getItems(domain string) ([]Item, error) {
	res, err := http.Get("https://" + sub + domain + "/v1/courses_rest_api/learn-http/items")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var items []Item
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func logItems(items []Item) {
	for _, item := range items {
		fmt.Println(item.Name)
	}
}
