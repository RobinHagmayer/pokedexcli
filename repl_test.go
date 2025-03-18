package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "hElLo",
			expected: []string{"hello"},
		},
		{
			input:    " HellO   ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf(`
Test Failed. Input:
%v
expected length: %v
actual length: %v`, c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			expectedWord := c.expected[i]
			word := actual[i]

			if word != expectedWord {
				t.Errorf(`
Test Failed. Input:
%v
expected cleanInput: %s
actual cleanInput: %s`, c.input, expectedWord, word)
			}
		}
	}
}
