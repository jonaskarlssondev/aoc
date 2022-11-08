package main

import "testing"

func TestOne(t *testing.T) {
	tests := []struct {
		input  string
		moves  int
		output string
	}{
		{
			input:  "389125467",
			moves:  10,
			output: "92658374",
		},
		{
			input:  "389125467",
			moves:  100,
			output: "67384529",
		},
	}

	for _, test := range tests {
		got := One(test.input, test.moves)

		if got != test.output {
			t.Fatalf("Got: %s, Expected: %s", got, test.output)
		}
	}
}
