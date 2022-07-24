package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func main() {
	one(3, 1)
	two()
}

func one(xInc, yInc int) int {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	x := 0

	trees := 0

	for y := yInc; y < len(input); y += yInc {
		line := input[y]

		x = (x + xInc) % len(line)

		if line[x] == '#' {
			trees++
		}

	}

	fmt.Printf("Hit trees: %d for x: %d, y: %d.\n", trees, xInc, yInc)
	return trees
}

func two() {
	sum := 1

	sum *= one(1, 1)
	sum *= one(3, 1)
	sum *= one(5, 1)
	sum *= one(7, 1)
	sum *= one(1, 2)

	fmt.Println("Sum is: ", sum)
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
