package main 

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: " TEST! CaSe? ",
			expected: []string{"test", "case"},
		},
		{
			input: "here'S a, longER! TES,,./T case!",
			expected: []string{"here's","a","longer","tes,,./t", "case"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) { 
			t.Errorf("Actual length != expected length")
			continue
		}
		
		for i, word := range actual {
			if word != c.expected[i] { 
				t.Errorf("Actual word: '%s' not equal to expected word '%s'", word, c.expected[i])
			}
		}
	}

}
