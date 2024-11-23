package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	users, err := getUsers(url)
	if err != nil {
		log.Fatalf("Error getting users:", err)
	}
	logUser(users)
}

func getUsers(url string) ([]User, error) {
	var userList []User

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return userList, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return userList, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&userList); err != nil {
		return userList, err
	}

	return userList, err
}
