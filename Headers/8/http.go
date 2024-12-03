package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err := encoder.Encode(data)
	if err != nil {
		return User{}, fmt.Errorf("error encoding data: %v", err)
	}

	req, err := http.NewRequest("PUT", fullURL, &buf)
	if err != nil {
		return User{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("error: received non-OK status code: %d", resp.StatusCode)
	}

	var updatedUser User
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&updatedUser)
	if err != nil {
		return User{}, fmt.Errorf("error decoding response data: %v", err)
	}
	return updatedUser, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return User{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("error: received non-OK status code: %d", resp.StatusCode)
	}

	var newUser User
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		return User{}, fmt.Errorf("error decoding response data: %v", err)
	}
	return newUser, nil
}
