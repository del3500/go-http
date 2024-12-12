package main

func fetchTasks(baseURL, availability string) []Issue {
	switch availability {
	case "Low":
		availability = "1"
	case "Medium":
		availability = "3"
	case "High":
		availability = "5"
	}
	fullURL := baseURL + "?sort=estimate&limit=" + availability
	return getIssues(fullURL)
}
