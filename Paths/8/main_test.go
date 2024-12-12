package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

	type testCase struct {
		availability string
		expected     []Issue
	}

	tests := []testCase{
		{
			availability: "Low",
			expected: []Issue{
				{Title: "Add more console.logs because why not", Estimate: 5},
			},
		},
		{
			availability: "Medium",
			expected: []Issue{
				{Title: "Add more console.logs because why not", Estimate: 5},
				{Title: "Optimize database queries for better performance", Estimate: 5},
				{Title: "Resolve production deployment issues", Estimate: 5},
			},
		},
	}

	if withSubmit {
		tests = append(tests, testCase{
			availability: "High",
			expected: []Issue{
				{Title: "Add more console.logs because why not", Estimate: 5},
				{Title: "Optimize database queries for better performance", Estimate: 5},
				{Title: "Resolve production deployment issues", Estimate: 5},
				{Title: "Implement user authentication flow", Estimate: 6},
				{Title: "Increase code coverage by deleting untested code", Estimate: 7},
			},
		})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		issues := fetchTasks(url, test.availability)

		passed := true
		for i, expectedIssue := range test.expected {
			if i >= len(issues) {
				passed = false
				break
			}

			if expectedIssue.Title != issues[i].Title || expectedIssue.Estimate != issues[i].Estimate {
				passed = false
				break
			}
		}

		if passed {
			passCount++
			fmt.Printf(`---------------------------------
Rarity:     %s
Expected:
%v
Actual:
%v
PASS
`, test.availability, logIssues(test.expected), logIssues(issues))
		} else {
			failCount++
			t.Errorf(`---------------------------------
Rarity:     %s
Expected:
%v
Actual:
%v
FAIL`, test.availability, logIssues(test.expected), logIssues(issues))
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
