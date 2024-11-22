package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"golang.org/x/exp/rand"
)

type Project struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Assignees int    `json:"assignees"`
}

func generateKey() string {
	const characters = "ABCDEF0123456789"
	var result string
	rand.New(rand.NewSource(0))
	for i := 0; i < len(characters); i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}

func getProjectResponse(apiKey, url string) (Project, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Project{}, err
	}
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Project{}, err
	}

	defer resp.Body.Close()

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return Project{}, err
	}
	return project, nil
}

func putProject(apiKey, url string, project Project) error {
	client := &http.Client{}
	data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
