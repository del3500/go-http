package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		url      string
		expected []Item
	}

	tests := []testCase{
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/items?limit=1",
			[]Item{{Name: "Healing Potion", Quality: 4}},
		},
		{
			"https://api.boot.dev/v1/courses_rest_api/learn-http/items?limit=2",
			[]Item{
				{Name: "Healing Potion", Quality: 4},
				{Name: "Light Leather", Quality: 1},
			},
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		items, _ := getItems(test.url)

		if !reflect.DeepEqual(items, test.expected) {
			failCount++
			t.Errorf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, test.url, test.expected, items)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
URL:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, test.url, test.expected, items)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
