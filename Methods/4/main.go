package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/rand"
)

func main() {
	userToCreate := User{
		Role:       "Junior Dev",
		Experience: 2,
		Remote:     true,
		User: struct {
			Name     string `json:"name"`
			Location string `json:"location"`
			Age      int    `json:"age"`
		}{
			Name:     "asd",
			Location: "PH",
			Age:      20,
		},
	}

	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	// apiKey := generateApiKey()
	fake := "1"

	fmt.Println("Retrieving user data...")
	userDataFirst, err := getUsers(url, fake)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataFirst)
	fmt.Println("---------------")

	fmt.Println("Creating new character...")
	creationResponse, err := createUser(url, fake, userToCreate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	jsonData, _ := json.Marshal(creationResponse)
	fmt.Printf("Creation response body: %s\n", string(jsonData))
	fmt.Println("---------------")

	fmt.Println("Retrieving user data...")
	userDataSecond, err := getUsers(url, fake)
	if err != nil {
		fmt.Println("Error retrieving user:", err)
		return
	}

	logUsers(userDataSecond)
	fmt.Println("---------------")
}

func generateApiKey() string {
	const characters = "ABCDEF0123456789"
	var result string
	rand.New(rand.NewSource(0))
	for i := 0; i < len(characters); i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}
