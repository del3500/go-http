package main

import (
	"log"
)

func main() {
	items, err := getItems(domain)
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}
	logItems(items)
}
