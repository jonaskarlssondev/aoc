package main

import (
	"aoc/util"
	"fmt"
	"os"
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

	highest := 0

	for _, line := range input {
		row, col := findSeat(line)

		seat := row*8 + col

		if seat > highest {
			highest = seat
		}
	}

	fmt.Println("Highest ID:", highest)
}

func findSeat(line string) (row, col int) {
	minRow := 0
	maxRow := 127

	for row := 0; row < 8; row++ {
		if line[row] == 'F' {
			maxRow -= (maxRow - minRow + 1) / 2
		} else {
			minRow += (maxRow - minRow + 1) / 2
		}
	}

	minCol := 0
	maxCol := 7

	for col := 7; col < 10; col++ {
		if line[col] == 'L' {
			maxCol -= (maxCol - minCol + 1) / 2
		} else {
			minCol += (maxCol - minCol + 1) / 2
		}
	}

	row = minRow
	col = minCol
	return
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	check := [128][8]bool{}

	for _, line := range input {
		row, col := findSeat(line)
		check[row][col] = true
	}

	for r := 0; r < len(check); r++ {
		for c := 0; c < len(check[r]); c++ {
			if !check[r][c] {
				switch c {
				case 0:
					if r == 0 {
						break
					}
					if check[r][c+1] && check[r-1][7] {
						fmt.Println("Seat:", r*8+c)
					}
				case 7:
					if r == len(check)-1 {
						break
					}
					if check[r+1][0] && check[r][c-1] {
						fmt.Println("Seat:", r*8+c)
					}
				default:
					if check[r][c-1] && check[r][c+1] {
						fmt.Println("Seat:", r*8+c)
					}
				}
			}
		}
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
