package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		address   string
		expectErr bool
	}

	tests := []testCase{
		{"boot.dev", false},
		{"example.com", false},
		{"cloudflare.com", false},
	}

	// Additional test cases for submission
	if withSubmit {
		tests = append(tests, []testCase{
			{"iana.org", false},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output, err := getIPAddress(test.address)
		splitIP := strings.Split(output, ".")
		if err != nil && !test.expectErr {
			failCount++
			t.Errorf(`---------------------------------
URL:			%v
ExpectedErr:	%v
GotErr:			%v
Fail`, test.address, test.expectErr, err != nil)
		} else if len(splitIP) != 4 {
			failCount++
			t.Errorf(`---------------------------------
URL:			%v
Expected IP:	%v
Got:			%v
Fail`, test.address, true, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:			%v
IP Address:		%v
Pass
`, test.address, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
