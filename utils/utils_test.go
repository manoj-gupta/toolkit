package utils

import (
	"testing"
)

func TestIsPalindromeString(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"eatae", true},
		{"tea", false},
		{"tat", true},
	}

	for _, test := range tests {
		got := IsPalindromeString(test.input)

		if got != test.want {
			t.Errorf("Got not same as want")
		}
	}
}

func TestReverseString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"eatae", "eatae"},
		{"tea", "aet"},
		{"etat", "tate"},
	}

	for _, test := range tests {
		got := ReverseString(test.input)

		if got != test.want {
			t.Errorf("Got not same as want")
		}
	}
}
