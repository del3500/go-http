package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func isValidJSON(input string) bool {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	return err == nil
}

func TestIsValidJSON(t *testing.T) {
	type testCase struct {
		input string
	}
	tests := []testCase{
		{itemList},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{playerObject},
		}...)
	}

	for _, test := range tests {
		if output := isValidJSON(test.input); !output {
			t.Errorf(`Test Failed. Input:
%v
  =>
expected isValidJSON: %v
actual isValidJSON: %v
`,
				test.input, true, output)
		} else {
			fmt.Printf(`Test Passed. Input:
%v
  =>
expected isValidJSON: %v
actual isValidJSON: %v
`,
				test.input, true, output)
		}
		fmt.Println("==============================")
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
