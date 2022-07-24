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

	sum := 0
	answeredQuestions := make(map[string]bool)
	for i, line := range input {
		if line != "" {
			for _, question := range line {
				answeredQuestions[string(question)] = true
			}
		}

		if line == "" || i == len(input)-1 {
			sum += len(answeredQuestions)
			answeredQuestions = make(map[string]bool)
		}
	}

	fmt.Println("Sum: ", sum)
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	sum := 0
	count := 0
	answeredQuestions := []rune{}
	for i, line := range input {
		if line != "" {
			if count == 0 {
				answeredQuestions = []rune(line)
			} else {
				answeredQuestions = intersect(answeredQuestions, []rune(line))
			}

			count++
		}

		if line == "" || i == len(input)-1 {
			sum += len(answeredQuestions)
			answeredQuestions = []rune{}
			count = 0
		}
	}

	fmt.Println("Sum: ", sum)
}

func intersect(a, b []rune) []rune {
	set := make([]rune, 0)
	hash := make(map[rune]bool)

	for _, v := range a {
		hash[v] = true
	}

	for _, v := range b {
		if hash[v] {
			set = append(set, v)
		}
	}

	return set
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
