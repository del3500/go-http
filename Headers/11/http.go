package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	/*req.Header.Set("Authorization", "Bearer "+apiKey)*/
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println("request to delete location unsuccessful")
		return err
	}
	fmt.Println("deletion successful")
	return nil
}
