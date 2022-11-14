package main

import "testing"

func TestCoords(t *testing.T) {
	tests := []struct {
		input  string
		output [2]float32
	}{
		{
			input:  "sesenwnw",
			output: [2]float32{0, 0},
		},
		{
			input:  "eew",
			output: [2]float32{1, 0},
		},
		{
			input:  "sene",
			output: [2]float32{1, 0},
		},
		{
			input:  "nwnwne",
			output: [2]float32{-0.5, 3},
		},
	}

	for _, test := range tests {
		got := getFinalCoords(test.input)

		if got != test.output {
			t.Fatalf("Got: %v, Expected: %v", got, test.output)
		}
	}
}

func TestTwo(t *testing.T) {
	test := struct {
		input  []string
		output int
	}{
		input: []string{
			"sesenwnenenewseeswwswswwnenewsewsw",
			"neeenesenwnwwswnenewnwwsewnenwseswesw",
			"seswneswswsenwwnwse",
			"nwnwneseeswswnenewneswwnewseswneseene",
			"swweswneswnenwsewnwneneseenw",
			"eesenwseswswnenwswnwnwsewwnwsene",
			"sewnenenenesenwsewnenwwwse",
			"wenwwweseeeweswwwnwwe",
			"wsweesenenewnwwnwsenewsenwwsesesenwne",
			"neeswseenwwswnwswswnw",
			"nenwswwsewswnenenewsenwsenwnesesenew",
			"enewnwewneswsewnwswenweswnenwsenwsw",
			"sweneswneswneneenwnewenewwneswswnese",
			"swwesenesewenwneswnwwneseswwne",
			"enesenwswwswneneswsenwnewswseenwsese",
			"wnwnesenesenenwwnenwsewesewsesesew",
			"nenewswnwewswnenesenwnesewesw",
			"eneswnwswnwsenenwnwnwwseeswneewsenese",
			"neswnwewnwnwseenwseesewsenwsweewe",
			"wseweeenwnesenwwwswnew",
		},
		output: 2208,
	}

	got := Two(test.input)

	if got != test.output {
		t.Fatalf("Got: %v, Expected: %v", got, test.output)
	}

}
