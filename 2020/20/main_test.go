package main

import "testing"

func TestParseOne(t *testing.T) {
	test := struct {
		input  []string
		output []image
	}{
		input: []string{
			"Tile 1621:",
			".#.##...#.",
			"#..#..#.#.",
			"#.#..#..##",
			".....#..#.",
			".#..#...##",
			"#....#...#",
			".#........",
			"#.#.#....#",
			"...#...#..",
			".#..#....#",
			"",
			"Tile 3671:",
			"..#.#.###.",
			"#.##....##",
			"#.........",
			"##..#.#...",
			"#..###....",
			"..#.#....#",
			"##..###..#",
			"..#......#",
			".........#",
			"......###.",
			"",
		},
		output: []image{
			{
				id:     1621,
				top:    ".#.##...#.",
				right:  "..#.##.#.#",
				bottom: ".#..#....#",
				left:   ".##..#.#..",
				data: [][]rune{
					[]rune("..#..#.#"),
					[]rune(".#..#..#"),
					[]rune("....#..#"),
					[]rune("#..#...#"),
					[]rune("....#..."),
					[]rune("#......."),
					[]rune(".#.#...."),
					[]rune("..#...#."),
				},
			},
			{
				id:     3671,
				top:    "..#.#.###.",
				right:  ".#...####.",
				bottom: "......###.",
				left:   ".####.#...",
				data: [][]rune{
					[]rune(".##....#"),
					[]rune("........"),
					[]rune("#..#.#.."),
					[]rune("..###..."),
					[]rune(".#.#...."),
					[]rune("#..###.."),
					[]rune(".#......"),
					[]rune("........"),
				},
			},
		},
	}

	got := ParseInput(test.input)

	if len(got) != 2 {
		t.Errorf("Got: %v, want: %v", got, test.output)
	}

	for i, img := range test.output {
		if got[i].id != img.id ||
			got[i].top != img.top ||
			got[i].right != img.right ||
			got[i].bottom != img.bottom ||
			got[i].left != img.left {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}

func TestReverseMatrix(t *testing.T) {
	test := struct {
		input  [][]rune
		output [][]rune
	}{
		input: [][]rune{
			[]rune("..#"),
			[]rune("..#"),
			[]rune("..#"),
		},
		output: [][]rune{
			[]rune("#.."),
			[]rune("#.."),
			[]rune("#.."),
		},
	}

	got := reverseMatrix(test.input)

	if len(test.output) != len(got) {
		t.Fatalf("Got %v, want: %v", got, test.output)
	}

	for i, r := range test.output {
		if string(got[i]) != string(r) {
			t.Fatalf("Got %v, want: %v", got, test.output)
		}
	}
}

func TestFlip(t *testing.T) {
	test := struct {
		input  [][]rune
		output [][]rune
	}{
		input: [][]rune{
			[]rune("..#"),
			[]rune(".#."),
			[]rune("#.."),
		},
		output: [][]rune{
			[]rune("#.."),
			[]rune(".#."),
			[]rune("..#"),
		},
	}

	got := flip(test.input)

	if len(test.output) != len(got) {
		t.Fatalf("Got %v, want: %v", got, test.output)
	}

	for i, r := range test.output {
		if string(got[i]) != string(r) {
			t.Fatalf("Got %v, want: %v", got, test.output)
		}
	}
}

func TestTranspose(t *testing.T) {
	test := struct {
		input  [][]rune
		output [][]rune
	}{
		input: [][]rune{
			[]rune("..#"),
			[]rune("#.#"),
			[]rune("#.."),
		},
		output: [][]rune{
			[]rune(".##"),
			[]rune("..."),
			[]rune("##."),
		},
	}

	got := transpose(test.input)

	if len(test.output) != len(got) {
		t.Fatalf("Got %v, want: %v", got, test.output)
	}

	for i, r := range test.output {
		if string(got[i]) != string(r) {
			t.Fatalf("Got %v, want: %v", got, test.output)
		}
	}
}

func TestRotate90(t *testing.T) {
	test := struct {
		input  [][]rune
		output [][]rune
	}{
		input: [][]rune{
			[]rune("#.#"),
			[]rune("##."),
			[]rune("#.."),
		},
		output: [][]rune{
			[]rune("###"),
			[]rune(".#."),
			[]rune("..#"),
		},
	}

	got := rotate90(test.input)

	if len(test.output) != len(got) {
		t.Fatalf("Got %v, want: %v", got, test.output)
	}

	for i, r := range test.output {
		if string(got[i]) != string(r) {
			t.Fatalf("Got %v, want: %v", got, test.output)
		}
	}
}
