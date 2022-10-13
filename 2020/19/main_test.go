package main

import "testing"

func TestParseOne(t *testing.T) {
	test := struct {
		input  []string
		output []string
	}{
		input: []string{
			"0: 4 1 5",
			"1: 2 3 | 3 2",
			"2: 4 4 | 5 5",
			"3: 4 5 | 5 4",
			"4: \"a\"",
			"5: \"b\"",
		},
		output: []string{
			"aaaabb", "aaabab", "abbabb", "abbbab", "aabaab", "aabbbb", "abaaab", "ababbb",
		},
	}

	got := ParseInput(test.input, false).valid
	if len(got) != len(test.output) {
		t.Errorf("Got: %v, want: %v", got, test.output)
	}
	for _, res := range got {
		if !contains(test.output, res) {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}
