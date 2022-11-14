package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	fmt.Printf("Part 1: %d\n", One(input))
}

func One(data []string) int {
	pub0, _ := strconv.Atoi(data[0])
	pub1, _ := strconv.Atoi(data[1])

	val := 1

	loops := make([]int, 2)
	for l := 1; loops[0] == 0 || loops[1] == 0; l++ {
		val *= 7
		val %= 20201227

		if val == pub0 {
			loops[0] = l
		}

		if val == pub1 {
			loops[1] = l
		}
	}

	if loops[0] != 0 {
		key := 1
		for i := 0; i < loops[0]; i++ {
			key *= pub1 // Use other public key
			key %= 20201227
		}

		return key
	}

	if loops[1] != 0 {
		key := 1
		for i := 0; i < loops[0]; i++ {
			key *= pub0 // Use other public key
			key %= 20201227
		}

		return key
	}

	return -1
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
