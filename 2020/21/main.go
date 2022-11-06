package main

import (
	"aoc/util"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	data := parseInput(input)
	fmt.Printf("Part 1: %d\n", One(data))
}

type food struct {
	ingredients []string
	allergens   []string
}

func One(data []food) int {
	possible := map[string][]string{}
	appearances := map[string]int{}

	for _, food := range data {
		for _, i := range food.ingredients {
			appearances[i]++
		}

		for _, a := range food.allergens {
			if possible[a] == nil {
				possible[a] = food.ingredients
			} else {
				possible[a] = intersect(possible[a], food.ingredients)
			}
		}
	}

	for {
		multiple := true
		for a, p := range possible {
			if len(p) != 1 {
				multiple = false
			} else {
				for a2, p2 := range possible {
					if a != a2 {
						possible[a2] = remove(p2, p[0])
					}
				}
			}
		}

		if multiple {
			break
		}
	}

	for _, i := range possible {
		delete(appearances, i[0])
	}

	sum := 0
	for _, c := range appearances {
		sum += c
	}

	keys := []string{}
	for a := range possible {
		keys = append(keys, a)
	}
	sort.Strings(keys)

	dangerous := []string{}
	for _, a := range keys {
		dangerous = append(dangerous, possible[a][0])
	}

	fmt.Println("Part 2: " + strings.Join(dangerous, ","))
	return sum
}

func remove(s []string, v string) []string {
	res := []string{}
	for _, e := range s {
		if e != v {
			res = append(res, e)
		}
	}

	return res
}

func intersect(s1, s2 []string) []string {
	res := []string{}
	exists := map[string]bool{}

	for _, s := range s1 {
		exists[s] = true
	}

	for _, s := range s2 {
		if exists[s] {
			res = append(res, s)
			delete(exists, s)
		}
	}

	return res
}

func parseInput(s []string) []food {
	foods := []food{}

	for _, r := range s {
		sep := strings.Split(r, " (contains ")
		ingredients := strings.Split(sep[0], " ")
		allergens := strings.Split(sep[1][:len(sep[1])-1], ", ")

		f := food{
			ingredients: ingredients,
			allergens:   allergens,
		}

		foods = append(foods, f)
	}

	return foods
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
