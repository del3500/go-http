package main

import "fmt"

func main() {
	const URL = "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	uuid := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"

	user, err := getUserByID(URL, uuid)
	if err != nil {
		fmt.Println(err)
		return
	}
	logUser(user)
}
