package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Item struct {
	Name    string `json:"name"`
	Quality int    `json:"quality"`
}

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()
	var item []Item
	data, err := io.ReadAll(res.Body)
	json.Unmarshal(data, &item)

	return item, nil
}
