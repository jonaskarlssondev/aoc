package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	// Convert to matrix of single char strings
	seats := matrix(input)

	changes := -1
	for changes != 0 {
		// Loop and then assign copy to seats matrix
		seats, changes = eval(seats, 4, false)
	}

	sum := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == "#" {
				sum++
			}
		}
	}

	fmt.Printf("Occupied seats: %d\n", sum)
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	// Convert to matrix of single char strings
	seats := matrix(input)

	changes := -1
	for changes != 0 {
		// Loop and then assign copy to seats matrix
		seats, changes = eval(seats, 5, true)
	}

	sum := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == "#" {
				sum++
			}
		}
	}

	fmt.Printf("Occupied seats: %d\n", sum)
}

func matrix(input []string) [][]string {
	var matrix [][]string

	for _, line := range input {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix
}

func eval(seats [][]string, tol int, part2 bool) ([][]string, int) {
	changes := 0
	var copy [][]string

	for r, row := range seats {
		copy = append(copy, make([]string, len(seats[0])))
		for c, seat := range row {
			if seat == "." {
				// Dont check neighours if it doesnt matter
				copy[r][c] = "."
				continue
			}

			n := neighbours(r, c, seats, part2)
			if seat == "L" && n == 0 {
				copy[r][c] = "#"
				changes++
			} else if seat == "#" && n >= tol {
				copy[r][c] = "L"
				changes++
			} else {
				copy[r][c] = seat
			}
		}
	}

	return copy, changes
}

func neighbours(x, y int, seats [][]string, part2 bool) int {
	neighbours := 0
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}

			row := x + i
			col := y + j

			n, found := check(row, col, seats)
			if part2 {
				for !found {
					row += i
					col += j
					n, found = check(row, col, seats)
				}
			}

			neighbours += n
		}
	}

	return neighbours
}

// Returns 1 if it found an empty seat, and then false if floor was found
func check(row, col int, seats [][]string) (int, bool) {
	if row < 0 || col < 0 || row >= len(seats) || col >= len(seats[0]) {
		return 0, true
	}

	if seats[row][col] == "#" {
		return 1, true
	} else if seats[row][col] == "L" {
		return 0, true
	}

	// If floor we want to continue search
	return 0, false
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
