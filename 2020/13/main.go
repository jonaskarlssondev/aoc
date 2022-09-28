package main

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
	"strconv"
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

	time, _ := strconv.Atoi(input[0])

	lines := strings.Split(input[1], ",")
	buses := busses(lines)

	min := math.MaxInt32
	bus := 0
	for _, b := range buses {
		bustime, _ := strconv.Atoi(b)

		timeToNext := bustime - (time % bustime)
		if timeToNext < min {
			min = timeToNext
			bus = bustime
		}
	}
	fmt.Printf("Shortest time til next: %d\n", min*bus)
}

func busses(s []string) []string {
	var busses []string
	for _, v := range s {
		if v != "x" {
			busses = append(busses, v)
		}
	}
	return busses
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	lines := strings.Split(input[1], ",")
	busoffset := make(map[int]int, 0)
	for i, v := range lines {
		if v != "x" {
			bus, _ := strconv.Atoi(v)
			busoffset[i] = bus
		}
	}

	earliest := 0
	inc := 1
	// This only works because we are guaranteed to find a match
	for i, bus := range busoffset {
		// TLDR, use Least Common Multiple for bus, bus+index, and LCM for previous buses

		// For each bus, create an increment that matches this bus and all previous busses
		// If we have bus3 at index 1, and bus5 at index 3, then they need to depart using the following pattern
		// BusX at 0 -> every minute, Bus3 at 1 every 3+1 minute, bus5 at 3 every 5+3 minute.
		// However it's not until every 12th minute that index and (bus+index) aligns for bus3
		// Bus5 thus needs to align 5, (5+3), and 12, which means a check every 120 minute (LCM)
		for (earliest+i)%bus != 0 {
			earliest += inc
		}
		inc *= bus
	}

	fmt.Printf("Earliest time found is: %d\n", earliest)
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
