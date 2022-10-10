package main

import "testing"

func TestParse(t *testing.T) {

	test := struct {
		input  []string
		output map[spot]bool
	}{
		input: []string{
			".#.", "..#", "###",
		},
		output: map[spot]bool{
			{0, 0, 0, 0}: false,
			{1, 0, 0, 0}: true,
			{2, 0, 0, 0}: false,
			{0, 1, 0, 0}: false,
			{1, 1, 0, 0}: false,
			{2, 1, 0, 0}: true,
			{0, 2, 0, 0}: true,
			{1, 2, 0, 0}: true,
			{2, 2, 0, 0}: true,
		},
	}

	got := ParseInput(test.input)
	for spot, val := range got {
		out, ok := test.output[spot]
		if !ok || out != val {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}

func TestOne(t *testing.T) {
	tests := []struct {
		input  map[spot]bool
		runs   int
		output int
	}{
		{
			input: map[spot]bool{
				{0, 0, 0, 0}: false,
				{1, 0, 0, 0}: true,
				{2, 0, 0, 0}: false,
				{0, 1, 0, 0}: false,
				{1, 1, 0, 0}: false,
				{2, 1, 0, 0}: true,
				{0, 2, 0, 0}: true,
				{1, 2, 0, 0}: true,
				{2, 2, 0, 0}: true,
			},
			runs:   1,
			output: 11,
		},
		{
			input: map[spot]bool{
				{0, 0, 0, 0}: false,
				{1, 0, 0, 0}: true,
				{2, 0, 0, 0}: false,
				{0, 1, 0, 0}: false,
				{1, 1, 0, 0}: false,
				{2, 1, 0, 0}: true,
				{0, 2, 0, 0}: true,
				{1, 2, 0, 0}: true,
				{2, 2, 0, 0}: true,
			},
			runs:   2,
			output: 21,
		},
		{
			input: map[spot]bool{
				{0, 0, 0, 0}: false,
				{1, 0, 0, 0}: true,
				{2, 0, 0, 0}: false,
				{0, 1, 0, 0}: false,
				{1, 1, 0, 0}: false,
				{2, 1, 0, 0}: true,
				{0, 2, 0, 0}: true,
				{1, 2, 0, 0}: true,
				{2, 2, 0, 0}: true,
			},
			runs:   3,
			output: 38,
		},
		{
			input: map[spot]bool{
				{0, 0, 0, 0}: false,
				{1, 0, 0, 0}: true,
				{2, 0, 0, 0}: false,
				{0, 1, 0, 0}: false,
				{1, 1, 0, 0}: false,
				{2, 1, 0, 0}: true,
				{0, 2, 0, 0}: true,
				{1, 2, 0, 0}: true,
				{2, 2, 0, 0}: true,
			},
			runs:   6,
			output: 112,
		},
	}

	for _, test := range tests {
		got := One(test.input, test.runs)
		if got != test.output {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}
