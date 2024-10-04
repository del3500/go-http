package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		inputUrl string
		expected ParsedURL
	}

	tests := []testCase{
		{
			"http://dragonslayer:pwn3d@fantasyquest.com:8080/maps?sort=rank#id",
			ParsedURL{
				protocol: "http",
				username: "dragonslayer",
				password: "pwn3d",
				hostname: "fantasyquest.com",
				port:     "8080",
				pathname: "/maps",
				search:   "sort=rank",
				hash:     "id",
			},
		},
		{
			"https://fantasyquest.com/items?sort=quality",
			ParsedURL{
				protocol: "https",
				username: "",
				password: "",
				hostname: "fantasyquest.com",
				port:     "",
				pathname: "/items",
				search:   "sort=quality",
				hash:     "",
			},
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{"", ParsedURL{}},
			{"://example.com", ParsedURL{}},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		parsedUrl := newParsedURL(test.inputUrl)
		if !reflect.DeepEqual(parsedUrl, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Fail
`, test.inputUrl, test.expected, parsedUrl)

		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, test.inputUrl, test.expected, parsedUrl)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
