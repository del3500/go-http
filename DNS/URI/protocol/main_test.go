package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		email    string
		expected string
	}

	tests := []testCase{
		{"slayer@fquest.app", "mailto:slayer@fquest.app"},
		{"killer@fquest.app", "mailto:killer@fquest.app"},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"", "mailto:"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMailtoLinkForEmail(test.email)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Email:		%v
Expecting:  %v
Actual:     %v
Fail`, test.email, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Email:		%v
Expecting:  %v
Actual:     %v
Pass
`, test.email, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
