package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://api.boot.dev/v1/courses_rest_api/learn-http/projects"

type Projects struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Assignees int    `json:"assignees"`
}

func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}
	defer res.Body.Close()

	var projects []Projects
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&projects)
	if err != nil {
		log.Fatalf("error encoding response: %v", err)
	}

	for _, p := range projects {
		fmt.Printf("Title: %s\n", p.Title)
		fmt.Printf("ID: %s\n", p.Id)
		fmt.Printf("Completed: %t\n", p.Completed)
		fmt.Printf("Assignees: %d\n", p.Assignees)
		fmt.Println("=======================")
	}
}
