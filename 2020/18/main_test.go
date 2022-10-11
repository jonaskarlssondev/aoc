package main

import "testing"

func TestParseOne(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "1 + 2 * 3 + 4 * 5 + 6",
			output: "12+3*4+5*6+",
		},
		{
			input:  "1 + (2 * 3) + (4 * (5 + 6))",
			output: "123*+456+*+",
		},
	}

	for _, test := range tests {
		got := parseOp(test.input, false).rpn
		if got != test.output {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}

func TestParseTwo(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "1 + 2 * 3 + 4 * 5 + 6",
			output: "12+34+*56+*",
		},
	}

	for _, test := range tests {
		got := parseOp(test.input, true).rpn
		if got != test.output {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}

func TestCompute(t *testing.T) {
	tests := []struct {
		input  operation
		output int
	}{
		{
			input:  operation{"12+3*4+5*6+"},
			output: 71,
		},
		{
			input:  operation{"123*+456+*+"},
			output: 51,
		},
		{
			input:  operation{"12+34+56+**"},
			output: 231,
		},
	}

	for _, test := range tests {
		got := test.input.compute()
		if got != test.output {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}
