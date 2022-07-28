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

type bag struct {
	colour string
	count  int
}

func buildBagMap(input []string) map[string][]bag {
	bagMap := make(map[string][]bag)

	// Sort all possible containers into a map where the key is the root colour and each container is
	// described by colour and the number of those.
	for _, line := range input {
		root, contains := split(line, " bags contain ")

		contained := []bag{}
		containedBags := strings.Split(contains, ", ")
		for _, cb := range containedBags {
			if cb == "no other bags." {
				continue
			}

			// Ignore number and space after, and only care about the color description
			colour, _ := split(cb[2:], " bag")
			n, _ := split(cb, " ")
			count, err := strconv.Atoi(n)
			if err != nil {
				exit(err)
			}

			b := bag{
				colour: colour,
				count:  count,
			}
			contained = append(contained, b)
		}

		bagMap[root] = contained
	}

	return bagMap
}

func one() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	bagMap := buildBagMap(input)

	set := make(map[string]bool)

	// Loop all root colours and find recursively find if the search colour is a child of any container for the root colour.
	for e := range bagMap {
		if e == "shiny gold" {
			continue
		}

		checked := make(map[string]bool)
		if canContain(bagMap, checked, e, "shiny gold") {
			set[e] = true
		}
	}

	fmt.Println("Count:", len(set))
}

func canContain(bags map[string][]bag, checked map[string]bool, colour string, search string) bool {
	if val, ok := checked[colour]; ok {
		return val
	}

	if colour == search {
		return false
	}

	if contains(bags[colour], search) {
		checked[colour] = true
		return true
	}

	for _, c := range bags[colour] {
		if canContain(bags, checked, c.colour, search) {
			return true
		}
	}

	checked[colour] = false
	return false
}

func contains(containers []bag, colour string) bool {
	for _, v := range containers {
		if v.colour == colour {
			return true
		}
	}

	return false
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	bagMap := buildBagMap(input)

	count := 0
	for _, bag := range bagMap["shiny gold"] {
		count += bag.count + bag.count*numberOfBags(bagMap, bag.colour)
	}

	fmt.Println("Count:", count)
}

func numberOfBags(bags map[string][]bag, colour string) int {
	if len(bags[colour]) == 0 {
		return 0
	}

	count := 0
	for _, b := range bags[colour] {
		count += b.count + b.count*numberOfBags(bags, b.colour)
	}

	return count
}

func split(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	return split[0], split[1]
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
