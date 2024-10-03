package main

import (
	"fmt"
)

const baseUrl = "https://api.boot.dev"

func main() {
	locations, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/locations?limit=1")
	if err != nil {
		fmt.Println("Error getting locations:", err)
		return
	}
	fmt.Println("Location:")
	logResources(locations)
	fmt.Println("---")

	items, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/items?limit=1")
	if err != nil {
		fmt.Println("Error getting items:", err)
		return
	}
	fmt.Println("Item:")
	logResources(items)
	fmt.Println("---")

	users, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/users?limit=1")
	if err != nil {
		fmt.Println("Error getting users:", err)
		return
	}
	fmt.Println("User:")
	logResources(users)
}
