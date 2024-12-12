package main

import (
	"fmt"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		id       string
		expected int
	}

	tests := []testCase{
		{"invalid-id", http.StatusNotFound},
		{"0194fdc2-fa2f-4cc0-81d3-ff12045b73c8", http.StatusOK},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users/" + test.id
		code := getUserCode(url)

		if code != test.expected {
			failCount++
			t.Errorf(`---------------------------------
ID:         %s
URL:        %s
Expecting:  %v
Actual:     %v
Fail
`, test.id, url, test.expected, code)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
ID:         %s
URL:        %s
Expecting:  %v
Actual:     %v
Pass
`, test.id, url, test.expected, code)
		}
	}

	fmt.Println("------------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
