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

	_, acc := execute(input)
	fmt.Println("Accumulator:", acc)
}

func execute(input []string) (bool, int) {
	acc := 0
	checked := make(map[int]bool)
	i := 0
	for i < len(input) {
		// Don't run same instruction twice.
		_, visited := checked[i]
		if visited {
			break
		}
		checked[i] = true

		op, cmd := util.Split(input[i], " ")

		// Get increment without sign.
		inc, err := strconv.Atoi(cmd[1:])
		if err != nil {
			exit(err)
		}

		switch op {
		case "nop":
			i++
		case "jmp":
			if cmd[0] == '+' {
				i += inc
			} else {
				i -= inc
			}
		case "acc":
			if cmd[0] == '+' {
				acc += inc
			} else {
				acc -= inc
			}
			i++
		}
	}

	return i == len(input), acc
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	for i, line := range input {
		instructions := make([]string, len(input))
		copy(instructions, input)

		op, cmd := util.Split(line, " ")
		if op == "jmp" {
			instructions[i] = "nop " + cmd
		} else if op == "nop" {
			instructions[i] = "jmp " + cmd
		}

		valid, acc := execute(instructions)
		if valid {
			fmt.Println("Found valid instruction:", acc)
			break
		}
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
