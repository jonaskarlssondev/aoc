package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
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

	heading := 90
	headingMap := map[int]string{0: "N", 90: "E", 180: "S", 270: "W"}
	ew := 0
	ns := 0

	for _, line := range input {
		command := string(line[0])
		num, _ := strconv.Atoi(line[1:])
		cmd := command
		if command == "F" {
			cmd = headingMap[heading]
		}
		move(&ew, &ns, &heading, cmd, num)
	}

	fmt.Printf("Sum is: %d\n", abs(ew)+abs(ns))
}

func move(ew, ns, heading *int, cmd string, num int) {
	switch cmd {
	case "N":
		*ns += num
	case "S":
		*ns -= num
	case "E":
		*ew += num
	case "W":
		*ew -= num
	case "L":
		*heading = (*heading - num + 360) % 360
	case "R":
		*heading = (*heading + num + 360) % 360
	}
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	ew := 0
	ns := 0
	ewm := 10
	nsm := 1

	for _, line := range input {
		command := string(line[0])
		num, _ := strconv.Atoi(line[1:])

		if command == "F" {
			ew += ewm * num
			ns += nsm * num
		}
		move2(&ewm, &nsm, command, num)
	}

	fmt.Printf("Sum is: %d\n", abs(ew)+abs(ns))
}

func move2(ewm, nsm *int, cmd string, num int) {
	switch cmd {
	case "N":
		*nsm += num
	case "S":
		*nsm -= num
	case "E":
		*ewm += num
	case "W":
		*ewm -= num
	case "L":
		n := *nsm
		e := *ewm
		switch num {
		case 90:
			*ewm = -n
			*nsm = e
		case 180:
			*ewm = -e
			*nsm = -n
		case 270:
			*ewm = n
			*nsm = -e
		}
	case "R":
		n := *nsm
		e := *ewm
		switch num {
		case 90:
			*ewm = n
			*nsm = -e
		case 180:
			*ewm = -e
			*nsm = -n
		case 270:
			*ewm = -n
			*nsm = e
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
