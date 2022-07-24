package main

import (
	"aoc/util"
	"fmt"
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

	valid := 0

	for _, entry := range input {
		policy, password := split(entry, ": ")
		minStr, rest := split(policy, "-")
		maxStr, letter := split(rest, " ")
		min, _ := strconv.Atoi(minStr)
		max, _ := strconv.Atoi(maxStr)
		appearances := strings.Count(password, letter)
		if appearances >= min && appearances <= max {
			valid++
		}
	}

	fmt.Println("Valid: ", valid)
}

func split(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	return split[0], split[1]
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	valid := 0

	for _, entry := range input {
		policy, password := split(entry, ": ")
		strOne, rest := split(policy, "-")
		strTwo, letter := split(rest, " ")
		one, _ := strconv.Atoi(strOne)
		two, _ := strconv.Atoi(strTwo)

		x := string(password[one-1])
		y := string(password[two-1])

		if (x == letter || y == letter) && (x != y) {
			valid++
		}
	}

	fmt.Println("Valid: ", valid)
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
