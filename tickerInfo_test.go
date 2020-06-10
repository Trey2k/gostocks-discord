package main

import (
	"strconv"
	"testing"
)

func TestIsValidTicker(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"amzn", true},
		{"amzns", false},
		{"go", true},
		{"agagrg", false},
		{"g1", false},
		{"SPY", true},
		{"BYND", true},
		{"lol", true},
	}

	for _, test := range tests {
		if output := IsValidTicker(test.input); output != test.expected {
			t.Error("Test failed, inputted: '" + test.input + "', expected: '" + strconv.FormatBool(test.expected) + "', received: '" + strconv.FormatBool(output) + "'")
		}
	}
}
