package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	fmt.Printf("Part 1: %d\n", One(input))
	fmt.Printf("Part 2: %d\n", Two(input))
}

func One(data []string) int {
	tiles := map[[2]float32]bool{}

	for _, line := range data {
		pos := getFinalCoords(line)
		tiles[pos] = !tiles[pos]
	}

	sum := 0
	for _, isBlack := range tiles {
		if isBlack {
			sum++
		}
	}

	return sum
}

func Two(data []string) int {
	tiles := map[[2]float32]bool{}

	// Initial state
	for _, line := range data {
		pos := getFinalCoords(line)
		tiles[pos] = !tiles[pos]
	}

	for i := 0; i < 100; i++ {
		copy := map[[2]float32]bool{}

		toCheck := map[[2]float32]bool{}
		for pos, black := range tiles {
			toCheck[pos] = black

			// Any tile that neighbours a black tile might not be in "tiles"
			// and should be checked
			check := getNeighbours(pos[0], pos[1])
			for _, n := range check {
				if _, ok := toCheck[n]; !ok {
					toCheck[n] = false // Must be white
				}
			}
		}

		for pos, black := range toCheck {
			// number of black neighbours for this tile
			neighbours := nNeighbours(tiles, pos)

			if black && (neighbours == 0 || neighbours > 2) {
				copy[pos] = false
			} else if !black && neighbours == 2 {
				copy[pos] = true
			} else {
				copy[pos] = tiles[pos]
			}
		}

		tiles = copy
	}

	sum := 0
	for _, isBlack := range tiles {
		if isBlack {
			sum++
		}
	}

	return sum
}

func nNeighbours(tiles map[[2]float32]bool, tile [2]float32) int {
	n := 0
	x, y := tile[0], tile[1]

	if tiles[[2]float32{x + 1, y}] {
		n++
	} // e
	if tiles[[2]float32{x + 0.5, y - 1}] {
		n++
	} // se
	if tiles[[2]float32{x - 0.5, y - 1}] {
		n++
	} // sw
	if tiles[[2]float32{x - 1, y}] {
		n++
	} // w
	if tiles[[2]float32{x - 0.5, y + 1}] {
		n++
	} // nw
	if tiles[[2]float32{x + 0.5, y + 1}] {
		n++
	} // ne

	return n
}

func getNeighbours(x, y float32) [6][2]float32 {
	return [6][2]float32{
		[2]float32{x + 1, y},
		[2]float32{x + 0.5, y - 1},
		[2]float32{x - 0.5, y - 1},
		[2]float32{x - 1, y},
		[2]float32{x - 0.5, y + 1},
		[2]float32{x + 0.5, y + 1},
	}
}

func getFinalCoords(move string) [2]float32 {
	pos := [2]float32{0, 0}

	for i := 0; i < len(move); i++ {
		str := string(move[i])
		if str != "w" && str != "e" {
			str = string(move[i : i+2])
			i++
		}

		delta := deltaCoordinates[str]
		pos[0] += delta[0]
		pos[1] += delta[1]
	}

	return pos
}

// Converts a step to a new tile to changes in location
// This is so e.g. a se->ne is translated same as e.
var deltaCoordinates = map[string][2]float32{
	"e":  [2]float32{1, 0},
	"se": [2]float32{0.5, -1},
	"sw": [2]float32{-0.5, -1},
	"w":  [2]float32{-1, 0},
	"nw": [2]float32{-0.5, 1},
	"ne": [2]float32{0.5, 1},
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
